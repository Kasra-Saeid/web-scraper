package colly

import "github.com/gocolly/colly"

type Colly struct {
	Collector colly.Collector
}

func New(collector colly.Collector) Colly {
	return Colly{Collector: collector}
}

func (c Colly) ScrapeCards(url, parrentClass, itemClass string, attrs map[string]string, texts []string) {
	c.Collector.OnHTML(parrentClass, func(e *colly.HTMLElement) {
		e.ForEach(itemClass, func(i int, he *colly.HTMLElement) {
			for class, attr := range attrs {
				he.ChildAttr(class, attr)
			}
			for _, text := range texts {
				he.ChildText(text)
			}
		})
	})
	c.Collector.Visit(url)
}
