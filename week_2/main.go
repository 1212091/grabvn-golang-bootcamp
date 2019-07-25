package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"./util"
)

func main() {
	files, err := ioutil.ReadDir("./wiki")
	if err != nil {
		log.Fatal(err)
	}

	wordMapChannel := make(chan map[string]int)
	var wg sync.WaitGroup

	readFilesInFolder(files, &wg, wordMapChannel)

	go closeChannel(&wg, wordMapChannel)
	finalWordsMap := sumWordInAllFiles(wordMapChannel)
	printFinalWordsMap(finalWordsMap)
}

func readFilesInFolder(files []os.FileInfo, wg *sync.WaitGroup, wordMapChannel chan map[string]int) {
	for _, f := range files {
		data, err := ioutil.ReadFile("./wiki/" + f.Name())
		if err != nil {
			log.Fatal(err)
			return
		}
		wg.Add(1)
		go util.Slice(string(data), wordMapChannel, wg)
	}
}

func sumWordInAllFiles(wordMapChannel <-chan map[string]int) map[string]int {
	finalWordsMap := make(map[string]int)
	for wordMap := range wordMapChannel {
		for word, frequency := range wordMap {
			if _, ok := finalWordsMap[word]; ok {
				finalWordsMap[word] += frequency
			} else {
				finalWordsMap[word] = frequency
			}
		}
	}
	return finalWordsMap
}

func printFinalWordsMap(finalWordsMap map[string]int) {
	fmt.Println("Stop here")
	for word, frequency := range finalWordsMap {
		fmt.Println("word: ", word, " frequency: ", frequency)
	}
}

func closeChannel(wg *sync.WaitGroup, wordMapChannel chan map[string]int) {
	fmt.Println("Wait here")
	wg.Wait()
	close(wordMapChannel)
}
