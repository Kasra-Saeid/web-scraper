package service

import (
	"web_scraper/internal/scraping/domain/port"
)

type Scraper struct {
	scraperPkg port.Scraper
}

func NewScraper(scraperPkg port.Scraper) Scraper {
	return Scraper{scraperPkg: scraperPkg}
}

func (s Scraper) Scrape(url, parrentClass, itemClass string, attrs map[string]string, texts []string) {
	s.scraperPkg.ScrapeCards(url, parrentClass, itemClass, attrs, texts)
}
