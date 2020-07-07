package main

import "fmt"

func main() {
	m := make(map[string]string)
	fmt.Println(m)

	m["sekar"] = "Sekardayu Hana Pradiani"
	fmt.Println(m)
	
	m["sekar"] = "My Girlfriend"
	fmt.Println(m)
	
	delete(m, "sekar")
	fmt.Println(m)
	
	v, ok := m["sekar"]
	fmt.Println(v, ok)
}
