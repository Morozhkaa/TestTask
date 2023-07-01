package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"parser/models"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("failed to create file, err: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// create a new сollector instance
	scrapeUrl := "https://hypeauditor.com/top-instagram-all-russia/"
	c := colly.NewCollector(
		colly.AllowedDomains("hypeauditor.com"),
	)
	ind := 1

	// process rows with information about influencers
	c.OnHTML("div.row__top", func(h *colly.HTMLElement) {
		var err error
		var num int

		// rating processing: get delta value with + or -
		if h.ChildText("span.ml-2") != "" {
			text := h.ChildText("div.delta.delta-value")
			num, err = strconv.Atoi(h.ChildText("span.ml-2"))
			if err != nil {
				return
			}
			if text[len(text)-1]-'0' == 103 { // recognize the red triangle
				num = -num
			}
		}

		// form a string with categories separated by ;
		var categories string
		h.ForEach("div.tag__content.ellipsis", func(i int, h *colly.HTMLElement) {
			categories += h.Text + ";"
		})
		if len(categories) > 0 {
			categories = categories[:len(categories)-1]
		}

		// fill influencer instance with values
		data := &models.Influencer{
			Id:         fmt.Sprint(ind),
			Rank:       fmt.Sprint(num),
			ImageUrl:   h.ChildAttr("img.avatar__img", "src"),
			Nickname:   h.ChildText("div.contributor__name-content"),
			Name:       h.ChildText("div.contributor__title"),
			Link:       "https://hypeauditor.com" + h.ChildAttr("a", "href"),
			Categories: categories,
			Followers:  h.ChildText("div.row-cell.subscribers"),
			Country:    h.ChildText("div.row-cell.audience"),
			EngAuth:    h.ChildText("div.row-cell.authentic"),
			EngAvg:     h.ChildText("div.row-cell.engagement"),
			PageUrl:    "https://hypeauditor.com" + h.ChildAttr("a.button.button--theme-secondary.button--size-md", "href"),
		}

		// save string to csv file
		writer.Write([]string{data.Id, data.Rank, data.ImageUrl, data.Nickname, data.Name, data.Link,
			data.Categories, data.Followers, data.Country, data.EngAuth, data.EngAvg, data.PageUrl})
		ind++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.Visit(scrapeUrl)
}
