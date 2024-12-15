package main

import (
	"encoding/json"
	"os"
)

type JSONData struct {
	Concepts       map[string][]string `json:"concepts"`
	LanguagesNames map[string][]string `json:"languagesNames"`
	Languages      []string            `json:"languages"`
	JavaScript     map[string]string   `json:"js"`
	HTML           map[string]string   `json:"html"`
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
