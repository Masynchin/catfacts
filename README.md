# catfacts
CLI that fetches cat facts from cat-fact.herokuapp.com (or as [github repo](https://github.com/alexwohlbruck/cat-facts))

## Installing

You can install this program with command `go get github.com/Masynchin/catfacts`

## Usage

### General

You can run `catfacts` without any options provided and get 5 random cat facts

### Options

With `-n {positive number}` option you can get as many cat facts as you want. It can be 1, 2, 10 or even 500 (according to [API limit](https://alexwohlbruck.github.io/cat-facts/docs/endpoints/facts#query-parameters))
