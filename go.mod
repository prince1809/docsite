module github.com/prince1809/docsite

go 1.12

require (
	github.com/Depado/bfchroma v1.1.1
	github.com/alecthomas/chroma v0.6.3
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shurcooL/sanitized_anchor_name v1.0.0
	github.com/sourcegraph/go-jsonschema v0.0.0-20221230021921-34aaf28fc4ac
	github.com/sourcegraph/jsonschemadoc v0.0.0-20200429204751-398086c46c99
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/net v0.7.0
	golang.org/x/tools v0.1.12
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1
