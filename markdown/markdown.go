package markdown

import (
	"context"
	"net/url"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/styles"
	"gopkg.in/russross/blackfriday.v2"
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

// Options customize how Run parses and HTML-renders the Markdown document.
type Options struct {
	// Base is the base URL (typically including only the path, such as "/" or "/hello/") to use when
	// resolving relative links.
	Base *url.URL

	// ContentFilePathToLinkPath converts references to file paths of other content files to the URL
	// path to use in links. For example, contentFilePathToLinksPath("a/index.md") == "a".
	ContentFilePathToLinkPath func(string) string

	// Funcs are custom functions that can be invoked within Markdown documents with
	// <markdownfuncjfunction-name arg="val" />
	Funcs FuncMap

	// FuncInfo contains information passed to Markdown functions about the cuttent execution
	// context.
	FuncInfo FuncInfo
}

// FuncMap contains named functions that can be invoked within Markdown documents (see
// (Options).Func).
type FuncMap map[string]func(context.Context, FuncInfo, map[string]string) (string, error)

// FuncInfo contains information passed to Markdown functions about the current execution context.
type FuncInfo struct {
	Version string // the version of the content containing the page to render
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

func hasSingleChildOfType(node *blackfriday.Node, typ blackfriday.NodeType) bool {
	seenLink := false
	for child := node.FirstChild; child != nil; child = child.Next {
		switch {
		case child.Type == blackfriday.Text && len(child.Literal) == 0:
			continue
		case child.Type == blackfriday.Link && !seenLink:
			seenLink = true
		default:
			return false
		}
	}
	return seenLink
}
