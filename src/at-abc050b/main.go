package main

import "lib/io"

func main() {
	defer io.Flush()
	n := io.NextInt()
	t := make([]int, n+1)
	s := 0
	for i := 1; i <= n; i++ {
		t[i] = io.NextInt()
		s += t[i]
	}
	m := io.NextInt()
	p := make([]int, m+1)
	x := make([]int, m+1)
	for i := 1; i <= m; i++ {
		p[i] = io.NextInt()
		x[i] = io.NextInt()
	}
	for i := 1; i <= m; i++ {
		o := s
		o -= t[p[i]]
		o += x[i]
		io.Println(o)
	}
}
