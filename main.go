package main

import (
	"github.com/cckn/go-utils/gadget"
	log "github.com/sirupsen/logrus"
)

type Player interface {
	Play(string)
	Stop()
}

//
// func TryOut(player Player) {
// 	player.Play("Tast Track")
// 	player.Stop()
// 	recorder := player.(gadget.TapeRecorder)
// 	recorder.Record()
// }

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("타입단언")

	var player Player = gadget.TapePlayer{}
	recorder, ok := player.(gadget.TapeRecorder)

	if ok {
		recorder.Record()
	} else {
		log.Println("Player was not a TapeRecorder")
	}
}
