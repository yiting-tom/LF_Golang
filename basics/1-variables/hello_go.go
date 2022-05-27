package main

import (
	"fmt"
)

func main() {
	/*********** var v.s. := ************
	 *
	 * var
	 *   1. inside & ouside the function scope.
	 *   2. can separatly assign.
	 *
	 * :=
	 *   1. only inside the fucntion scope.
	 *   2. can't separatly assign.
	 *
	 *************************************/
	var v1 string = "var v1 string ="
	var v2 = "var v2 ="
	v3 := "v2 :="

	// separatly claim.
	var str1 string  // default: ""
	var int2 int     // default: 0
	var boo3 bool    // default: false
	var flo4 float32 // default: .0

	// multiple variable define.
	var m1, m2, m3 int = 1, 2, 3
	var n1, n2, n3 = "str", 1, false
	o1, o2, o3 := "str", 0, true

	// multiple variable define in block.
	var (
		a  int
		_b string = "string"
	)

	// const read-only variable.
	const PI float32 = 3.1415926
	// PI = 123  Error (declared const)

	fmt.Println(v1, v2, v3)
	fmt.Println(str1, int2, boo3, flo4)
	fmt.Println(m1, m2, m3)
	fmt.Println(n1, n2, n3)
	fmt.Println(o1, o2, o3)
	fmt.Printf("a has value: %v and type: %T\n", a, a)

	// Warning %!T (Missing)
	fmt.Printf("_b has value: %v and type: %T\n", _b, _b)
	fmt.Println(PI)
}
