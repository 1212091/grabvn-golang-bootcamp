package util

import (
	"regexp"
	"strings"
	"sync"
)

// Slice reads the content of file and put it in the word map
func Slice(fileContent string, wordMapChannel chan map[string]int, wg *sync.WaitGroup) {
	wordsMap := make(map[string]int)
	var regex = `[^A-Za-z0-9 ]+`
	var re = regexp.MustCompile(regex)
	formattedFileContent := re.ReplaceAllString(fileContent, "")

	wordsList := strings.Split(formattedFileContent, " ")

	for _, word := range wordsList {
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
	}
	wordMapChannel <- wordsMap
	defer wg.Done()
}
