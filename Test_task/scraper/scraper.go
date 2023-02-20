package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("work")
	c := colly.NewCollector(
		colly.AllowedDomains("notion.site"),
	)
	c.Visit("https://iodized-saturn-c55.notion.site/8b85c6d0b84d4a369430482612332177?v=44e2251ec5ef4d4cbeaf889b9e5da920")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})

}
