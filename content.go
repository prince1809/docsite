package docsite

import "github.com/prince1809/docsite/markdown"

// ContentPage requests a Markdown-formatter documentation page. To create a ContentPage, use of the Site methods.
type ContentPage struct {
	Path        string            // the canonical URL path (without ".md" or "/index.md")
	FilePath    string            // the filename on disk
	Data        []byte            // the page's file contents
	Doc         markdown.Document // the Markdown doc
	BreadCrumbs []breadcrumbEntry // ancestor breadcrumb for this page
}

type breadcrumbEntry struct {
	Label    string
	URL      string
	IsActive bool
}
