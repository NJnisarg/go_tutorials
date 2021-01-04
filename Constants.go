package main

import "fmt"

// iota is scoped to the block of variables
const (
	i = iota
	j = iota
	k = iota
)

// iota counter resets
// _ is a write only variable. Hence you can not read it, but you can assign to it to throw away a value
// You can manipulate iota with a constant expression
const (
	_ = iota + 7
	i2
	j2
)

// Also useful to make pattern based constants
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	const myConst int = 4
	fmt.Printf("%v, %T\n", myConst, myConst)

	// We can't set the const from an expression that needs to be evaluated at runtime
	// const a int = math.Sin(4.5)

	// If you dont declare the data type of the const, it will do implicit type conversion
	const a = 7 // var a = 7 a+b wont work since implicit conversion is not done
	const b int16 = 11

	fmt.Printf("%v, %T\n", a+b, a+b)

	// Enumerated constants
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)

	fmt.Printf("%v %T\n", i2, i2)
	fmt.Printf("%v %T\n", j2, j2)

	fmt.Printf("%v %T\n", KB, KB)
	fmt.Printf("%v %T\n", MB, MB)
	fmt.Printf("%v %T\n", GB, GB)

}
