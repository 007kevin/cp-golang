package main

import (
	"lib/io"
	"sort"
)

func main() {
	n := io.NextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}
	sort.Sort(sort.IntSlice(a))
	for i := 0; i < n; i++ {
		io.Print(a[i], " ")
	}
	io.Println()
}
