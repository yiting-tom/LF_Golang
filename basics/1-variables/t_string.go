package main

import (
	"fmt"
	"unsafe"
)

/**
1. 2 representations, "" for formated string or `` for raw output
2. string is a sequence of bytes
3. string is immutable
4. can be +ed together
*/

func main() {
	var defa string // default value is ""
	var str1 string = "string1"
	var str2 = `string2` // default type is float64

	// 3. string is immutable
	// str1[0] = "H" // error: cannot assign to string index 0

	// 4. can be +ed together
	str3 := str1 + str2 + // multiple lines concatination are allowed
		str1 + str2

	fmt.Printf("default: %v\tstr1: %v\tstr2: %v\tstr3: %v\n", defa, str1, str2, str3)

	fmt.Println("\n=========== len (bytes) ===========")
	fmt.Printf("default: %v\tstring1: %v\tstring2: %v\n", len(defa), len(str1), len(str2))

	fmt.Println("\n=========== Size (bytes) ===========")
	fmt.Printf("default: %v\tstring1: %v\tstring2: %v\n", unsafe.Sizeof(defa), unsafe.Sizeof(str1), unsafe.Sizeof(str2))
}
