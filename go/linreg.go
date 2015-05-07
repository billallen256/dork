package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// example follows code from http://onlinestatbook.com/2/regression/intro.html

func sum(a *[]float64) float64 {
	s := 0.0

	for _, x := range *a {
		s += x
	}

	return s
}

func mean(a *[]float64) float64 {
	return sum(a) / float64(len(*a))
}

func std(a *[]float64) float64 {
	m := mean(a)
	s := 0.0

	for _, x := range *a {
		s += math.Pow(x-m, 2.0)
	}

	return math.Sqrt(s / float64(len(*a)))
}

func readVals(f *os.File) (*[]float64, *[]float64) {
	reader := csv.NewReader(f)
	xs := make([]float64, 0, 10)
	ys := make([]float64, 0, 10)

	for {
		r, err := reader.Read()

		if err != nil {
			break
		}

		x, xerr := strconv.ParseFloat(r[0], 64)
		y, yerr := strconv.ParseFloat(r[1], 64)

		if xerr != nil || yerr != nil {
			continue
		}

		xs = append(xs, x)
		ys = append(ys, y)
	}

	return &xs, &ys
}

func corr(X *[]float64, Y *[]float64) float64 {
	xmean := mean(X)
	ymean := mean(Y)
	//x := make([]float64, len(*X))
	//y := make([]float64, len(*Y))
	xy := make([]float64, len(*X))
	x2 := make([]float64, len(*X))
	y2 := make([]float64, len(*Y))

	for i := range *X {
		x := (*X)[i] - xmean
		y := (*Y)[i] - ymean
		xy[i] = x * y
		x2[i] = math.Pow(x, 2.0)
		y2[i] = math.Pow(y, 2.0)
	}

	return sum(&xy) / math.Sqrt(sum(&x2)*sum(&y2))
}

func randArray(size int) *[]float64 {
	r := make([]float64, size)

	for i := 0; i < size; i++ {
		r[i] = rand.Float64()
	}

	return &r
}

func main() {
	f, err := os.Open(os.Args[1])
	defer f.Close()

	if err != nil {
		os.Exit(1)
	}

	//xs, ys := readVals(f)
	xs := randArray(100)
	ys := randArray(100)

	for i := range *xs {
		fmt.Println((*xs)[i], (*ys)[i])
	}

	mx := mean(xs)
	my := mean(ys)
	sx := std(xs)
	sy := std(ys)
	r := corr(xs, ys)

	slope := r * (sy / sx)
	intercept := my - (slope * mx)

	fmt.Println("mean x", mx, "mean y", my)
	fmt.Println("std x", sx, "std y", sy)
	fmt.Println("r", r)
	fmt.Println("b", slope)
	fmt.Println("A", intercept)
}
