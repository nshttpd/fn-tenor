package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	fdk "github.com/fnproject/fdk-go"
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

	var apikey string
	var ok bool

	fdk.SetHeader(out, "Content-Type", "application/json")

	if apikey, ok = os.LookupEnv(tenorAPIKey); !ok {
		fdk.WriteStatus(out, 400)
		io.WriteString(out, `{"error":"improper configuration"}`)
		return
	}

	fnctx := fdk.Context(ctx)

	if fnctx.Config["FN_METHOD"] != "GET" {
		fdk.WriteStatus(out, 404)
		io.WriteString(out, `{"error":"route not found"}`)
		return
	}

	u := fmt.Sprintf(tenorTrendingAPI, apikey, tenorRequestLimit)
	log.Printf("uri : %s", u)

	var resp *http.Response
	var err error

	if resp, err = http.Get(u); err != nil {
		fdk.WriteStatus(out, 502)
		io.WriteString(out, `{"error":"bad response from remote"}`)
		log.Printf("error on request to tenor for trending results: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fdk.WriteStatus(out, 200)
		out.Write(body)
	} else {
		fdk.WriteStatus(out, 500)
		io.WriteString(out, `{"error":"error reading response from remote"}`)
		log.Println("error reading response from tenor")
		log.Println(err)
	}

	return

}
