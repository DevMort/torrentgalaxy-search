package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
)

type Entry struct {
	title string
	link  string
}

func main() {
	if len(os.Args) != 2 {
		panic("Only one argument is needed (the search term/s)!")
	}

	search_term := os.Args[1]
	c := colly.NewCollector(colly.AllowedDomains(
		"torrentgalaxy.to",
	))

	c.OnHTML("div.tgxtable", func(e *colly.HTMLElement) {
		e.DOM.Find("a.txlight").Each(func(_ int, entry *goquery.Selection) {
			title, _ := entry.Attr("title")
			if title != "comments" {
				link, _ := entry.Attr("href")

				new_entry := Entry{
					link:  link,
					title: title,
				}

				fmt.Printf("%s: %s\n%s\n\n", color.GreenString("TITLE"), new_entry.title, color.YellowString("https://torrentgalaxy.to"+new_entry.link))
			}
		})
	})

	c.Visit("https://torrentgalaxy.to/torrents.php?search=" + strings.ReplaceAll(search_term, " ", "+") + "&sort=seeders&order=desc")
}
