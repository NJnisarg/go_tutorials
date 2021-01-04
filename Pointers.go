package main

import "fmt"

// MyStruct is a struct in go
type MyStruct struct {
	name string
	age  int
}

func main() {

	// Basics of pointers.
	// We cannot do pointer arithematic in golang
	var a int = 56
	var b *int = &a
	fmt.Println(a, *b)
	*b = 7
	fmt.Println(a, *b)

	// Assigning pointers for structs directly
	var ms *MyStruct
	ms = &MyStruct{
		name: "nisarg",
	}

	fmt.Printf("%T, %v\n", *ms, ms)

	// Using the new keyword
	var ms2 *MyStruct = new(MyStruct)

	// Both syntax give the same result
	// We don't have to use ms2->age like C++
	(*ms2).name = "nisarg"
	ms2.age = 21

	fmt.Println(*ms2)
}
