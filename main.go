package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	// Template
	defer fileClose()
	log.Println("8. 구조체")

	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	log.Println(myStruct.number)
	log.Println(myStruct.word)
	log.Println(myStruct.toggle)
}
