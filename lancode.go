package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	data := openJson()
	fmt.Println("Select language. h for help. 0 for abort")
	var language string
	fmt.Scanln(&language)
	if language == "h" {
		printHelp()
		fmt.Println("Select a language. 0 for abort")
		fmt.Scanln(&language)
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
	if err != nil {
		color.Red(fmt.Sprintf("Invalid concept index! Did you enter number between 0 and %v?", len(data.Concepts[language])))
	}
	if conceptIndex == 0 {
		os.Exit(0)
	}
	if conceptIndex > len(data.Concepts[language]) || conceptIndex < 1 {
		color.Red(fmt.Sprintf("Invalid concept index! Did you enter number between 0 and %v?", len(data.Concepts[language])))
		os.Exit(1)
	}
}
