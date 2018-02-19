/*
	Google I/O 2012 - Go Concurrency Patterns
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1] // index [1] é o valor
		}(url) // passando a url como parametro
	}

	return c
}

func main() {
	t1 := titulo("https://www.youtube.com/", "https://golang.org/")
	t2 := titulo("https://www.google.com.br", "https://www.bing.com/")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
