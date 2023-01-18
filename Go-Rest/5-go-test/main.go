package main

import (
	"fmt"
	"sync"
)

//original output

// rahiljaiswal@Rahils-MacBook-Pro 5-go-test % go run main.go
// found: https://golang.org/ "The Go Programming Language"
// found: https://golang.org/pkg/ "Packages"
// found: https://golang.org/ "The Go Programming Language"
// found: https://golang.org/pkg/ "Packages"
// not found: https://golang.org/cmd/
// not found: https://golang.org/cmd/
// found: https://golang.org/pkg/fmt/ "Package fmt"
// found: https://golang.org/ "The Go Programming Language"
// found: https://golang.org/pkg/ "Packages"
// found: https://golang.org/pkg/os/ "Package os"
// found: https://golang.org/ "The Go Programming Language"
// found: https://golang.org/pkg/ "Packages"
// not found: https://golang.org/cmd/

var visited map[string]bool

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var mu sync.Mutex
var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	// if depth <= 0 {
	// 	return
	// }
	defer wg.Done()
	_, v := visited[url]
	if v {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	mu.Lock()
	visited[url] = true
	mu.Unlock()

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, fetcher)
		//wg.Wait()
	}
	return
}

func main() {
	visited = make(map[string]bool)
	wg.Add(1)
	Crawl("https://golang.org/", fetcher)
	wg.Wait()
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}



////


///
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	doSomething(s)
	fmt.Println(s)
}
func doSomething(s []int) {
	s[0] = 99
	fmt.Println(s)
}

// [ 1, 2, 3]
// [ 99, 2, 3]
// [ 1, 2, 3]

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup            //waitgroup
	strings := make([]string, 0, 50) //
	for i := 0; i < 50; i++ {
		strings = append(strings, fmt.Sprint("string", i))
	}

	wg.Add(len(strings))
	for _, str := range strings {
		//fmt.Println("Str >", str)
		go func(st string) {
			defer wg.Done()
			fmt.Println("Inside Anony >", str)
		}(str)
	}
	wg.Wait()
}
