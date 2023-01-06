package service

import (
	"web_scraper/internal/scraping/domain/model"
	"web_scraper/internal/scraping/domain/port"
)

type Scraping interface {
	ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content
	WriteContents(allRows []model.Content, path *string)
}

type Scraper struct {
	scraperPkg port.Scraper
	repo       port.Repo
}

func NewScraper(scraperPkg port.Scraper, repo port.Repo) Scraper {
	return Scraper{scraperPkg: scraperPkg, repo: repo}
}

func (s Scraper) ScrapeCards(parrentClass string, itemClass string, website *model.Website) []model.Content {
	return s.scraperPkg.ScrapeCards(parrentClass, itemClass, website)
}

func (s Scraper) WriteContents(allRows []model.Content, path *string) {
	s.repo.WriteContents(allRows, path)
	s.repo.Close()
}
