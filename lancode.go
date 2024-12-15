package main

import (
	"fmt"
	"os"
)

func main() {
	data := openJson()

	fmt.Print("Select language. h for help. 0 for abort: ")
	var language string
	fmt.Scanln(&language)
	if language == "h" {
		for _, language := range data.LanguagesNames {
			fmt.Printf("%s - %s\n", language[0], language[1])
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

	fmt.Println("Select a concept. 0 for abort")
	for index, concept := range data.Concepts[language] {
		fmt.Printf("%o - %s \n", index+1, concept)
	}
	var conceptIndex int
	_, err := fmt.Scanln(&conceptIndex)
	checkConceptIndex(len(data.Concepts[language]), err, conceptIndex)
	switch language {
	case "js":
		javaScriptConventions(data, conceptIndex)
	case "html":
		htmlConventions(data, conceptIndex)
	}
}
