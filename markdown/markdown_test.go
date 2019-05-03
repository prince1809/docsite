package markdown

import (
	"context"
	"testing"
)

func check(t *testing.T, got, want Document) {

}

func TestRun(t *testing.T) {
	ctx := context.Background()
	t.Run("no metadata", func(t *testing.T) {
		doc, err := Run(ctx, []byte(`# My title
Hello world github/linguist#1 **cool**, and #1!`), Options{})
		if err != nil {
			t.Fatal(err)
		}
		check(t, *doc, Document{
			Title: "My title",
			HTML: []byte(`<h1 id="my-title"><a name="my-title" class="anchor" href="#my-title" rel="nofollow" aria-hidden="true"></a>My title</h1>
<p>Hello world github/linguist#1 <strong>cool</strong>, and #1!</p>
`),
		})
	})
}

func TestRenderer(T *testing.T) {

}
