package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/NagriHussain/bootstrap-application/utils"
)

type publicAPI struct {
	Count   int `json:"count"`
	Entries []struct {
		API         string `json:"API"`
		Description string `json:"Description"`
		Auth        string `json:"Auth"`
		HTTPS       bool   `json:"HTTPS"`
		Cors        string `json:"Cors"`
		Link        string `json:"Link"`
		Category    string `json:"Category"`
	} `json:"entries"`
}

const (
	pubAPIEndpoint  = "api.publicapis.org"
	schema          = "https"
	constCategories = "categories"
	constEntries    = "entries"
	constCategory   = "category"
	defaultPort     = "9090"
)

func getCategories() []string {
	var categoryList []string
	body := utils.HTTPGETCall(getURL(constCategories))
	if err := json.Unmarshal(body, &categoryList); err != nil {
		log.Fatal(err)
	}
	return categoryList
}

func chooseCategories(categoryListLen int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(categoryListLen)
}

func getURL(path string) string {
	url := url.URL{
		Scheme: schema,
		Host:   pubAPIEndpoint,
		Path:   path,
	}
	return url.String()
}

func prepAPICall() string {
	categoryList := getCategories()
	randINT := chooseCategories(len(categoryList))
	category := categoryList[randINT]
	fmt.Fprintln(os.Stderr, utils.Now()+" Making an API call for category - "+category)
	u, err := url.Parse(getURL(constEntries))
	if err != nil {
		print(err)
	}
	q := u.Query()
	q.Set(constCategory, category)
	u.RawQuery = q.Encode()
	return u.String()
}

//callAPI calls the free api api.publicapis.org and generates random JSON for consumption
func callAPI() []byte {
	u := prepAPICall()
	body := utils.HTTPGETCall(u)
	var pubDS publicAPI
	if err := json.Unmarshal(body, &pubDS); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, utils.Now()+" API Data "+string(body))
	return body
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := callAPI()
	w.Write(body)
}

func main() {
	APIAppPort := os.Getenv("APIAPP_PORT")
	if APIAppPort == "" {
		APIAppPort = defaultPort
	}
	fmt.Fprintln(os.Stderr, utils.Now()+" Starting API application on http://localhost:"+APIAppPort)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+APIAppPort, nil))
}
