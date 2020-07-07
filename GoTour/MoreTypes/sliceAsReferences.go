package main

import "fmt"

func main() {
	beatles := [4]string{
		"John Lennon",
		"Paul McCartney",
		"George Harrison",
		"Ringo Starr",
	}
	
	fmt.Println(beatles)
	bestDuo := beatles[0:2]
	fmt.Println(bestDuo)
	bestDuo[0] = "John"
	fmt.Println(bestDuo)
	fmt.Println(beatles)
}