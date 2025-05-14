package main

import "fmt"

func fact() {
	var n int
	fmt.Println("Enter an Positive Integer: ")
	fmt.Scanln(&n)
	for n < 0 {
		fmt.Println("Sorry, only positive numbers, enter again: ")
		fmt.Scanln(&n)
	}
	if n == 0 {
		print("The factorial of 0 is 1")
	} else {
		var f int = 1
		for i := 0; i < n+1; i++ {
			f *= i
		}
		fmt.Printf("The factorial of %d is %d", n, f)
	}

}

func main() {
	fmt.Println("Hello")
}
