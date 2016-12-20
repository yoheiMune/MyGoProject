package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main() {

	doc, err := goquery.NewDocument("http://www.yoheim.net/")
	if err != nil {
		fmt.Println("usr scarapiing failed.")
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})
 }