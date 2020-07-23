package main

import (
	"github.com/cckn/datafile"
	log "github.com/sirupsen/logrus"
)

// import log "github.com/sirupsen/logrus"

func main() {

	defer fileClose()

	slice, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Error(err)
	}

	log.Println(slice)
}
