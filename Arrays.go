package main

import "fmt"

func main() {

	// Array initialization syntax
	var students = [3]string{"Nisarg", "Yash", "Manan"}
	teachers := [...]string{"Annappa", "Mohit", "BRC"}
	var facAd [2]string
	facAd[0] = "Manu"
	facAd[1] = "Soumya"

	fmt.Println(students)
	fmt.Printf("%v, %T\n", teachers, teachers)
	fmt.Println(facAd)
	fmt.Println(len(teachers))

	// 2D matrix
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// Deep copy, not reference assignment(not pointer)
	a := [3]int{0, 0, 1}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// Slices. As the name suggests, they form a slice of an array generally
	// Slices work like references(pointers), hence they refer to a section of an array.
	c := []int{0, 0, 1} // Not specifying the size here means slice
	d := c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	// other syntax to create slice
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b1 := a1[:] // All elements
	c1 := a1[3:]
	d1 := a1[:6]
	e1 := a1[3:6]
	a1[5] = 100 // All slices point to same memory
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)
	fmt.Println(len(d1)) // Len is the actual len and cap is the capacity of the slice
	fmt.Println(cap(d1))

	// make syntax
	a2 := make([]int, 3, 100) // datatype, len and cap
	fmt.Println(a2)
	fmt.Println(len(a2))
	fmt.Println(cap(a2))

	a2 = append(a2, 1)
	fmt.Println(a2)
	fmt.Println(len(a2))
	fmt.Println(cap(a2))

	b2 := []int{7, 8}
	a2 = append(a2, 1, 2, 3)
	a2 = append(a2, []int{4, 5}...) // Golang spread operator
	a2 = append(a2, b2...)
	fmt.Println(a2)
	fmt.Println(len(a2))
	fmt.Println(cap(a2))

}
