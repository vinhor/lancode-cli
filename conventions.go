package main

import (
	"fmt"

	"github.com/fatih/color"
)

func javaScriptConventions(json JSONData, conceptIndex int) {
	keyword := color.New(color.FgHiGreen).SprintFunc()
	specialOperator := color.New(color.FgHiYellow).SprintFunc()
	name := color.New(color.FgHiBlue).SprintFunc()
	number := color.New(color.FgMagenta).SprintFunc()
	function := color.New(color.FgHiCyan).SprintFunc()
	bracket := color.New(color.FgBlue).SprintFunc()
	operator := color.New(color.FgCyan).SprintFunc()
	specialKeyword := color.New(color.FgHiRed).SprintFunc()
	stringColor := color.New(color.FgYellow).SprintFunc()
	class := color.New(color.FgHiMagenta).SprintFunc()

	concept := json.Concepts["js"][conceptIndex-1]
	convention := json.JavaScript[concept]
	fmt.Println("Use", convention)
	switch concept {
	case "constant":
		fmt.Printf("%s %s = %s\n", keyword("const"), name("PRODUCT_MANAGER"), stringColor("\"Anna Mayer\""))
	case "variable":
		fmt.Printf("%s %s = %s\n", keyword("let"), name("numberOfEmployes"), number("52"))
	case "function":
		fmt.Printf("%s %s = %s %s %s\n", keyword("const"), function("increaseEmployeeCounter"), bracket("()"), specialOperator("=>"), bracket("{"))
		fmt.Printf("  %s %s%s\n", specialKeyword("return"), name("numberOfEmployes"), operator("++"))
		fmt.Printf("%s\n", bracket("}"))
	case "class":
		fmt.Printf("%s %s %s\n", keyword("class"), class("SeniorEngineer"), bracket("{"))
		fmt.Printf("  %s%s%s%s\n", specialKeyword("constructor"), bracket("("), name("name"), bracket(") {"))
		fmt.Printf("    %s.%s = %s\n", keyword("this"), name("name"), name("name"))
		fmt.Printf("    %s.%s = %s\n", keyword("this"), name("level"), number("6"))
		fmt.Printf("  %s\n", bracket("}"))
		fmt.Printf("%s\n", bracket("}"))
	}
}
