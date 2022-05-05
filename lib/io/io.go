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
var buffer = make([]byte, 0, 1024*1024)
var writer = bufio.NewWriter(os.Stdout)

func init() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(buffer, 256*1024*1024)
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

func NextInt64() int64 {
	x, err := strconv.ParseInt(Next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func Println(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func Print(a ...interface{}) {
	fmt.Fprint(writer, a...)
}

func Flush() {
	writer.Flush()
}
