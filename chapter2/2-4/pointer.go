package __4

import (
	"fmt"
	"flag"
)

func checkVirtualAddress() {
	cat := 1
	str := "banana"

	// 为什么0xc000016058 64bit地址只有40bit
	fmt.Printf("%p %p \n", &cat, &str)
}

func readPointer() {
	fmt.Println()
	house := "Malibu Point 10880, 90265"

	ptr := &house

	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr address: %p\n", ptr)
	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)
}

func swapByPointer(a *int, b *int) {
	fmt.Println()
	t := *a

	*a = *b

	*b = t
}

func swapPointer(a, b *int) {
	fmt.Println()
	fmt.Println("before swap", a, b)
	fmt.Println("before swap value", *a, *b)
	a, b = b, a
	fmt.Println("after swap", a, b)
	fmt.Println("after swap value", *a, *b)
}
/**
指针获取命令行输入
 */
func flagParse() {
	fmt.Println()
	var mode = flag.String("mode", "", "process mode")
	flag.Parse()
	fmt.Println(*mode)
}



func main() {
	checkVirtualAddress()

	readPointer()

	x, y := 1, 2
	swapByPointer(&x, &y)

	fmt.Println(x, y)

	swapPointer(&x, &y)

	flagParse()

}
