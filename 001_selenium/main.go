package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
)

func main() {

	fmt.Println("Hello from Selenium sample.")

	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start phantomjs driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
	}

	if err := page.Navigate("http://www.yoheim.net"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}
	page.Screenshot("/tmp/a.png")
}
