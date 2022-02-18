package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile(line string) {
	f, err := os.OpenFile("links.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	newLine := line + "\n"
	n, err := f.WriteString(newLine)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
	defer f.Close()
}

func main() {
	var baseUrl = "commons.wikimedia.org"
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(baseUrl),
	)

	imgCollector := c.Clone()

	// On every a element which has href attribute call callback
	c.OnHTML(".gallerytext", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		imgLink := e.Request.AbsoluteURL(link)
		imgCollector.Visit(imgLink)
	})

	imgCollector.OnHTML("#file", func(e *colly.HTMLElement) {
		img := e.ChildAttr("a", "href")
		fmt.Printf("Image link: %s\n", img)
		writeFile(img)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	imgCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting image link", r.URL)
	})

	// Start scraping on wikimedia images
	c.Visit("https://commons.wikimedia.org/wiki/Category:Drawings_by_User:LuxAmber")
}
