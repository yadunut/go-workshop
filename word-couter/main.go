package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	filename = "wordlist.txt"
)

func check(err error) {
	if err != nil {
		fmt.Println("error: " + err.Error())
		os.Exit(1)
	}
}

func main() {

	bb, err := ioutil.ReadFile(filename)
	check(err)
	wordMap := make(map[string]int)

	file := string(bb)

	words := strings.Split(file, " ")
	for _, word := range words {
		wordMap[word]++
	}

	for word, count := range wordMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}
