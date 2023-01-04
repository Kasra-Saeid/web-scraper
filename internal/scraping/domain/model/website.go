package model

type Website struct {
	Url        string
	Attributes []Attribute
	HtmlTexts  []HtmlText
	Pages      []int
}

type Attribute struct {
	HtmlQuery string
	Name      string
	Text      string
}

type HtmlText struct {
	HtmlQuery string
	Text      string
}

func NewAttribute(htmlQuery, name string) *Attribute {
	return &Attribute{
		HtmlQuery: htmlQuery,
		Name:      name,
		Text:      "",
	}
}

func NewHtmlText(htmlQuery string) *HtmlText {
	return &HtmlText{
		HtmlQuery: htmlQuery,
		Text:      "",
	}
}

func NewWebsite(url string, attrs []Attribute, htmlTexts []HtmlText, pages []int) *Website {
	return &Website{Url: url, Attributes: attrs, HtmlTexts: htmlTexts, Pages: pages}
}
