package main

import "fmt"

func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func recursiveSummation(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + recursiveSummation(nums[1:])
}

func main() {
	// Nonfunctional programming - Updating string
	name := "Sherlock"
	name = name + " Holmes"
	fmt.Println(name)

	// Functional programming - Updating string
	firstname := "Sherlock"
	lastname := "Holmes"
	fullname := firstname + " " + lastname // Or we could use string formatting
	fmt.Println(fullname)

	// Nonfunctional programming - Avoid updating arrays
	friends1 := [3]string{"Joe", "John"}
	friends1[2] = "James"
	fmt.Println(friends1)

	// Functional programming - Avoid updating arrays
	friends2 := []string{"Joe", "John"}
	friends2 = append(friends2, "James")
	fmt.Println(friends2)

	// Nonfunctional programming - Avoid updating maps
	fruits1 := map[string]int{"Orange": 1}
	fruits1["Banana"] = 2
	fmt.Println(fruits1)

	// Functional programming - Avoid updating maps
	fruits2 := map[string]int{"Orange": 2}
	newFruits := map[string]int{"Banana": 3}

	allFruits := make(map[string]int, len(fruits2)+len(newFruits))

	for key, val := range fruits2 {
		allFruits[key] = val
	}

	for key, val := range newFruits {
		allFruits[key] = val
	}
	fmt.Println(allFruits)

	// Currying example
	add10 := add(10)
	add20 := add10(20)
	fmt.Println(add20)

	// Recursion example - eliminating the usage of loops
	fmt.Println(factorial(3))
	fmt.Println(recursiveSummation([]int{1, 2, 3, 4, 5}))
}
