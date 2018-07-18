package main

import "fmt"

func calcUsingLoop(n int) []int {
	p := make([]int, n, n)
	p[0] = 1
	p[1] = 1
	for i := 2; i < n; i++ {
		p[i] = p[i-1] + p[i-2]
	}
	return p
}

func main() {
	p := calcUsingLoop(10)
	fmt.Println(p)
}
