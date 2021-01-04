package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true")
	}

	// Stmt initializer syntax. First part is an initializer.
	// The variables generated in the first half are scoped only to the if block
	if a := true; a {
		fmt.Println("a is available only in this block scope")
	}

	// You have to write else on the same line as the closing brace of if section
	if 4 > 5 {
		fmt.Println("yeah true")
	} else {
		fmt.Println("This is false")
	}

	if 4 > 5 {
		// something here
	} else if 4 >= 6 {
		// something here
	} else {
		// something here
	}

	// Switch stmt

	// Using conditional stmt in case block
	i := 10

	switch {
	case i >= 10:
		fmt.Println("We can also use a variable and conditional expr in the case section")
	}

	// Initializer syntax
	switch a := 20; a {
	case 20:
		fmt.Println("You can also use the initializer syntax like if stmt")
	}

	// Normal syntax with mutiple matches allowed
	switch 6 {
	case 1, 2, 3:
		fmt.Println("Can match multiple cases with or operation")
		// break is not needed, it is implied
		fmt.Println("Can print multiple lines")
		break
		fmt.Println("break will get you out of the entire switch block")
	case 4, 5, 6:
		fmt.Println("Reached case 4,5,6")
		fallthrough
	case 7, 8, 9:
		fmt.Println("You can reach here by fall through which will not evaluate anything")
	default:
		fmt.Println("If nothing works, it will reach here")
	}

	// Type Switch
	var v interface{} = 10
	switch v.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("This is a string")
	default:
		fmt.Println("We dont know the type yet")
	}

}
