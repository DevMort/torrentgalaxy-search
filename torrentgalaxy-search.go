package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
)

type Entry struct {
	title    string
	link     string
	username string
	size     string
	seeds    string
}

func main() {
	if len(os.Args) != 2 {
		panic("Only one argument is needed (the search term/s)!")
	}

	search_term := os.Args[1]
	c := colly.NewCollector(colly.AllowedDomains(
		"torrentgalaxy.to",
	))

	var entries []Entry
	get_entry(c, &entries)

	c.Visit("https://torrentgalaxy.to/torrents.php?search=" + strings.ReplaceAll(search_term, " ", "+") + "&sort=seeders&order=desc")

	for i := len(entries) - 1; i >= 0; i-- {
		fmt.Printf("%s (%s, %s, %s)\n%s\n\n", entries[i].title, entries[i].username, color.CyanString(entries[i].size), color.RedString(entries[i].seeds), color.YellowString("https://torrentgalaxy.to"+entries[i].link))
	}
}

func get_entry(c *colly.Collector, entries *[]Entry) {
	c.OnHTML("div.tgxtablerow.txlight", func(e *colly.HTMLElement) {
		title, _ := e.DOM.Find("a.txlight").Attr("title")
		if title != "comments" {
			link, _ := e.DOM.Find("a.txlight").Attr("href")
			entry := Entry{
				title:    title,
				link:     link,
				username: e.DOM.Find("a.username").Text(),
				size:     e.DOM.Find("span.badge.badge-secondary.txlight").Text(),
				seeds:    e.DOM.Find("font[color=green]").Last().Text(),
			}

			*entries = append(*entries, entry)
		}
	})
}
