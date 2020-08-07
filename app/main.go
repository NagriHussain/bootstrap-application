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
)

const defaultPort = "8080"

var userCounter uint64 = 0

type response struct {
	UserAgent  string    `json:"user_agent"`
	RemoteAddr string    `json:"remote_address"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment the user count by 1
	atomic.AddUint64(&userCounter, 1)
	var resp response
	resp.UserAgent = r.UserAgent()
	resp.Time = time.Now()
	resp.RemoteAddr = r.RemoteAddr
	resp.Counter = atomic.LoadUint64(&userCounter)
	data, err := json.Marshal(&resp)
	fmt.Fprintln(os.Stderr, string(data))
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)
}

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = defaultPort
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
