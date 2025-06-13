package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Error      error
}

func checkURL(url string) Result {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	return Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   duration,
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("надо больше урлов ^_^")
		os.Exit(1)
	}

	fmt.Println("URL | Status Code | Response Time")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	for _, url := range os.Args[1:] {
		result := checkURL(url)
		if result.Error != nil {
			fmt.Printf("%s | ERROR: %v\n", result.URL, result.Error)
			continue
		}
		fmt.Printf("%s | %d | %v\n", result.URL, result.StatusCode, result.Duration)
	}
}
