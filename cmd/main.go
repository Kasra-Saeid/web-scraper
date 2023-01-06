package main

import (
	"fmt"
	"web_scraper/internal/scraping"
	"web_scraper/internal/scraping/domain/model"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	scraping := scraping.New()
	scrapingService := scraping.GetScrapingServcie()
	website := model.NewWebsite(
		"https://arzdigital.com/breaking/",
		[]model.Attribute{*model.NewAttribute(".arz-breaking-news__item-link", "title"), *model.NewAttribute(".arz-breaking-news__item-link > .arz-breaking-news__container > .arz-breaking-news__info > .arz-breaking-news__publish-time > time", "datetime")},
		[]model.HtmlText{*model.NewHtmlText(".arz-breaking-news-post__info-rating-pump.arz-breaking-news-post__info-rating-value")},
		[]int{1},
	)
	c := scrapingService.ScrapeCards(".arz-breaking-news__list", ".arz-breaking-news__item", website)

	for _, content := range c {
		fmt.Println(content.Title)
	}
}
