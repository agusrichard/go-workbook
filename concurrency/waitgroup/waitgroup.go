package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"
// )

// var urls = []string{
// 	"https://google.com",
// 	"https://tutorialedge.net",
// 	"https://twitter.com",
// }

// func fetch(url string, wg *sync.WaitGroup) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	wg.Done()
// 	fmt.Println(resp.Status)
// 	return resp.Status, nil
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("HomePage Endpoint Hit")
// 	var wg sync.WaitGroup

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go fetch(url, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Returning Response")
// 	fmt.Fprintf(w, "Responses")
// }

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

// func main() {
// 	handleRequests()
// }

// func main() {
// 	fmt.Println("Working on WaitGroup")

// 	var wg sync.WaitGroup
// 	start := time.Now()
// 	defer func() {
// 		dur := time.Since(start)
// 		fmt.Println("Dur", dur)
// 	}()

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func() {
// 			dur := time.Duration(rand.Intn(1000)) * time.Millisecond
// 			time.Sleep(dur)
// 			fmt.Println("Hello", dur)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Done on WaitGroup")
// }

// func main() {
// 	fmt.Println("Working on WaitGroup")

// 	var wg sync.WaitGroup
// 	start := time.Now()
// 	defer func() {
// 		dur := time.Since(start)
// 		fmt.Println("Dur", dur)
// 	}()

// 	num := make(chan int, 100)
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			fmt.Println("Sending a num")
// 			randInt := rand.Intn(1000)
// 			dur := time.Duration(randInt) * time.Millisecond
// 			time.Sleep(dur)

// 			num <- i
// 			fmt.Println("Done sending a num")
// 		}(i)
// 	}

// 	var nums []int
// 	go func() {
// 		for n := range num {
// 			wg.Done()
// 			nums = append(nums, n)
// 		}
// 	}()

// 	wg.Wait()

// 	fmt.Println("nums", nums)
// 	fmt.Println("Done on WaitGroup")
// }

type Response struct {
	name       string
	age        int
	ageDoubled int
	nums       []int
}

func main() {
	fmt.Println("Start working on WaitGroup")
	start := time.Now()
	defer func() {
		fmt.Println("Done working on WaitGroup")
		fmt.Println("Duration", time.Since(start))
	}()

	name := make(chan string)
	defer close(name)
	go func(name chan<- string) {
		fmt.Println("Sending a name")
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(dur)
		name <- "John"
		fmt.Println("Done sending a name")
	}(name)

	age := make(chan int)
	ageDoubled := make(chan int)
	defer close(age)
	go func(age chan<- int) {
		fmt.Println("Sending a age")
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(dur)
		age <- 23
		ageDoubled <- 23 * 2
		fmt.Println("Done sending a age")
	}(age)

	nums := make(chan []int)
	defer close(nums)
	go func(nums chan<- []int) {
		var list []int
		var wg sync.WaitGroup

		num := make(chan int, 100)
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int, wg *sync.WaitGroup) {
				fmt.Println("Sending a num")
				dur := time.Duration(rand.Intn(5000)) * time.Millisecond
				time.Sleep(dur)
				num <- i
				fmt.Println("Done sending a num", i)
			}(i, &wg)
		}

		go func(nums chan<- []int) {
			for n := range num {
				list = append(list, n)
				wg.Done()
			}
		}(nums)

		wg.Wait()

		nums <- list
	}(nums)

	response := Response{
		name:       <-name,
		age:        <-age,
		ageDoubled: <-ageDoubled,
		nums:       <-nums,
	}

	fmt.Printf("Response: %+v\n", response)
}
