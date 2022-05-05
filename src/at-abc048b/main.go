package main

import "lib/io"

func main() {
	a := io.NextInt64()
	b := io.NextInt64()
	x := io.NextInt64()
	a = ceil(a, x)
	b = floor(b, x)
	if a > b {
		io.Println(0)
	} else {
		io.Println((b-a)/x + 1)
	}
}

func ceil(a, x int64) int64 {
	if a%x == 0 {
		return a
	}
	return (a/x)*x + x
}

func floor(a, x int64) int64 {
	if a%x == 0 {
		return a
	}
	return (a / x) * x
}
