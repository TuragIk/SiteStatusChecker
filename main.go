package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type SiteResult struct {
	URL     string
	Status  string // "UP" or "DOWN"
	Latency time.Duration
}

func checkSite(site string, timeout *int) string {
	client := http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}
	resp, err := client.Get(site)
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

func readFile(filename *string, sites []string) []string {
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error reading file.\n")
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading file.\n")
			os.Exit(1)
		}
		for _, site := range record {
			sites = append(sites, site)
		}
	}
	return sites
}

func worker(id int, jobs <-chan string, results chan<- SiteResult, wg *sync.WaitGroup, timeout *int) {
	defer wg.Done()
	for site := range jobs {
		start := time.Now()
		status := checkSite(site, timeout)
		elapsed := time.Since(start)
		results <- SiteResult{site, status, elapsed}
	}
}
func main() {
	fmt.Println("Site Status Checker")
	var wg sync.WaitGroup
	filename := flag.String("f", "", "File path to read URLs from")
	timeout := flag.Int("t", 5, "Max timeout (seconds) for checking site. Default: 5.")
	flag.Parse()
	var sites []string
	if *filename != "" {
		sites = readFile(filename, sites)
	}
	sites = append(sites, flag.Args()...)

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		content, err := io.ReadAll(os.Stdin)
		if err == nil {
			sites = append(sites, strings.Fields(string(content))...)
		} else {
			fmt.Printf("Error reading piped data.\n")
		}
	}

	if len(sites) == 0 {
		fmt.Printf("No sites submitted. Exiting.\n")
		os.Exit(1)
	}

	numJobs := len(sites)
	jobs := make(chan string, numJobs)
	results := make(chan SiteResult, numJobs)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg, timeout)
	}
	for _, site := range sites {
		jobs <- site
	}
	close(jobs)
	wg.Wait()
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("[%s] %s (%s)\n", result.Status, result.URL, result.Latency)
	}
}
