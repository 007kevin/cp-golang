package main

import "lib/io"

func main() {
	defer io.Flush()
	h := io.NextInt()
	io.NextInt()
	for i := 0; i < h; i++ {
		s := io.Next()
		io.Println(s)
		io.Println(s)
	}
}
