package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"

	"github.com/yhat/scrape"
)

const (
	ua = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://whatsmyuseragent.com/", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	// resp, err := http.Get("http://whatsmyuseragent.com/")
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	info := scrape.ByClass("info")
	data := scrape.FindAll(root, info)
	for _, v := range data {
		fmt.Printf("%s\n", scrape.Text(v))
	}
}
