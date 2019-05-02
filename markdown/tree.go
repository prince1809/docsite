package markdown

import blackfriday "gopkg.in/russross/blackfriday.v2"

// SectionNode is a section and its children.
type SectionNode struct {
	Title    string         // section title
	URL      string         // section URL (usually an anchor links)
	Level    int            // heading level (1-6)
	Children []*SectionNode // subsections
}

func newTree(node *blackfriday.Node) []*SectionNode {
	stack := []*SectionNode{{}}
	return stack[0].Children
}
