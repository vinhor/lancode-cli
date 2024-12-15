package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func langMenu(data JSONData) string {
	languageHighlight := color.New(color.FgCyan).SprintFunc()

	fmt.Print("Select language. h for help. 0 for abort: ")
	var language string
	fmt.Scanln(&language)
	if language == "h" {
		for _, language := range data.LanguagesNames {
			fmt.Printf("%s - %s\n", languageHighlight(language[0]), language[1])
		}
		fmt.Print("Select language.  0 for abort: ")
		fmt.Scanln(&language)
		checkLanguage(data.Languages, language, false)
	} else {
		checkLanguage(data.Languages, language, true)
	}
	if language == "0" {
		os.Exit(0)
	}

	fmt.Print("You have chosen ")
	color.Red(data.LanguagesNames[language][1])
	for index, concept := range data.Concepts[language] {
		fmt.Printf("%o - %s \n", index+1, concept)
	}

	return language
}

func conceptMenu(data JSONData, language string) (int, string) {
	fmt.Print("Select a concept. l for language menu. 0 for abort: ")
	var conceptIndexString string
	_, err := fmt.Scanln(&conceptIndexString)
	if err != nil {
		panic(err)
	}
	if conceptIndexString == "0" {
		os.Exit(0)
	}
	if conceptIndexString == "l" {
		language = langMenu(data)
		return conceptMenu(data, language)
	}

	conceptIndex, err := strconv.Atoi(conceptIndexString)
	if err != nil {
		panic(err)
	}

	return conceptIndex, language
}
