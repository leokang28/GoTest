package main

import "fmt"

func replace() {
	angle := "heros never die"
	angleBytes := []byte(angle)

	for i := 6; i < 11; i++ {
		angleBytes[i] = ' '
	}

	fmt.Println(string(angleBytes))

}

func unicode() {
	str := "嘻哈bcde"
	strBytes := []rune(str)

	for i := 0; i<len(strBytes); i++ {
		fmt.Println(string(strBytes[i]))
	}
}

func main() {
	replace()

	unicode()
}
