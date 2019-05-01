# docsite

A documentation site generator that fits [Sourcegraph](https://sourcegraph.com)'s needs:

- Markdown source files that are browseable on the file system and readable as plain text (without directives or complex front matter or configuration)
- Served by an HTTP server, not generated as static HTML files, to eliminate the need for external site host configuration (which we found to be error-proobe)
- Usable within sourcegraph to self-host docs for the current product version (with the same rendering and structure)

## Usage

```shell
go get github.com/prince1809/docsite/cmd/docsite
docsite -h
```

- `docsite check`: check for common problems (such as brokej¥n links)
- `docsite serve`: serve the site over HTTP

To use docsite for docs.sourcegraph.com see "[Documentation site](https://docs.sourcegraph.com/dev/documentation/site)"in the Sourcegraph documentation.

### Checks

The `docsite check` command runs various checks on your documentation site to find problems:

- invlaid links
- broken links
- disconnected pages (with no inlinks from other pages)

If any problems are found, it exists with a non-zero status code.

To ignore the disconnected page check for a page, and YAML `ignoreDisconnectedPageChecks: true` to the top matter in the beginning of the `.md` file. For example:

```
---
ignoreDisconnectedPageChecks: true
---

# My page title
```

## site data

The site data describes the location of its templates, assets, and content. It is a JSON object with the following properties.

- `content`: a VFS URL for the Markdown content files.
- `baseURLPath`: the URL path where the site is available (such as `/` or `/help/`).
- `templates`: a VFS URL for the [Go-style HTML templates](https://golang.org/pkg/html/template/) used to render site pages.
- `assets`: a VFS URL for the static assets referred to in the HTML templates (such as CSS stylesheets).
- `assetsBaseURLPath`: the URL path where the assets are available (such as `/assets/`).
- `allowRevisions`: a regular expression specifying the allowed revisions (optional)
- `check` (optional): an object containing a single property `ignoreURLPattern`, which is a [RE2 regexp](https://golang.org/pkg/regexp/syntax) of URLs to ignore when checking for broken URLs with `docsite check`.

The possible values for VFS URLs are:

- A **relative path to a local directory** (such as `../myrepo/doc`). The path is interpreted relative to the `docsite,json` file (if it exists) or the current working directory (if site data is specified in `DOCSITE_CONFIG`).
- An **absolute URL to a Zip archive** (with `http` or `https` scheme). The URL can contain a fragment (such as `#mydir/`) to refer to a specific directory in the archive.