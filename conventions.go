package main

import (
	"fmt"

	"github.com/fatih/color"
)

func javaScriptConventions(json JSONData, conceptIndex int) {
	keyword := color.New(color.FgHiRed).SprintFunc()
	name := color.New(color.FgBlue).SprintFunc()
	number := color.New(color.FgMagenta).SprintFunc()

	concept := json.Concepts["js"][conceptIndex-1]
	convention := json.JavaScript[concept]
	fmt.Println("Use", convention)
	switch concept {
	case "variable":
		fmt.Printf("%s %s = %s\n", keyword("let"), name("numberOfEmployes"), number("52"))
	}
}
