package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	//Exercise 1:
	factorialIterative := func(n int) uint64 {
		if n < 0 {
			fmt.Println("Factorial of negative number doesn't exist.")
			return 0 // Or handle error appropriately
		}
		var result uint64 = 1
		for i := 1; i <= n; i++ {
			result *= uint64(i)
		}
		return result
	}
	num := 5
	fmt.Printf("Factorial of %d (Iterative): %d\n", num, factorialIterative(num))
	num = 0
	fmt.Printf("Factorial of %d (Iterative): %d\n", num, factorialIterative(num))
	num = -3
	fmt.Printf("Factorial of %d (Iterative): %d\n", num, factorialIterative(num))

	//Exercise 2:
	findlargest := func(arr []int) int {
		if len(arr) == 0 {
			fmt.Println("Array is empty.")
			return 0 // Or handle error appropriately
		}
		largest := arr[0]
		for _, value := range arr {
			if value > largest {
				largest = value
			}
		}
		return largest
	}
	numbers := []int{3, 5, 7, 2, 8}
	largest := findlargest(numbers)
	fmt.Println("Largest number in the array:", largest)
}

// Exercise 3:
