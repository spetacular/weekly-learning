package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var model = make(map[string]int)
var alphalet = "abcdefghijklmnopqrstuvwxyz"

func edit1(word string) []string {
	s := make([]string, 1)
	wordLen := len(word)
	//add 1 word
	for i := range word {
		for _, c := range alphalet {
			s = append(s, word[0:i]+string(c)+word[i:])
		}
	}

	//delete 1 word
	for i := range word {
		s = append(s, word[0:i]+word[i+1:])
	}

	//change 1 word
	for i := range word {
		for _, c := range alphalet {
			s = append(s, word[0:i]+string(c)+word[i+1:])
		}
	}

	// change position
	for i := range word[0 : wordLen-1] {
		s = append(s, word[0:i]+string(word[i+1])+string(word[i])+word[i+2:])
	}

	return s
}

func candidates(word string) []string {
	s := make([]string, 1)
	s = append(s, word)
	s1 := append(s, edit1(word)...)

	return s1
}

func calc(word string) string {
	words := candidates(word)
	// fmt.Println(words)
	var maxProb = 1
	var correctWord = word

	for _, w := range words {
		prob, prs := model[w]
		if prs == true {
			if prob > maxProb {
				maxProb = prob
				correctWord = w
			}
		}
	}
	return correctWord
}

func main() {
	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		model[word]++
		// fmt.Println(word)
	}

	fmt.Println(calc("tae"))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
