package port

type Scraper interface {
	ScrapeCards(url, parrentClass, itemClass string, attrs map[string]string, texts []string)
}
