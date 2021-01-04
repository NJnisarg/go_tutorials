package main

import "fmt"

func main() {

	// 1.) Boolean

	var val bool // Uninitialized in go means it will be zero, null, false and so on
	m := (5 < 6)
	fmt.Println(val)
	fmt.Println(m)

	// 2.) Numeric
	// Signed ints ==> int8, int16, int32, int64, int (default) (might be int32 or int64 based on the platform)
	// Unsigned ints ==> Same as above just prefixed with 'u' ==> uint8, uint16, .... so on
	var i int64 = 9823439873433
	var j uint32 = 234
	var k uint32 = 3454
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println(k / j) // Integer division

	// i/j is not allowed since operations on non matching data types is not allowed
	fmt.Println(uint32(i) / j)

	n := 3.15 // Defaults to float64
	n = 2.1e14
	n = 4e10

	var f float32 = 45.78

	fmt.Println(n)
	fmt.Println(f + f)

	var c1 complex64 = 1 + 2i
	var c2 complex128 = complex(4, 5)

	fmt.Printf("%v %T", real(c1), imag(c2))

	// 3.) String
	// A string is a array of bytes. It is immutable. A byte in golang is uint8
	var s string = "this is a string"
	fmt.Println(s)

	b := []byte(s)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", b, b)

	// 4.) Rune
	// A rune is a uint32 used to represent a utf32 character.
	var r rune = 't'
	fmt.Printf("%v, %T\n", r, r)
}
