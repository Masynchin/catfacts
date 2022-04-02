package main

import (
	"flag"
	"log"
)

var factsCount = flag.Int("n", 5, "facts count, cannot be lower than 1")

func main() {
	flag.Parse()

	facts, err := NewCatFacts(*factsCount).Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, f := range facts {
		f.Print()
	}
}
