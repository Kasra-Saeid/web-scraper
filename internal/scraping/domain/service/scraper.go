package service

import (
	"web_scraper/internal/scraping/domain/model"
	"web_scraper/internal/scraping/domain/port"
)

type Scraping interface {
	ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content
}

type Scraper struct {
	scraperPkg port.Scraper
}

func NewScraper(scraperPkg port.Scraper) Scraper {
	return Scraper{scraperPkg: scraperPkg}
}

func (s Scraper) ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content {
	return s.scraperPkg.ScrapeCards(parrentClass, itemClass, website)
}
