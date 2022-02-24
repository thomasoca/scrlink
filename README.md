# scrlink

A simple go program to scrape image links from [wikimedia](commons.wikimedia.org) category page using [colly](https://github.com/gocolly/colly) scraping library.

It would take the wikimedia category page (example: https://commons.wikimedia.org/wiki/Category:Drawings_by_User:LuxAmber) and write the image links from the page into a text output.

### Usage

Do a `go build` or a `go install` to run this program as a binary for convenience purpose.

```sh
./scrlink [WIKIMEDIA URL] [FILENAME]
```

example:

```sh
./scrlink https://commons.wikimedia.org/wiki/Category:Drawings_by_User:LuxAmber test.txt
```

