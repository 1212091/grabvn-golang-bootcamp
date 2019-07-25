package main

import (
	"io/ioutil"
	"log"
	"os"

	"./filereader"
	"./wordmap"
)

func main() {
	files, err := ioutil.ReadDir("./wiki")
	if err != nil {
		log.Fatal(err)
	}
	processData(files)
}

func processData(files []os.FileInfo) {
	// create pineline
	chanInput := filereader.ReadFiles(files)
	wordMapChannel1 := wordmap.CreateWordMapFromFileContent(chanInput)
	wordMapChannel2 := wordmap.CreateWordMapFromFileContent(chanInput)
	mergedWordMapChannel := wordmap.MergeWordMapChannels(wordMapChannel1, wordMapChannel2)
	finalWordsMap := wordmap.SumWordInAllFiles(mergedWordMapChannel)
	wordmap.PrintFinalWordsMap(finalWordsMap)
}
