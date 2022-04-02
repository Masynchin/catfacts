package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CatFacts is the service that provides cat facts.
type CatFacts struct {
	url   string
	count int
}

// NewCatFacts instantiates CatFacts.
func NewCatFacts(count int) CatFacts {
	url := fmt.Sprintf(
		"https://cat-fact.herokuapp.com/facts/random?amount=%v",
		count,
	)
	return CatFacts{url, count}
}

// Get returns cat facts.
func (f CatFacts) Get() ([]Fact, error) {
	if *factsCount < 1 {
		return nil, errors.New("Facts count must be positive number")
	}

	resp, err := http.Get(f.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Server did not respond correctly")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return unmarshalFacts(data, f.count)
}

// unmarshalFacts fetches facts from API response.
func unmarshalFacts(data []byte, count int) ([]Fact, error) {
	if count == 1 {
		return unmarshalSingleFact(data)
	} else {
		return unmarshalManyFacts(data)
	}
}

// unmarshalSingleFact converts API response with single fact
// into array of facts.
func unmarshalSingleFact(data []byte) ([]Fact, error) {
	var f Fact
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	return []Fact{f}, nil
}

// unmarshalManyFacts converts API response with many facts
// into array of facts.
func unmarshalManyFacts(data []byte) ([]Fact, error) {
	var facts []Fact
	if err := json.Unmarshal(data, &facts); err != nil {
		return nil, err
	}
	return facts, nil
}
