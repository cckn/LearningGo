package main

import log "github.com/sirupsen/logrus"

// import log "github.com/sirupsen/logrus"

func main() {

	defer fileClose()

	var notes []string
	notes = make([]string, 7)

	log.Print(notes)

}
