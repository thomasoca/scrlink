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

func writeFile(filename string, line string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	newLine := line + "\n"
	n, err := f.WriteString(newLine)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
	defer f.Close()
}

func scrapper(url string, filename string) {
	var baseUrl = "commons.wikimedia.org"

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(baseUrl),
	)

	// Instantiate collector for image link
	imgCollector := c.Clone()

	// On every a element which has href attribute call callback
	// Get gallerytext class
	c.OnHTML(".gallerytext", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		imgLink := e.Request.AbsoluteURL(link)
		imgCollector.Visit(imgLink)
	})

	imgCollector.OnHTML("#file", func(e *colly.HTMLElement) {
		img := e.ChildAttr("a", "href")
		fmt.Printf("Image link: %s\n", img)
		writeFile(filename, img)
	})

	// Start scraping on wikimedia images
	c.Visit(url)
}

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	switch {
	case argsLen < 2:
		fmt.Println("Missing some arguments, please check your input")
	case argsLen > 2:
		fmt.Println("Too many arguments")
	}
	wikimediaUrl := args[1]
	fileName := args[2]

	scrapper(wikimediaUrl, fileName)
}
