package wordmap

import "sync"

func MergeWordMapChannels(wordMapChannels ...<-chan map[string]int) <-chan map[string]int {
	var waitGroup sync.WaitGroup

	mergedWordMapChannel := make(chan map[string]int, 100)

	waitGroup.Add(len(wordMapChannels))

	for _, wordMapChannel := range wordMapChannels {
		go pushDataToMergedChannel(&waitGroup, wordMapChannel, mergedWordMapChannel)
	}

	go closeMergedChannel(&waitGroup, mergedWordMapChannel)

	return mergedWordMapChannel
}

func SumWordInAllFiles(wordMapChannel <-chan map[string]int) map[string]int {
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

func pushDataToMergedChannel(waitGroup *sync.WaitGroup, wordMapChannel <-chan map[string]int, mergedWordMapChannel chan map[string]int) {
	for wordMap := range wordMapChannel {
		mergedWordMapChannel <- wordMap
	}
	waitGroup.Done()
}

func closeMergedChannel(waitGroup *sync.WaitGroup, wordMapChannel chan map[string]int) {
	waitGroup.Wait()
	close(wordMapChannel)
}
