package port

import "web_scraper/internal/scraping/domain/model"

type Scraper interface {
	ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content
}
