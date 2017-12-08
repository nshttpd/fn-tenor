package main

import (
	"context"
	"io"

	fdk "github.com/fnproject/fdk-go"
	"github.com/nshttpd/fn-tenor/tenor"
)

const (
	tenorAPIKey       = "TENOR_API_KEY"
	tenorRequestLimit = 5
	tenorTrendingAPI  = "https://api.tenor.com/v1/autocomplete?type=trending&key=%s&limit=%d"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(tenorHandler))
}

func tenorHandler(ctx context.Context, in io.Reader, out io.Writer) {

	fdk.SetHeader(out, "Content-Type", "application/json")

	fnctx := fdk.Context(ctx)

	if fnctx.Config["FN_METHOD"] != "GET" {
		fdk.WriteStatus(out, 404)
		io.WriteString(out, `{"error":"route not found"}`)
		return
	}

	d := tenor.GetTenorTrending()
	if d != nil {
		fdk.WriteStatus(out, 200)
		out.Write(d)
	} else {
		fdk.WriteStatus(out, 500)
		io.WriteString(out, `{"error":"error reading response from remote"}`)
	}

	return

}
