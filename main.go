package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func readDir(path string, depth int) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		newPath := fmt.Sprintf("%s/%s", path, file.Name())

		if file.IsDir() {
			log.Println("File : ", newPath)
			readDir(newPath, depth+1)
		} else {
			log.Println("File : ", newPath)
		}
	}
}

func main() {
	defer fileClose()
	log.Println("파일 나열 하기")
	readDir(".", 0)
}
