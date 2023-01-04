package scraping

import (
	collyPkg "github.com/gocolly/colly"
	"web_scraper/internal/scraping/domain/service"
	"web_scraper/pkg/colly"
)

type Scraping struct {
	scrapingService service.Scraping
}

func New() Scraping {
	collector := collyPkg.NewCollector()
	scraperPkg := colly.New(collector)
	scrapingService := service.NewScraper(scraperPkg)
	scraping := Scraping{scrapingService: scrapingService}
	return scraping
}

func (s Scraping) GetScrapingServcie() service.Scraping {
	return s.scrapingService
}
