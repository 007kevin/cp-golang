package main

import (
	"lib/io"
)

func main() {
	defer io.Flush()
	words := []string{
		"dreamer",
		"eraser",
		"dream",
		"erase",
	}
	s := io.Next()
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if dp[i] {
			for _, w := range words {
				if matchAt(s, w, i) {
					dp[i+len(w)] = true
				}
			}
		}
	}
	if dp[len(s)] {
		io.Println("YES")
	} else {
		io.Println("NO")
	}
}

func matchAt(s string, w string, index int) bool {
	if index+len(w) > len(s) {
		return false
	}
	for i := 0; i < len(w); i++ {
		if s[index+i] != w[i] {
			return false
		}
	}
	return true
}
