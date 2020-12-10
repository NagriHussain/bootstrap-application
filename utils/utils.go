package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//Now returns a String format current time
func Now() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func HTTPGETCall(u string) []byte {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return body
}
