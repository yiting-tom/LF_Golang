package main

import (
	"fmt"
	"unsafe"
)

/**
 */

func main() {
	var defa bool // default value is false
	var tru = true
	var fal = false // default type is float64

	fmt.Printf("default: %v\ttrue: %v\tfalse: %v\n", defa, tru, fal)

	fmt.Println("\n=========== Size (bytes) ===========") // The data type of size isn't dependent on the platform.
	fmt.Printf("default: %v\ttrue: %v\tfalse: %v\n", unsafe.Sizeof(defa), unsafe.Sizeof(tru), unsafe.Sizeof(fal))
}
