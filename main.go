package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	fmt.Println("Site Status Checker")
	var wg sync.WaitGroup
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Printf("No sites submitted, exiting")
		os.Exit(1)
	}
	for _, site := range arguments {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			resp, err := http.Get(site)
			if err != nil {
				fmt.Printf("[DOWN] %s\n", site)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				fmt.Printf("[UP] %s\n", site)
			} else {
				fmt.Printf("[DOWN] %s\n", site)
			}
		}(site)
	}
	wg.Wait()
	fmt.Printf("test")
}
