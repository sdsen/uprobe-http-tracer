package main

import "net/http"

func main() {
	urls := []string{"http://google.com", "http://reddit.com", "http://pastebin.com/tools"}
	for _, url := range urls {
		http.Get(url)
	}
}


