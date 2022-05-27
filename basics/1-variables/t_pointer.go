package main

import "fmt"

/**
value data type:
	int, float, bool, `string`, `array`, `struct`
reference data type:
	pointer, slice, map, chan, interface
*/

func main() {
	var i int = 10
	var i_ptr *int = &i

	// error: cannot use i (type int) as type *int
	// var i_ptr int = &i

	// error: cannot use i_ptr (type *float32) as type *int
	// var i_ptr *float32 = &i

	fmt.Printf("i point to i_ptr: %p\n", i_ptr)
	fmt.Printf("i_ptr store the value of i: %v\n", *i_ptr)

}
