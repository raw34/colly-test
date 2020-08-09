package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strings"
)

func main() {
	Init()
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

			log.Print(no, title, url)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Print("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://yonikimo.com/story.html")
}

func Init() {
	logFile, err := os.OpenFile("logs/cmd.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
