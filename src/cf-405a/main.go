package main

import (
	"lib/io"
)

func main() {
	n := io.NextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= 100; j++ {
			h := 0
			for k := 0; k < n; k++ {
				if a[k] >= j {
					h++
				}
			}
			if n-i <= h {
				b[i]++
			}
		}
	}
	for i := 0; i < n; i++ {
		io.Print(b[i], " ")
	}
	io.Println()
}
