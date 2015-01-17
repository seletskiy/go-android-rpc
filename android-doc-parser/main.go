package main

import (
	"fmt"
	"os"

	"launchpad.net/xmlpath"
)

type Class struct {
	Link string
	Name string
}

var xpathClassA = xmlpath.MustCompile(`//td[@class='jd-linkcol']/a`)
var xpathClassAHref = xmlpath.MustCompile(`@href`)

func main() {
	reader, err := os.Open("package-summary.html")
	if err != nil {
		panic(err)
	}

	root, err := xmlpath.ParseHTML(reader)
	if err != nil {
		panic(err)
	}

	classes := []Class{}

	linksIterator := xpathClassA.Iter(root)
	for linksIterator.Next() {
		node := linksIterator.Node()
		href, _ := xpathClassAHref.String(node)
		classes = append(classes, Class{
			Link: href,
			Name: node.String(),
		})
	}

	fmt.Printf("%+v\n", classes)
}
