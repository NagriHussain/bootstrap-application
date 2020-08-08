package utils

import (
	"io/ioutil"
	"net/http"
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
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return body
}
