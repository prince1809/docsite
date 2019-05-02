package markdown

// Document is a parsed and HTML-rendered Markdown document.
type Document struct {
	// Meta is the document's  metadata in the Markdown "front matter", if any.
	Meta Metadata
}
