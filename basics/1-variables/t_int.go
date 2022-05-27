package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var defa int // default value is 0 and type is int
	var one int = 1
	var two = 2

	var i8 int8   // 1 byte: -128 to 127
	var i16 int16 // 2 bytes: -32768 to 32767
	var i32 int32 // 4 bytes: -2147483648 to 2147483647
	var i64 int64 // 8 bytes: -9223372036854775808 to 9223372036854775807

	var ui8 uint8   // 1 byte: 0 to 255
	var ui16 uint16 // 2 bytes: 0 to 65535
	var ui32 uint32 // 4 bytes: 0 to 4294967295
	var ui64 uint64 // 8 bytes: 0 to 18446744073709551615

	fmt.Printf("default: %v\tone: %v\ttwo: %v\n", defa, one, two)

	fmt.Println("\n=========== Size (bytes) ===========") // The data type of size is dependent on the platform.
	fmt.Printf("default: %v\tone: %v\ttwo: %v\n", unsafe.Sizeof(defa), unsafe.Sizeof(one), unsafe.Sizeof(two))
	fmt.Printf("i8: %v\ti16: %v\ti32: %v\ti64: %v\n", unsafe.Sizeof(i8), unsafe.Sizeof(i16), unsafe.Sizeof(i32), unsafe.Sizeof(i64))
	fmt.Printf("ui8: %v\tui16: %v\tui32: %v\tui64: %v\n", unsafe.Sizeof(ui8), unsafe.Sizeof(ui16), unsafe.Sizeof(ui32), unsafe.Sizeof(ui64))
}
