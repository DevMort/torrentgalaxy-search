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
	magnet   string
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

	j := 1
	for i := len(entries) - 1; i >= 0; i-- {
		fmt.Printf("[%v] %s (%s, %s, %s)\n\n", color.BlueString(fmt.Sprint(j)), entries[i].title, entries[i].username, color.CyanString(entries[i].size), color.RedString(entries[i].seeds))
		j++
	}

	fmt.Printf("\nWhich would you like to see the info of? ")
	var choice int
	fmt.Scanf("%2d", &choice)

	fmt.Printf("\n%s\n%s\n\n", entries[len(entries)-choice].magnet, color.YellowString("https://torrentgalaxy.to"+entries[len(entries)-choice].link))
}

func get_entry(c *colly.Collector, entries *[]Entry) {
	c.OnHTML("div.tgxtablerow.txlight", func(e *colly.HTMLElement) {
		title, _ := e.DOM.Find("a.txlight").Attr("title")
		if title != "comments" {
			link, _ := e.DOM.Find("a.txlight").Attr("href")
			magnet, _ := e.DOM.Find("i.glyphicon.glyphicon-magnet").Closest("a").Attr("href")
			entry := Entry{
				title:    title,
				link:     link,
				username: e.DOM.Find("a.username").Text(),
				size:     e.DOM.Find("span.badge.badge-secondary.txlight").Text(),
				seeds:    e.DOM.Find("font[color=green]").Last().Text(),
				magnet:   magnet,
			}
			*entries = append(*entries, entry)
		}
	})
}
