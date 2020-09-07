package api

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Answer models an answer to a Question
type Answer struct {
	AnswerID   int    `json:"answer_id"`
	Body       string `json:"body"`
	IsAccepted bool   `json:"is_accepted"`
	Link       string `json:"link"`
	Score      int    `json:"score"`
	Code       string `json:"code"`
}

// Extract extracts code blocks from the Answer body
func (a *Answer) Extract() (string, error) {
	// parse the body string into a DOM
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.Body))
	if err != nil {
		return "", fmt.Errorf("failed to parse document: %v", err)
	}

	// extract and return code blocks
	blocks := []string{}
	doc.Find("pre code").Each(func(i int, s *goquery.Selection) {
		blocks = append(blocks, s.Text())
	})

	return strings.TrimSpace(strings.Join(blocks, "\n")), nil
}
