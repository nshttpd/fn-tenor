package main

import (
	"context"
	"io"

	"net/url"

	"github.com/fnproject/fdk-go"
	"github.com/nshttpd/fn-tenor/tenor"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(tenorSearchHandler))
}

func tenorSearchHandler(ctx context.Context, in io.Reader, out io.Writer) {
	fdk.SetHeader(out, "Content-Type", "application/json")

	fnctx := fdk.Context(ctx)

	if fnctx.Config["FN_METHOD"] != "GET" {
		fdk.WriteStatus(out, 404)
		io.WriteString(out, `{"error":"route not found}"`)
		return
	}

	u, err := url.Parse(fnctx.Config["FN_REQUEST_URL"])
	if err != nil {
		fdk.WriteStatus(out, 500)
		io.WriteString(out, `{"error":"error parsing source url"}`)
		return
	}

	m, err := url.ParseQuery(u.RawQuery)

	q := m.Get("q")

	if q == "" {
		fdk.WriteStatus(out, 400)
		io.WriteString(out, `{"error":"missing query parameter"}`)
		return
	}

	d := tenor.SearchTenor(q)

	if d != nil {
		fdk.WriteStatus(out, 200)
		out.Write(d)
	} else {
		fdk.WriteStatus(out, 500)
		io.WriteString(out, `{"error":"error reading response from remote"}`)
	}

	return
}
