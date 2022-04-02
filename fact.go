package main

import "fmt"

// Fact is model representation of single fact from API's response.
type Fact struct {
	Text string `json:"text"`
}

// Print prints fact to console.
func (f Fact) Print() {
	fmt.Println(f.Text)
}
