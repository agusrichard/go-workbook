package main

// import (
// 	"fmt"
// 	"sync"
// )

// func myfunc(waitgroup *sync.WaitGroup) {
// 	fmt.Println("Inside the goroutine")
// 	waitgroup.Done()
// }

// func main() {
// 	fmt.Println("Hello World")

// 	var waitgroup sync.WaitGroup
// 	waitgroup.Add(1)
// 	go myfunc(&waitgroup)
// 	waitgroup.Wait()

// 	fmt.Println("Finished execution")
// }

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
