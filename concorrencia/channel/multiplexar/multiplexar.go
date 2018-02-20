package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {

	in := gen(
		"https://www.youtube.com/",
		"https://golang.org/",
		"https://www.bing.com/",
		"https://www.google.com.br",
	)

	//FAN OUT
	ic1 := title(in)
	ic2 := title(in)

	//FAN IN
	for n := range merge(ic1, ic2) {
		fmt.Println(n)
	}
}

func gen(items ...string) chan string {
	out := make(chan string)
	go func() {
		for _, i := range items {
			out <- i
		}
		close(out)
	}()
	return out
}

func title(item chan string) chan string {
	out := make(chan string)
	go func() {
		for i := range item {
			resp, _ := http.Get(i)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			out <- r.FindStringSubmatch(string(html))[1]
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	/*
		Start an output goroutine for each input channel in cs.
		Output copies values from c to out until c is closed, then calls wg.Done.
	*/
	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	/*
		Start a goroutine to close out once all the output goroutines are done.
		This must start after the wg.Add call.
	*/
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
