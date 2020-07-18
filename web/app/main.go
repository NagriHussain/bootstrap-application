/*Author Hussain Nagri
 */
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

const logFile = "main.log"

var userCounter uint64 = 0

type response struct {
	UserAgent  string    `json:"user_agent"`
	RemoteAddr string    `json:"remote_address"`
	Date       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func setUpHandler() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment the user count by 1
	atomic.AddUint64(&userCounter, 1)
	var resp response
	resp.UserAgent = r.UserAgent()
	resp.Date = time.Now()
	resp.RemoteAddr = r.RemoteAddr
	resp.Counter = atomic.LoadUint64(&userCounter)
	data, err := json.Marshal(&resp)
	log.Println(string(data))
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(data)
}

func main() {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
	setUpHandler()
}
