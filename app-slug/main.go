package main

import (
	"log"

	"github.com/resonantChaos22/toolkit"
)

func main() {
	toSlug := "Now is the time 123 "

	var tools toolkit.Tools

	slug, err := tools.Slugify(toSlug)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(slug)
}
