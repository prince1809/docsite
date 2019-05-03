package markdown

import (
	"context"

	"golang.org/x/net/html"
)

func evalMarkdownFuncs(ctx context.Context, htmlFragment []byte, opt Options) ([]byte, error) {
	panic("Implement me")
}

func getMarkdownFuncInvocation(tok html.Token) (funcName string, args map[string]string) {
	panic("Implement me")
}

func consumeUntilCloseTag(z *html.Tokenizer, funcName string) error {
	panic("implement me")
}
