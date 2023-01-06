package colly

import (
	"fmt"
	"log"
	"strconv"
	"web_scraper/internal/scraping/domain/model"

	"github.com/gocolly/colly"
)

type Colly struct {
	Collector *colly.Collector
}

func New(collector *colly.Collector) Colly {
	return Colly{Collector: collector}
}

func (c Colly) ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content {
	contentOfPage := make([]model.Content, 0)
	c.Collector.OnHTML(parrentClass, func(e *colly.HTMLElement) {
		e.ForEach(itemClass, func(i int, he *colly.HTMLElement) {
			newContent := model.Content{
				Title:    "",
				Date:     "",
				PosScore: 0,
				NegScore: 0,
			}
			for _, attr := range website.Attributes {
				fmt.Println(attr.Name)
				if attr.InAppName == "title" {
					newContent.Title = he.ChildAttr(attr.HtmlQuery, attr.Name)
				} else if attr.InAppName == "date" {
					newContent.Date = he.ChildAttr(attr.HtmlQuery, attr.Name)
				}
			}
			for _, htmlText := range website.HtmlTexts {
				if htmlText.InAppName == "bull_rate" {
					var err error
					newContent.PosScore, err = strconv.Atoi(he.ChildText(htmlText.HtmlQuery))
					if err != nil {
						log.Fatalln(err)
					}
				} else if htmlText.InAppName == "bear_rate" {
					var err error
					newContent.NegScore, err = strconv.Atoi(he.ChildText(htmlText.HtmlQuery))
					if err != nil {
						log.Fatalln(err)
					}
				}

			}
			contentOfPage = append(contentOfPage, newContent)

		})
	})
	for _, p := range website.Pages {
		c.Collector.Visit(fmt.Sprintf("%vpage/%d", website.Url, p))
	}
	c.Collector.Wait()
	return contentOfPage
}
