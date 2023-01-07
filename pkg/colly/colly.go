package colly

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"web_scraper/internal/scraping/domain/model"

	"github.com/gocolly/colly"
)

type Colly struct {
	Collector *colly.Collector
}

func New(collector *colly.Collector, opts ...Option) Colly {
	c := Colly{Collector: collector}
	for _, opt := range opts {
		opt(c.Collector)
	}
	return c
}

func (c Colly) ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content {
	contentOfPage := make([]model.Content, 0)
	c.Collector.OnHTML(parrentClass, func(e *colly.HTMLElement) {
		e.ForEach(itemClass, func(_ int, he *colly.HTMLElement) {
			newContent := model.Content{
				Title:    "",
				Date:     "",
				PosScore: 0,
				NegScore: 0,
			}
			for _, attr := range website.Attributes {
				if attr.InAppName == "date" {
					newContent.Date = he.ChildAttr(attr.HtmlQuery, attr.Name)
				} else if attr.InAppName == "title" {
					newContent.Title = he.ChildAttr(attr.HtmlQuery, attr.Name)
				}
			}
			for _, htmlText := range website.HtmlTexts {
				if htmlText.InAppName == "bull_rate" {
					newContent.PosScore = StringNumberToInt(he.ChildText(htmlText.HtmlQuery))
				} else if htmlText.InAppName == "bear_rate" {
					newContent.NegScore = StringNumberToInt(he.ChildText(htmlText.HtmlQuery))
				}

			}
			contentOfPage = append(contentOfPage, newContent)

		})
	})
	for _, p := range website.Pages {
		c.Collector.Visit(fmt.Sprintf("%vpage/%d", website.Url, p))
		fmt.Println(ShowScrapingProcess(p, website.Pages))
		fmt.Println(fmt.Sprintf("%vpage/%d", website.Url, p))
	}
	c.Collector.Wait()
	return contentOfPage
}

func RandomAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1",
		"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.1)",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 Edg/87.0.664.75",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363",
	}
	max, min := len(userAgents), 0
	randomNumberInAgentsLengthRange := rand.Intn(max-min) + min
	return userAgents[randomNumberInAgentsLengthRange]
}

func ShowScrapingProcess(currentPage int, pages []int) float32 {
	return float32(float32(currentPage)/float32(len(pages))) * 100
}

func StringNumberToInt(num string) int {
	resNumInString := ""
	var resNum float64 = 0
	var err error
	resNumInString = num
	resNum, err = strconv.ParseFloat(resNumInString, 64)
	if err != nil {
		if strings.HasSuffix(num, "k") {
			resNumInString = strings.Trim(num, "k")
			resNum, err = strconv.ParseFloat(resNumInString, 64)
			if err != nil {
				log.Fatalln(err)
			}
			return int(resNum) * 1000
		} else if strings.HasSuffix(num, "K") {
			resNumInString = strings.Trim(num, "K")
			resNum, err = strconv.ParseFloat(resNumInString, 64)
			if err != nil {
				log.Fatalln(err)
			}
			return int(resNum) * 1000
		} else {
			return 0
		}
	}
	return int(resNum)
}
