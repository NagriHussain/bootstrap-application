/*Author Hussain Nagri

 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/NagriHussain/bootstrap-application/utils"
)

const defaultPort = "8080"

var (
	userCounter    uint64 = 0
	aPIAppEndpoint string
)

type response struct {
	UserAgent  string    `json:"user_agent"`
	RemoteAddr string    `json:"remote_address"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
	APIData    string    `json:"api_data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment the user count by 1
	atomic.AddUint64(&userCounter, 1)
	var resp response
	resp.UserAgent = r.UserAgent()
	resp.Time = time.Now()
	resp.RemoteAddr = r.RemoteAddr
	resp.Counter = atomic.LoadUint64(&userCounter)
	// This is where API call to the other docker service will happen.
	resp.APIData = string(utils.HTTPGETCall(aPIAppEndpoint))

	data, err := json.Marshal(&resp)
	fmt.Fprintln(os.Stderr, string(data))
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)
}

func main() {
	webAppPort := os.Getenv("WEBAPP_PORT")
	//API_APP_ENDPOINT should be a string in the syntax of "http://localhost:9090"
	if webAppPort == "" {
		webAppPort = defaultPort
	}
	aPIAppEndpoint = os.Getenv("APIAPP_ENDPOINT")
	if aPIAppEndpoint == "" {
		fmt.Fprintln(os.Stderr, utils.Now()+" No API_APP_ENDPOINT found, make sure API_APP_ENDPOINT Environment variable is set.\n"+
			"Something like 'http://localhost:9090'\nTerminating process.")
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, utils.Now()+" Starting web application on http://localhost:"+webAppPort)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+webAppPort, nil))
}
