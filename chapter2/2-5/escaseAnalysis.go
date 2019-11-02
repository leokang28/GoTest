// 变量逃逸分析
package main

import "fmt"

type Data struct {}

func dummy(b int) int {
	var c int
	c = b
	return c
}

func dummy2() *Data {
	var d Data
	return &d
}

func void() {}

func main() {
	var a int

	void()

	fmt.Println(a, dummy(a))

	fmt.Println(dummy2())
}

// run with command 'go run -gcflags " -m -l" x.go'