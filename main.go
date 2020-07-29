package main

import (
	log "github.com/sirupsen/logrus"
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("error interface")

	var err error
	err = ComedyError("What's a programmer's favorite beer? Logger!")

	log.Println(err)
}
