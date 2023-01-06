package colly

import (
	"fmt"
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
	fmt.Println("here")
	contentOfPage := make([]model.Content, 0)
	c.Collector.OnHTML(parrentClass, func(e *colly.HTMLElement) {

		fmt.Println(*e)
		e.ForEach(itemClass, func(i int, he *colly.HTMLElement) {
			newContent := model.Content{
				Title:    "",
				Date:     "",
				PosScore: 0,
				NegScore: 0,
			}
			for _, attr := range website.Attributes {
				fmt.Println(attr.Name)
				if attr.Name == "title" {
					newContent.Title = he.ChildAttr(attr.HtmlQuery, attr.Name)
				} else if attr.Name == "datetime" {
					fmt.Println(he.ChildAttr(attr.HtmlQuery, attr.Name))
				}
			}
			for _, htmlText := range website.HtmlTexts {
				_ = htmlText
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
