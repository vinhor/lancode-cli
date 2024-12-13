package main

import (
	"encoding/json"
	"os"
)

type JSONData struct {
	Concepts  map[string][]string          `json:"concepts"`
	Languages map[string]map[string]string `json:"-"`
}

func openJson() JSONData {
	file, err := os.ReadFile("main.json")
	if err != nil {
		panic(err)
	}

	var data JSONData
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	} else {
		return data
	}

}