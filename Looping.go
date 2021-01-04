package main

import "fmt"

func main() {

	// Normal loop syntax
	for i := 2; i <= 7; i += 2 {
		fmt.Println(i)
	}

	// Multiple initializer and multiple increment
	i, j := 5, 6                     // We can put this in the first stmt too. This will define the variable to outer scope.
	for ; i <= 10; i, j = i+1, j+1 { // Do not forget the semi-colon if not passing the initializer
		fmt.Println(i, j)
	}

	// You can leave out the semi-colon only if both the init and increment are absent.
	// It basically becomes a while loop
	i = 0
	for i < 5 {
		i++
		if i < 3 {
			continue
		}
		if i > 3 {
			break
		}

		fmt.Println(i)

	}

	// We can label the loop block to break out of it
Loop:
	for i := 0; i < 7; i += 2 {
		for j := 0; j < 3; j++ {
			fmt.Println(i * j)
			if i*j > 4 {
				break Loop
			}
		}
	}

	// Range loop. Looping over iterables
	fmt.Println("Iterable looping")
	var medianAge map[string]int = map[string]int{
		"baroda":    25,
		"mumbai":    22,
		"bangalore": 26,
	}
	var s []int = []int{1, 2, 4}
	for k, v := range s {
		fmt.Println(k, v)
	}

	for k, v := range medianAge {
		fmt.Println(k, v)
	}
}
