package wordmap

import "fmt"

func PrintFinalWordsMap(finalWordsMap map[string]int) {
	for word, frequency := range finalWordsMap {
		fmt.Println("Word: ", word, " | frequency: ", frequency)
	}
}
