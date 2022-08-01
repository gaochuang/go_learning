package json_test

import (
	"encoding/json"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:age`
}

type JonInfo struct {
	Skill []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JonInfo   JonInfo   `json:"jon_info"`
}

var jsonStr = `{
  "basic_info":{
	 "name":"Mike",
   "age":30
  },
  "jon_info":{
	"skills":["Java","Go","C"]
  }
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*e)

	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}

}
