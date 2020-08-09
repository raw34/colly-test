package main

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains
		colly.AllowedDomains("yonikimo.com"),
	)

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))

		if err != nil {
			log.Fatal(err)
		}

		nodes := htmlquery.Find(doc, `//*[@id="top"]/table/tbody/tr`)

		for _, node := range nodes {
			no := htmlquery.FindOne(node, "td[1]/text()")
			title := htmlquery.FindOne(node, "td[2]/a/text()")
			url := htmlquery.FindOne(node, "td[2]/a/@href")

			log.Println(no, title, url)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://yonikimo.com/story.html")
}
