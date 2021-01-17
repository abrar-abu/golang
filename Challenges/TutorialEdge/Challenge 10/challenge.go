package main

import (
  "fmt"
  "sort"
  "strings"
)

type Word struct {
	Word      string
	Frequency int
}

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range strings.Split(text, " ") {
		if _, ok := wordCount[word]; ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}

func Top5Words(wordmap map[string]int) []Word {
	words := []Word{}
	for word, frequency := range wordmap {
		words = append(words, Word{Word: word,
			Frequency: frequency})
	}
	sort.Slice(words, func(i, j int) bool { return words[i].Frequency > words[j].Frequency })
	return words[:5]
}

func main() {
  fmt.Println("Word Frequency Test")

  text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

  results := CountWords(text)
  MostCommon := Top5Words(results)

  fmt.Println(MostCommon)
}
