package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Age  int    `json:"age"`
}

var EmployeeDb map[string]*Employee

func init() {
	EmployeeDb = map[string]*Employee{}

	EmployeeDb["Bob"] = &Employee{"Bob", 01234, 21}
	EmployeeDb["Joy"] = &Employee{"Joy", 01233, 25}
	EmployeeDb["David"] = &Employee{"David", 3223, 59}
}

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome"))
}

func GetEmployee(w http.ResponseWriter, req *http.Request, pa httprouter.Params) {
	qName := pa.ByName("name")

	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)

	if info, ok = EmployeeDb[qName]; !ok {
		w.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}

	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
		return
	}

	w.Write(infoJson)
}

func main() {

	router := httprouter.New()

	router.GET("/", Index)

	router.GET("/employee/:name", GetEmployee)

	log.Fatal(http.ListenAndServe(":9090", router))

}
