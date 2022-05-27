package main

import (
	"fmt"
	"unsafe"
)

/**
1. charactor is computable
2. if it's in ASCII -> byte type
3. if not -> int type
*/

func main() {
	var defa byte // default value is ""
	var one byte = 'a'
	var two = 'b' // default type is float64

	// charactors not in ASCII is not allowed stored in byte type
	// var int_special byte = '😎'

	var int_special int = '😎' // but it can be stored in int type (remember: print out the charactor need to use %c)

	fmt.Printf("default: %v\tone: %v\ttwo: %v\tint_special: %c\n", defa, one, two, int_special)

	fmt.Println("\n=========== Size (bytes) ===========") // The data type of size isn't dependent on the platform.
	fmt.Printf("default: %v\tone: %v\ttwo: %v\tint_special: %v\n", unsafe.Sizeof(defa), unsafe.Sizeof(one), unsafe.Sizeof(two), unsafe.Sizeof(int_special))
}
