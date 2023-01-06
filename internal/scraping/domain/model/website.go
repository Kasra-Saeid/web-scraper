package model

type Website struct {
	Url        string
	Attributes []Attribute
	HtmlTexts  []HtmlText
	Pages      []int
}

type Attribute struct {
	InAppName string
	HtmlQuery string
	Name      string
	Text      string
}

type HtmlText struct {
	InAppName string
	HtmlQuery string
	Text      string
}

func NewAttribute(inAppName, htmlQuery, name string) *Attribute {
	return &Attribute{
		InAppName: inAppName,
		HtmlQuery: htmlQuery,
		Name:      name,
		Text:      "",
	}
}

func NewHtmlText(inAppName, htmlQuery string) *HtmlText {
	return &HtmlText{
		InAppName: inAppName,
		HtmlQuery: htmlQuery,
		Text:      "",
	}
}

func NewWebsite(url string, attrs []Attribute, htmlTexts []HtmlText, pages []int) *Website {
	return &Website{Url: url, Attributes: attrs, HtmlTexts: htmlTexts, Pages: pages}
}
