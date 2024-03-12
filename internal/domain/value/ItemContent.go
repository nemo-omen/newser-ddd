package value

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

type ItemContent = string

// NewItemContent takes a given HTML string and
// returns a sanitized version with 'target="_blank"'
// added to all links.
func NewItemContent(content string) ItemContent {
	if len(content) < 1 {
		return ""
	}

	// sanitize, allowing
	// code blocks with language classes
	p := bluemonday.UGCPolicy()
	p.AllowElements("code")
	p.AllowAttrs("class").OnElements("code")
	safe := p.Sanitize(content)

	// add a target attribute to all links
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(safe))
	if err != nil {
		return safe
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("target", "_blank")
	})

	// stringify the doc
	updated, err := doc.Html()
	if err != nil {
		return safe
	}
	return updated
}
