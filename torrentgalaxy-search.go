package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
)

func main() {
	if len(os.Args) != 2 {
		panic("Only one argument is needed (the search term/s)!")
	}

	search_term := os.Args[1]
	c := colly.NewCollector(colly.AllowedDomains(
		"torrentgalaxy.to",
	))

	c.OnHTML("div.tgxtablerow.txlight", func(e *colly.HTMLElement) {
		title, _ := e.DOM.Find("a.txlight").Attr("title")
		if title != "comments" {
			link, _ := e.DOM.Find("a.txlight").Attr("href")
			username := e.DOM.Find("a.username").Text()
			size := e.DOM.Find("span.badge.badge-secondary.txlight").Text()
			seeds := e.DOM.Find("font[color=green]").Text()

			fmt.Printf("%s: %s (%s, %s, %s)\n%s\n\n", color.GreenString("TITLE"), title, username, color.CyanString(size), color.RedString(seeds), color.YellowString("https://torrentgalaxy.to"+link))
		}
	})

	c.Visit("https://torrentgalaxy.to/torrents.php?search=" + strings.ReplaceAll(search_term, " ", "+") + "&sort=seeders&order=desc")
}
