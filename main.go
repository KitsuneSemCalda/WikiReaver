package main

import (
	"WikiReaver/linkGenerator/url"
	urltest "WikiReaver/linkGenerator/urlTest"
	"flag"
	"fmt"
	"log"
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

	url := url.MountFullUrl(*lang, *term)
	isOk := urltest.CheckUrl(url)

	if isOk == nil {
		log.Println("Este link está disponivel: ", url)
		return
	}
}
