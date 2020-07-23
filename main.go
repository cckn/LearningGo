package main

import log "github.com/sirupsen/logrus"

// import log "github.com/sirupsen/logrus"

func main() {

	defer fileClose()

	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	log.Println(notes[3], notes[6], notes[0])
}
