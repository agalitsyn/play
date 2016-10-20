// https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, results chan string, fetchedURLs map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 || fetchedURLs[url] {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	fetchedURLs[url] = true
	if err != nil {
		results <- err.Error()
		return
	}

	results <- fmt.Sprintf("found: %s %q", url, body)
	for _, u := range urls {
		wg.Add(1)
		Crawl(u, depth-1, fetcher, results, fetchedURLs, wg)
	}

	return
}

func main() {
	results := make(chan string)
	// https://golang.org/pkg/sync/#WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher, results, make(map[string]bool), &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}