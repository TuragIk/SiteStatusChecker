package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type SiteResult struct {
	URL     string
	Status  string // "UP" or "DOWN"
	Latency time.Duration
}

func checkSite(site string) string {
	resp, err := http.Get(site)
	status := ""
	if err != nil || resp == nil {
		status = "DOWN"
		return status
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		status = "UP"
	} else {
		status = "DOWN"
	}
	return status
}

func worker(id int, jobs <-chan string, results chan<- SiteResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for site := range jobs {
		start := time.Now()
		status := checkSite(site)
		elapsed := time.Since(start)
		results <- SiteResult{site, status, elapsed}
	}
}
func main() {
	fmt.Println("Site Status Checker")
	var wg sync.WaitGroup
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Printf("No sites submitted, exiting.\n")
		os.Exit(1)
	}

	numJobs := len(arguments)
	jobs := make(chan string, numJobs)
	results := make(chan SiteResult, numJobs)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	for _, site := range arguments {
		jobs <- site
	}
	close(jobs)
	wg.Wait()
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("[%s] %s (%s)\n", result.Status, result.URL, result.Latency)
	}
}
