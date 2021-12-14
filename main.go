package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func goroutine1(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine 1")
	waitgroup.Done()
}

func goroutine2(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine 2")
	waitgroup.Done()
}

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("HTTP Error: %s", err)
		return "", err
	}
	wg.Done()
	log.Println(resp.Request.URL, resp.Status)
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
	fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	// fmt.Println("Execute first")

	// var waitgroup sync.WaitGroup
	// waitgroup.Add(3)
	// go func() {
	// 	fmt.Println("Inside my anonymous goroutine")
	// 	waitgroup.Done()
	// }()
	// go goroutine1(&waitgroup)
	// go goroutine2(&waitgroup)
	// waitgroup.Wait()

	// fmt.Println("Execute last")

	handleRequests()
}
