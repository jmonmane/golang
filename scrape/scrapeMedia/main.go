package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

const (
	ua = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36"
)

func login() {
	client := &http.Client{}
	// POST /ajax/login.html HTTP/1.1
	// Host: learn.infiniteskills.com
	// Connection: close
	// Content-Length: 64
	// Cache-Control: max-age=0
	// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
	// Origin: https://learn.infiniteskills.com
	// Upgrade-Insecure-Requests: 1
	// User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36
	// Content-Type: application/x-www-form-urlencoded
	// Referer: https://learn.infiniteskills.com/login.html
	// Accept-Encoding: gzip, deflate
	// Accept-Language: en-US,en;q=0.8
	// Cookie: is_learn=YmVhdWdhbGJyYWl0aA%3D%3D; iskillslearn=14532496193884
	//
	// username=beaugalbraith&password=+divxfactory-btcob5&remember=yes

	req, err := http.NewRequest("POST", "https://learn.infiniteskills.com/ajax/login.html", nil)
	cookie := http.Cookie{
		Name:  "username",
		Value: "beaugalbraith",
		Name:  "password",
		Value: " divxfactory-btcob5",
	}

	req.AddCookie(&cookie)
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", resp.Header)
	defer resp.Body.Close()
	fmt.Printf("%s", scrape.Text(root))
}

func main() {
	login()
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://learn.infiniteskills.com/player.html?sku=02277", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// req.Header.Set("User-Agent", ua)
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// root, err := html.Parse(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// jwvideo := scrape.ByClass("wrapper")
	// jv := scrape.FindAll(root, jwvideo)
	// for _, v := range jv {
	// 	fmt.Println(scrape.Text(v))
	// }
}
