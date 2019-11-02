package __3

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(" int8 range", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range", math.MinInt64, math.MaxInt64)

	fmt.Println(" uint8 range", 0, math.MaxUint8)
	fmt.Println("uint16 range", 0, math.MaxUint16)
	fmt.Println("uint32 rangfe", 0, math.MaxUint32)
	fmt.Printf("uint64 range 0 %d \n", math.MaxUint64)

	var (
		a int32 = 1047483647
		b = int16(a)
		c = float32(math.Pi)
	)

	// 32 -> 16进制，高位16位被隐式截断
	fmt.Printf("int32: 0x%x %d \n", a, a)
	fmt.Printf("int16: 0x%x %d \n", b, b)
	fmt.Printf("float32 to int: %d \n", int(c))
	//fmt.Println("uint64 range", 0, math.MaxUint64)
}
