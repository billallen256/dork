package main

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"os"
)

type trips struct {
	Routes []route `yaml:",flow"`
}

type route struct {
	Airports []string `yaml:",flow"`
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", fileName)
		os.Exit(1)
	}

	defer f.Close()

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not stat %s\n", fileName)
		os.Exit(1)
	}

	fileSize := fileInfo.Size()
	buff := make([]byte, fileSize)
	_, err = f.Read(buff)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %d bytes\n", len(buff))
		os.Exit(1)
	}

	var t trips
	yaml.Unmarshal(buff, &t)

	for _, r := range t.Routes {
		fmt.Println(r)
	}
}
