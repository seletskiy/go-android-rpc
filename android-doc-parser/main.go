package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type Step string

const (
	stepStart       Step = "stepStart"
	stepClassNav    Step = "stepClassNav"
	stepEnd         Step = "stepEnd"
	stepClassesList Step = "stepClassesList"
	stepClassLink   Step = "stepClassLink"
	stepUnknown     Step = "stepUnknown"
)

type Class struct {
	Link string
	Name string
}

type tokenizerState struct {
	Step       Step
	TagName    string
	TokenType  html.TokenType
	Attributes map[string]string
}

var classes []Class

func main() {
	reader, err := os.Open("package-summary.html")
	if err != nil {
		panic(err)
	}

	tokenizer := html.NewTokenizer(reader)

	state := tokenizerState{
		Step: stepStart,
	}

	for state.Step != "stepEnd" {
		tokenType := tokenizer.Next()
		tagName, hasAttr := tokenizer.TagName()
		state.TagName = string(tagName)
		state.TokenType = tokenType
		state.Attributes = getAttributes(hasAttr, tokenizer)

		state.Step = switchState(state, tokenizer)
	}

	for _, class := range classes {
		fmt.Printf("%60s %s\n", class.Name, class.Link)
	}
}

func switchState(
	state tokenizerState, tokenizer *html.Tokenizer,
) Step {
	//debug(state)

	if state.TokenType == html.ErrorToken {
		return "stepEnd"
	}

	switch state.Step {
	case stepStart:
		return parseClassNav(state, tokenizer)
	case stepClassNav:
		return parseClassesBegin(state, tokenizer)
	case stepClassesList:
		return parseClassesList(state, tokenizer)
	case stepClassLink:
		return parseClassLink(state, tokenizer)
	}

	return stepUnknown
}

func parseClassNav(
	state tokenizerState, tokenizer *html.Tokenizer,
) Step {
	nextStep := state.Step

	switch state.TokenType {
	case html.StartTagToken:
		if state.Attributes["id"] == "classes-nav" {
			nextStep = stepClassLink
		}
	}

	return nextStep
}

func parseClassesBegin(
	state tokenizerState, tokenizer *html.Tokenizer,
) Step {
	nextStep := state.Step

	switch state.TokenType {
	case html.TextToken:
		text := string(tokenizer.Text())
		if text == "Classes" {
			nextStep = stepClassesList
		}
	}

	return nextStep
}

func parseClassesList(
	state tokenizerState, tokenizer *html.Tokenizer,
) Step {
	nextStep := state.Step

	switch state.TokenType {
	case html.StartTagToken:
		if state.Attributes["class"] == "jd-linkcol" {
			nextStep = stepClassLink
		}
	}

	return nextStep
}

func parseClassLink(
	state tokenizerState, tokenizer *html.Tokenizer,
) Step {
	nextStep := state.Step

	switch state.TokenType {
	case html.StartTagToken:
		if state.TagName == "a" {
			href := state.Attributes["href"]
			tokenizer.Next()
			name := string(tokenizer.Text())
			classes = append(classes, Class{
				Link: href,
				Name: name,
			})
		}
	case html.EndTagToken:
		nextStep = stepClassesList
	}

	return nextStep
}

func getAttributes(hasAttr bool, tokenizer *html.Tokenizer) map[string]string {
	attributes := make(map[string]string)

	for hasAttr {
		key, value, hasMoreAttr := tokenizer.TagAttr()

		attributes[string(key)] = string(value)

		if !hasMoreAttr {
			break
		}
	}

	return attributes
}

func debug(state tokenizerState) {
	switch state.TokenType {
	case html.StartTagToken:
		log.Printf("%s <%s>", state.Step, state.TagName)
	case html.EndTagToken:
		log.Printf("%s </%s>", state.Step, state.TagName)
	}
}
