package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/fatih/color"
)

func checkConceptIndex(maxNumber int, err error) {
	if err != nil {
		color.Red(fmt.Sprintf("Invalid concept index! Did you enter number between 0 and %v?", maxNumber))
		os.Exit(1)
	}
}

func checkLanguage(languagesArray []string, input string, canPrintHelp bool) {
	languagesArray = append(languagesArray, "0")
	if canPrintHelp {
		languagesArray = append(languagesArray, "h")
	}
	if !slices.Contains(languagesArray, input) {
		color.Red("Invalid language!")
		os.Exit(1)
	}
}