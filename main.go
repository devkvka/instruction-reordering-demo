package main

import (
	"fmt"
	"sync"
)

var (
	x, y = 0, 0
	a, b = 0, 1
)

func t1() {
	x = 1
	a = y
}

func t2() {
	y = 1
	b = x
}

func main() {
	var wg sync.WaitGroup
	i := 0
	for ;a != 0 || b != 0; {
		a, b = 0, 0
		x, y = 0, 0

		wg.Go(t1)
		wg.Go(t2)

		wg.Wait()

		i++
	}

	fmt.Println("Found (a == 0 && b == 0) after", i, "iterations")
}
