package docsite

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prince1809/docsite/markdown"
	"github.com/sourcegraph/go-jsonschema/jsonschema"
	"github.com/sourcegraph/jsonschemadoc"
	"net/url"
	"strings"
)

// createMarkdownFuncs creates the standard set of Markdown functions expected by documentation
// content. Documentation pages can invoke these functions with special tags such as <div
// markdown-func=myfunc1 myfunc:arg1="foo" />. The only function currently defined in jsonschemadoc,
// which generates documentation for a JSON schema.

func createMarkdownFuncs(site *Site) markdown.FuncMap {
	return markdown.FuncMap{
		"jsonschemadoc": func(ctx context.Context, info markdown.FuncInfo, args map[string]string) (string, error) {
			inputPath := args["path"]
			if inputPath == "" {
				return "", errors.New("no path to JSON Schema file is specified (use <div markdown-func=jsonschemadoc:path=PATH>)")
			}

			content, err := site.Content.OpenVersion(ctx, info.Version)
			if err != nil {
				return "", err
			}
			data, err := ReadFile(content, inputPath)
			if err != nil {
				return "", err
			}

			var schema *jsonschema.Schema
			if err := json.Unmarshal(data, &schema); err != nil {
				return "", err
			}

			// Support JSON references to emit documentation for a sub-definition.
			if ref := args["ref"]; ref != "" {
				if !strings.HasPrefix(ref, "#") {
					return "", fmt.Errorf("invalid JSON Schema reference %q (only URI fragments are supported)", ref)
				}
				u, err := url.Parse(ref)
				if err != nil {
					return "", errors.WithMessage(err, "invalid JSON Schema reference")
				}
				if !strings.HasPrefix(u.Fragment, "/definitions/") || strings.Count(u.Fragment, "/") != 2 {
					return "", fmt.Errorf("unsupported JSON Schema reference %q (only simple #/definitions/Foo references are supported)", u.Fragment)
				}
				defName := strings.TrimPrefix(u.Fragment, "/definitions/")
				if schema.Definitions == nil || (*schema.Definitions)[defName] == nil {
					return "", fmt.Errorf("unable to resolve JSON Schema reference %q", u.Fragment)
				}
				schema = (*schema.Definitions)[defName]
			}

			out, err := jsonschemadoc.Generate(schema)
			if err != nil {
				return "", err
			}

			doc, err := markdown.Run(ctx, []byte("<div class='pre-wrap'>\n```javascript\n"+string(out)+"\n```\n</div>"), markdown.Options{})
			if err != nil {
				return "", err
			}
			return string(doc.HTML), nil
		},
	}
}
