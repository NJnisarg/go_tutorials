package main

import (
	"fmt"
	"strconv"
)

// Lower case name means its scoped to the package
var l float64 = 46

// L is a Upper case name means global access outside package as well. Exported
var L float32 = 77

// A block of variables. Useful for conciseness and logical grouping. Movie is a global exported variable here.
var (
	actorName string = "Elisabeth Sladen"
	companion string = "Sarah Jane Smith"
	Movie     string = "Movie name"
)

func main() {

	// All of the variables here are block scoped (or local to the block of code)

	// Declare and Assign Later
	var i int
	i = 42

	// Declare and Assign together
	var j float64 = 23

	// Short Declaration without data type specified. Default datatype is inferred
	k := 45
	var p = 7

	// Type Conversion. No auto type conversion, we have to explicitly convert
	var m = float32(i)
	var s = string(i) // This will convert i into a string of bytes. Hence we will get the unicode repr
	var s2 = strconv.Itoa(i)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf(" %v, %T\n", m, m)
	fmt.Printf(" %v, %T\n", s, s)
	fmt.Printf(" %v, %T\n", s2, s2)
	fmt.Println(actorName)
	fmt.Println(p)
}
