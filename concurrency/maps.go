package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	mymap["sekar"] = 23
	mymap["saskia"] = 21
	mymap["arifa"] = 23

	var myKey []string
	for key, value := range mymap {
		fmt.Println(key, value)
		myKey = append(myKey, key)
	}

	fmt.Println(myKey)

	// deleting key from map
	mymap["jane"] = 22
	delete(mymap, "jane")
	fmt.Println(mymap["jane"])
}
