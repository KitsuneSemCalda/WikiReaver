package main

import (
	"WikiReaver/linkGenerator/url"
	"flag"
	"fmt"
)

func main() {
	lang := flag.String("lang", "", "language to be searched default is 'en'")
	term := flag.String("term", "", "term to be searched default is null")

	flag.Parse()

	if *term == "" {
		fmt.Println("Please give a term to be searched")
		return
	}

	if *lang != "" {
		url.MountFullUrl(*lang, *term)
		return
	}

	url.MountFullUrl(*lang, *term)
}
