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
		s += (x-m)*(x-m)
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

func corr(X, Y *[]float64) float64 {
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
		x2[i] = x*x
		y2[i] = y*y
	}

	return sum(&xy) / math.Sqrt(sum(&x2)*sum(&y2))
}

func linreg(X, Y *[]float64) (slope, intercept float64) {
	mx := mean(X)
	my := mean(Y)
	sx := std(X)
	sy := std(Y)
	r := corr(X, Y)

	slope = r * (sy / sx)
	intercept = my - (slope * mx)
	return slope, intercept
}

func coeffDetermination(X, Y *[]float64) float64 {
	N := len(*X)
	mx := mean(X)
	my := mean(Y)
	sx := std(X)
	sy := std(Y)
	s := 0.0

	for i := 0; i < N; i++ {
		s += ((*X)[i] - mx) * ((*Y)[i] - my)
	}

	R := (1/float64(N)) * (s / (sx * sy))
	return R * R
}

func randArray(size int) *[]float64 {
	r := make([]float64, size)

	for i := 0; i < size; i++ {
		r[i] = rand.Float64()
	}

	return &r
}

func main() {
	//f, err := os.Open(os.Args[2])
	num64, err := strconv.ParseInt(os.Args[1], 10, 64)
	num := int(num64)
	//defer f.Close()

	if err != nil {
		os.Exit(1)
	}

	//xs, ys := readVals(f)
	xs := randArray(num)
	ys := randArray(num)

	//for i := range *xs {
	//	fmt.Println((*xs)[i], (*ys)[i])
	//}

	slope, intercept := linreg(xs, ys)

	//fmt.Println("mean x", mx, "mean y", my)
	//fmt.Println("std x", sx, "std y", sy)
	//fmt.Println("r", r)
	fmt.Println("b", slope)
	fmt.Println("A", intercept)

	r2 := coeffDetermination(xs, ys)
	fmt.Println("R**2", r2)
}
