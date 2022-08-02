package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "this is test page")
	})

	http.HandleFunc("/time/", func(w http.ResponseWriter, req *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		io.WriteString(w, timeStr)
	})

	http.ListenAndServe(":9090", nil)
}
