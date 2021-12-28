package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
}

var scanner Scanner = Scanner{bufio.NewScanner(os.Stdin)}

func init() {
	scanner.Split(bufio.ScanWords)
}

func Next() string {
	if !scanner.Scan() {
		panic(scanner.Err())
	}
	return scanner.Text()
}

func NextInt() int {
	x, err := strconv.Atoi(Next())
	if err != nil {
		panic(err)
	}
	return x
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func Print(a ...interface{}) {
	fmt.Print(a...)
}
