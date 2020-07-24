package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	// Template
	defer fileClose()
	log.Println("7. 데이터 라벨링")
	log.Println("키의 존재 여부 확인")

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool

	value, ok = counters["a"]
	log.Println(value, ok)
	value, ok = counters["b"]
	log.Println(value, ok)
	value, ok = counters["c"]
	log.Println(value, ok)
}
