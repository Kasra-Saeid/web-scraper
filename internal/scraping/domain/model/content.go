package model

type Content struct {
	Title    string
	Date     string
	PosScore int
	NegScore int
}

func NewContent(title, date string, posScore, negScore int) *Content {
	return &Content{
		Title:    title,
		Date:     date,
		PosScore: posScore,
		NegScore: negScore,
	}
}
