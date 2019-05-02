package docsite

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

// VersionedFileSystem represents multiple version of an http.FileSystem.
type VersionedFileSystem interface {
	OpenVersion(ctx context.Context, version string) (http.FileSystem, error)
}

// Site represents a documentation site, including all of its templates, assets, and content.
type Site struct {
	// Content is the versioned file system containing the Markdown files and assets (e.g. images)
	// embedded in them.
	Content VersionedFileSystem

	// Base is the base URL (typically including only the path, such as "/" or "/help/") where the
	// site is avilable
	Base *url.URL

	// Templates is the file system containing the go html/template templates used to render site
	// page
	Templates http.FileSystem

	// Assets it the file system containing the site-wide static asset files (e.g., global styles
	// and logo).
	Assets http.FileSystem

	// AssetBase is the base URL (sometimg only including the path, such as "/assets/") where the
	// assets are available.
	AssetBase *url.URL

	// CheckIgnoreURLPattern is a regex matching URLs to ignore in the check method.
	CheckIgnoreURLPattern *regexp.Regexp
}

// docsiteConfig is the shape of docsite.json
//
// See ["Site data" in README.md](../../README.md#site-data) for documentation on this type's
// fields.
type dockSiteConfig struct {
	Content           string
	BaseURLPath       string
	Templates         string
	Assets            string
	AssetsBaseURLPath string
	Check             struct {
		IgnoreURLPattern string
	}
}

type nonVersionedFileSystem struct{ http.FileSystem }

type versionedFileSystemURL struct {
	url   string
	mu    sync.Mutex
	cache map[string]fileSystemCacheEntry
}

type fileSystemCacheEntry struct {
	fs http.FileSystem
	at time.Time
}

const fileSystemCacheTTL = 5 * time.Minute

// PageData is the data avialable to the HTML template used to render a page.
type PageData struct {
	CurrentVersion  string // content version string requested
	ContentPagePath string // content path page requested

	ContentVersionNotFoundError bool // whether the requested version was not found
	ContentPageNotFoundError    bool // Whether the requested content page was not found

	// Content is the content page, when it is found.
	Content *ContentPage
}
