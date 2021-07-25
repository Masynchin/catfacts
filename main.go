package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	factsCount int
	urlTemplate = "https://cat-fact.herokuapp.com/facts/random?amount=%v"
)

func init() {
	flag.IntVar(&factsCount, "n", 5, "facts count, cannot be lower than 1")
}

type Fact struct {
	Text string `json:"text"`
}

func main() {
	flag.Parse()
	if factsCount < 1 {
		log.Fatal(errors.New("Facts count must be positive number"))
	}

	resp, err := makeRequest()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(errors.New("Server did not respond correctly"))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if factsCount == 1 {
		printSingleFact(data)
	} else {
		printManyFacts(data)
	}
}

func makeRequest() (*http.Response, error) {
	url := fmt.Sprintf(urlTemplate, factsCount)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func printSingleFact(data []byte) {
	var f Fact
	if err := json.Unmarshal(data, &f); err != nil {
		log.Fatal(err)
	}
	printFact(f)
}

func printManyFacts(data []byte) {
	var facts []Fact
	if err := json.Unmarshal(data, &facts); err != nil {
		log.Fatal(err)
	}
	for _, f := range facts {
		printFact(f)
	}
}

func printFact(f Fact) {
	fmt.Println(f.Text)
}
