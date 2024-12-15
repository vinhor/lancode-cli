package main

func main() {
	data := openJson()
	language := langMenu(data)
	conceptIndex, language := conceptMenu(data, language)

	checkConceptIndex(len(data.Concepts[language]), int(conceptIndex))
	switch language {
	case "js":
		javaScriptConventions(data, conceptIndex)
	case "html":
		htmlConventions(data, conceptIndex)
	}
}
