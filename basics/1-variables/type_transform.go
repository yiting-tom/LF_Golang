package main

import (
	"fmt"
)

/**
Transform other data type of a variable to string type.

1. fmt.Sprintf(): return a string with printf format
2. strconv module:
*/
func main() {
	int1 := 1
	flo2 := 2.0
	boo := false
	char := 'd'
	var str string

	str = fmt.Sprintf("%v", int1)
	fmt.Printf("int1 type %T str=%q\n", str, str)

	str = fmt.Sprintf("%v", flo2)
	fmt.Printf("flo2 type %T str=%q\n", str, str)

	str = fmt.Sprintf("%v", boo)
	fmt.Printf("boo type %T str=%q\n", str, str)

	str = fmt.Sprintf("%v", char)
	fmt.Printf("char type %T str=%q\n", str, str)
}
