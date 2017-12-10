package tenor

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	tenorAPIKey       = "TENOR_API_KEY"
	tenorRequestLimit = 5
	tenorTrendingAPI  = "https://api.tenor.com/v1/autocomplete?type=trending&key=%s&limit=%d"
	tenorSearchAPI    = "https://api.tenor.com/v1/search?q=%s&key=%s&limit=%d&safesearch=moderate"
)

func getAPIKey() (string, bool) {
	key, ok := os.LookupEnv(tenorAPIKey)
	if !ok {
		log.Println("could not get API key from environment")
		return "", false
	}
	return key, true
}

func GetTenorTrending() []byte {
	if key, ok := getAPIKey(); ok {
		u := fmt.Sprintf(tenorTrendingAPI, key, tenorRequestLimit)
		if resp, err := http.Get(u); err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("error reading trend data from tenor")
				log.Println(err)
				return nil
			}
			return body
		}
		log.Println("error fetching data form tenor")
	}
	return nil
}

func SearchTenor(q string) []byte {
	if key, ok := getAPIKey(); ok {
		u := fmt.Sprintf(tenorSearchAPI, q, key, tenorRequestLimit)
		if resp, err := http.Get(u); err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("error reading search data from tenor")
				log.Println(err)
				return nil
			}
			return body
		}
		log.Println("error featching search data from tenor")
	}
	return nil
}
