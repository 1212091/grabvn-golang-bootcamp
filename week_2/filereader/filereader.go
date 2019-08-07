package filereader

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFiles(files []os.FileInfo) chan string {
	input := make(chan string, 100)
	go readFilesAsync(files, input)
	return input
}

func readFilesAsync(files []os.FileInfo, input chan string) {
	for _, file := range files {
		content, err := ioutil.ReadFile("./wiki/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		input <- string(content)
	}
	close(input)
}
