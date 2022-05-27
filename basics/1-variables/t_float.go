package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var defa float32 // default value is 0 for float32 and float64
	var one float32 = 1.0
	var two = 2.0 // default type is float64

	var f32 float32 // 4 bytes: -3.4e38 to 3.4e38
	var f64 float64 // 8 bytes: -1.8e308 to 1.8e308

	fmt.Printf("default: %v\tone: %v\ttwo: %v\n", defa, one, two)

	fmt.Println("\n=========== Size (bytes) ===========") // The data type of size isn't dependent on the platform.
	fmt.Printf("default: %v\tone: %v\ttwo: %v\n", unsafe.Sizeof(defa), unsafe.Sizeof(one), unsafe.Sizeof(two))
	fmt.Printf("f32: %v\tf64: %v\n", unsafe.Sizeof(f32), unsafe.Sizeof(f64))
}
