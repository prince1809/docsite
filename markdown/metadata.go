package markdown

// Metadata is document metadata in the "front matter" of a Markdown document.
type Metadata struct {
	Title                       string `yaml:"title"`
	IgnoreDisconnectedPageCheck bool   `yaml:"ignoreDisconnectedPageCheck"`
}

func parseMetadata(input []byte) (meta Metadata, markdown []byte, err error) {
	panic("implement me")
}
