package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, pa httprouter.Params) {

	str := fmt.Sprintf("hello %s!\n", pa.ByName("name"))
	io.WriteString(w, str)
}

func main() {
	router := httprouter.New()

	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":9090", router))
}
