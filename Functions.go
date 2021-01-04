package main

import "fmt"

// Greeting is a struct we will use to demo methods in go
type Greeting struct {
	name string
	age  int
}

func main() {
	test1()
	Test2()

	func(name string) {
		fmt.Println("Hello ", name)
	}("Nisarg")

	var f func(string, int) (float64, float32, error) = func(name string, age int) (float64, float32, error) {
		fmt.Println("This is a function as a object/datatype. You can pass it around like a variable")
		fmt.Println(name, age)
		return 4.5, 3.2, nil
	}

	test3("Nisarg", 1, "string1", "string2", "string3", f, 1, 2, 3)

	// Using functions as methods on structs
	var g *Greeting = new(Greeting)
	g.name = "Nisarg"
	g.age = 21
	g.greet()
}

func test1() {
	fmt.Println("This is a test function")
}

// Test2 is a exported function
func Test2() {
	fmt.Println("This function has first letter uppercase. It will be exported to package scope")
}

func test3(name string, val int, v1, v2, v3 string, callback func(string, int) (float64, float32, error), values ...int) {
	fmt.Printf("%T\n", values)
	fmt.Println(name, val, v1, v2, v3, values)
	callback(name, val)
}

func (g *Greeting) greet() {
	fmt.Println("This is a method on a Greeting struct pointer")
	fmt.Println(g.name, g.age)
}
