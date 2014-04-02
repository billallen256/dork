package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func consumer(id int, v chan int) {
	for {
		fmt.Println(id, <-v)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numConsumers := 3
	numRands := 1000
	numbers := make(chan int, 10)

	for i:=0 ; i < numConsumers ; i++ {
		go consumer(i, numbers)
	}

	for i:=0 ; i < numRands ; i++ {
		numbers <- rand.Int()
	}
}
