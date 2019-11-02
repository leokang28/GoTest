package main

import (
	"fmt"
	"unicode/utf8"
	)

func main() {
	str := "克里斯蒂01cristine"
	for a, b := range str {
		fmt.Println(a)
		fmt.Printf("unicode: %c %d\n", b, b)

		fmt.Println(utf8.RuneCountInString(str))
	}

}