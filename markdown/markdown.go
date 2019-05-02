package markdown

import (
	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/styles"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Document is a parsed and HTML-rendered Markdown document.
type Document struct {
	// Meta is the document's  metadata in the Markdown "front matter", if any.
	Meta Metadata

	// Title is taken from the metadata (if it exists) or else from the context of the first
	// heading.
	Title string

	// HTML is the rendered Markdown content.
	HTML []byte

	// Tree is the tree sections (used to show a table contents).
	Tree []*SectionNode
}

// NewParser creates a new Markdown parser (the same one used by Run).
func NewParser(renderer blackfriday.Renderer) *blackfriday.Markdown {
	return blackfriday.New(
		blackfriday.WithRenderer(renderer),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
	)
}

// NewBfRenderer creates the default blackfriday renderer to be passed to NewPrser()
func NewBfRenderer() blackfriday.Renderer {
	return bfchroma.NewRenderer(
		bfchroma.ChromaStyle(styles.VisualStudio),
		bfchroma.Extend(
			blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			}),
		),
	)
}
