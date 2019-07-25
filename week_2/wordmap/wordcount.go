package wordmap

import (
	"regexp"
	"strings"
)

// Slice reads the content of file and put it in the word map
func createWordMapFromFileContentAsync(input <-chan string, wordMapChannel chan map[string]int) {
	for fileContent := range input {
		wordsMap := make(map[string]int)
		regex := regexp.MustCompile(`[^A-Za-z0-9 ]+`)
		formattedFileContent := regex.ReplaceAllString(fileContent, "")

		wordsList := strings.Split(formattedFileContent, " ")

		for _, word := range wordsList {
			if _, ok := wordsMap[word]; ok {
				wordsMap[word]++
			} else {
				wordsMap[word] = 1
			}
		}
		wordMapChannel <- wordsMap
	}
	close(wordMapChannel)
}

func CreateWordMapFromFileContent(input <-chan string) <-chan map[string]int {
	wordMapChannel := make(chan map[string]int, 100)
	go createWordMapFromFileContentAsync(input, wordMapChannel)
	return wordMapChannel
}
