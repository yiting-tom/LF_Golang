package main

import (
	"fmt"
	"strconv"
)

// Ref: https://pkg.go.dev/strconv

func main() {
	one := 1
	two := 2.0
	boo := true
	var str string

	// convert to int
	// if error occurs, return 0
	i, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("Atoi error:", err)
	} else {
		fmt.Println("Atoi success:", i)
	}

	// Convert to int64
	str = strconv.Itoa(one)
	fmt.Printf("str type %T str=%q\n", str, str)

	fmt.Println(
		"\n/*********************************",
		"	others convert to string",
		"*********************************/",
	)
	// transform to int64 is needed before convert to string.
	// 10: base 10
	str = strconv.FormatInt(int64(one), 10)
	str = strconv.FormatUint(uint64(one), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	// 'f': foramt is float could be like 'f', 'e', 'E'
	// 3: precision is 3
	// 64: bit size of float64 is 64
	str = strconv.FormatFloat(two, 'f', 3, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(boo)
	fmt.Printf("str type %T str=%q\n", str, str)

	fmt.Println(
		"\n/*********************************",
		"	string convert to others",
		"*********************************/",
	)
	i64, err := strconv.ParseInt("123", 10, 64) // base 10, bit size of int64 is 64
	fmt.Printf("i64 type %T i64=%v\n", i64, i64)

	ui64, err := strconv.ParseUint("123", 10, 64) // base 10, bit size of uint64 is 64
	fmt.Printf("ui64 type %T ui64=%v\n", ui64, ui64)

	f32, err := strconv.ParseFloat("1.23", 32) // base 10, bit size of float32 is 32
	fmt.Printf("f32 type %T f32=%v\n", f32, f32)

	tru, err := strconv.ParseBool("true")
	fmt.Printf("boo type %T boo=%v\n", tru, tru)

	fmt.Println(
		"\n/*********************************",
		"	others formatting",
		"*********************************/",
	)
	fone := strconv.FormatInt(int64(one), 2)
	fmt.Printf("fone type %T fone=%q\n", fone, fone)

	ftwo := strconv.FormatFloat(two, 'e', -1, 32)
	fmt.Printf("ftwo type %T ftwo=%q\n", ftwo, ftwo)

	fmt.Println(
		"\n/*********************************",
		"	string formatting",
		"*********************************/",
	)
	str = strconv.Quote("哈囉，世界")
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.QuoteRune('😎')
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.QuoteToASCII("哈囉，世界")
	fmt.Printf("str type %T str=%q\n", str, str)
}
