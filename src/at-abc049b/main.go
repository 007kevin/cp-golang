package main

import "lib/io"

func main() {
	h := io.NextInt()
	w := io.NextInt()
	c := make([][]byte, h)
	for i := 0; i < h; i++ {
		c[i] = []byte(io.Next())
	}
	cc := make([][]byte, 2*h)
	for i := 0; i < 2*h; i++ {
		cc[i] = make([]byte, w)
	}
	for i := 0; i < 2*h; i++ {
		for j := 0; j < w; j++ {
			cc[i][j] = c[(i)/2][j]
		}
	}
	for i := 0; i < 2*h; i++ {
		io.Println(string(cc[i]))
	}
}
