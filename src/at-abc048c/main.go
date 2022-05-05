package main

import (
	"lib/io"
)

func main() {
	n := io.NextInt()
	x := io.NextInt64()
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt64()
	}
	var cnt int64
	for i := 1; i < n; i++ {
		diff := a[i] + a[i-1] - x
		if diff > 0 {
			a[i] -= diff
			if a[i] < 0 {
				a[i-1] += a[i]
				a[i] = 0
			}
			cnt += diff
		}
	}
	io.Println(cnt)
}
