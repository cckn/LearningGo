package main

import (
	log "github.com/sirupsen/logrus"
)

type MyType string

func (receiver MyType) sayHi() {

	log.Println(receiver)
}

func main() {

	// Template
	defer fileClose()
	log.Println("8. 나만의 타입 - 사용자 정의 타입")
	log.Println("메서드")

	value := MyType("a MyType value")
	value.sayHi()
	value2 := MyType("a")
	value2.sayHi()
}
