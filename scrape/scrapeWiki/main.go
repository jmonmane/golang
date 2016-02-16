package main

import (
	"fmt"
	"os"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

const (
	tableID = "tableMorningJan2120162"
)

func scraper() {
	fd, err := os.Open("/mnt/hgfs/Downloads/wiki.html")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	root, err := html.Parse(fd)
	if err != nil {
		panic(err)
	}
	t := html.NewTokenizer(root)

	// matcher := func(n *html.Node) bool {
	// 	if n.DataAtom == atom.Table {
	// 		return true
	// 	}
	// 	return false
	// }

	// rowMatcher := func(n *html.Node) bool {
	// 	if n.DataAtom == atom.Tr {
	// 		return true
	// 	}
	// 	return false
	// }
	tableMatcher := scrape.ById(tableID)
	table := scrape.FindAll(root, tableMatcher)

	for _, v := range table {
		if t.Token().Data == "tr" {
			fmt.Printf("%s\n", scrape.Text(v))
		} else {
			t.Next()
		}
	}

	// for , v := range table  {
	// 	fmt.Printf("%s\n", scrape.Text(v))
	// }

}

func main() {
	scraper()

}
