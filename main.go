package main

import log "github.com/sirupsen/logrus"

type Number int

func (n *Number) Double() {
	*n *= 2
	log.Println(*n)
}

func main() {

	// Template
	defer fileClose()
	log.Println("포인터 리시버 매개 변수")

	var n = Number(10)
	log.Println(n)

	n.Double()
	log.Println(n)
}
