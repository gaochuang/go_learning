package json_test

import (
	"encoding/json"
	"testing"
)

type PodStatus struct {
	Name         string `json:"name"`
	StartTime    string `json:"starttime"`
	RestartCount int    `json:"restartcount"`
}

type PodSpec struct {
	Image    string `json:"image"`
	NodeName string `json:"nodename"`
}

type PodStruct struct {
	PodStatus PodStatus `json:"PodStatus"`
	PodSpec   PodSpec   `json:"PodSpec"`
}

var PodStr = `{
	"PodStatus":{
      "name":"webapp",
	  "starttime":"2022-08-01",
	  "restartcount":0
	},
	"PodSpec":{
		"image":"docker.registry.webapp:v1",
		"nodename":"master-node"
	}
}`

func TestPodJson(t *testing.T) {
	pod1 := PodStruct{}
	err := json.Unmarshal([]byte(PodStr), &pod1)
	if err != nil {
		t.Error(err)
	}

	t.Log(pod1)

	if data, err := json.Marshal(&pod1); err == nil {
		t.Log(string(data))
	} else {
		t.Error(err)
	}
}
