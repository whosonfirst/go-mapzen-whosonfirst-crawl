package main

import (
       "flag"
       "fmt"
	"com.mapzen/whosonfirst"
)

func main() {

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("crawl ", root)

	callback := func(path string, b whosonfirst.JSONBlob) error {
		fmt.Println("inflated ", path)
		return nil
	}

	c := whosonfirst.NewCrawler(root)
	_ = c.CrawlAndInflate(callback)
}