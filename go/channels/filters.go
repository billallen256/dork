package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Filter struct {
	Name string
	in   chan string
	out  chan string
}

func NewFilter(name string, f func(in string) string) *Filter {
	i := make(chan string, 10)
	o := make(chan string, 10)

	go func() {
		for {
			o <- f(<-i)
		}
	}()

	return &Filter{name, i, o}
}

func (filter *Filter) Process(in string) {
	filter.in <- in
}

func (filter *Filter) Result() string {
	return <-filter.out
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s\n", fileName)
		os.Exit(1)
	}

	// create the filters
	filters := make([]*Filter, 0, 3)
	filters = append(filters, NewFilter("upper", func(in string) string { return strings.ToUpper(in) }))
	filters = append(filters, NewFilter("len", func(in string) string { return fmt.Sprintf("%d %s", len(in), in) }))
	toFilters := make(chan string, 10)
	toPrinter := make(chan string, 10)

	// connect the filters
	go func() {
		for {
			next := <-toFilters

			for _, filter := range filters {
				filter.Process(next)
				next = filter.Result()
			}

			toPrinter <- next
		}
	}()

	// read from the file
	go func() {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Reading", line)
			toFilters <- line
		}
	}()

	// print the output
	for p := range toPrinter {
		fmt.Println(p)
	}
}
