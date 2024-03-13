package value

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ItemContent = string

// NewItemContent takes a given HTML string and
// returns a sanitized version with 'target="_blank"'
// added to all links.
func NewItemContent(content string) ItemContent {
	if len(content) < 1 {
		return ""
	}

	ctx := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
	}

	p := bluemonday.UGCPolicy()
	p.AllowElements("code")
	p.AllowAttrs("class").OnElements("code")
	p.AllowAttrs("target").OnElements("a")

	ns, err := html.ParseFragment(strings.NewReader(content), ctx)
	if err != nil {
		return p.Sanitize(content)
	}

	var node *html.Node
	for _, n := range ns {
		if n.Type == html.ElementNode {
			node = n
			break
		}
	}

	// add a target attribute to all links
	doc := goquery.NewDocumentFromNode(node)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("target", "_blank")
	})

	// sanitize, allowing
	// code blocks with language classes
	safe := p.Sanitize(content)

	// stringify the document
	h, err := goquery.OuterHtml(doc.Selection)
	if err != nil {
		return safe
	}

	return h
}
