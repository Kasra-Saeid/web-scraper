package scraping

import (
	"web_scraper/internal/scraping/domain/service"
	"web_scraper/pkg/colly"
	"web_scraper/pkg/csv_file"

	collyPkg "github.com/gocolly/colly"
)

type Scraping struct {
	scrapingService service.Scraping
}

func New() Scraping {
	collector := collyPkg.NewCollector(collyPkg.Async(true))
	scraperPkg := colly.New(collector)
	csvPkg := csv_file.New()
	scrapingService := service.NewScraper(scraperPkg, csvPkg)
	scraping := Scraping{scrapingService: scrapingService}
	return scraping
}

func (s Scraping) GetScrapingServcie() service.Scraping {
	return s.scrapingService
}
