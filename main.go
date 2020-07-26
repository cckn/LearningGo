package main

import (
	log "github.com/sirupsen/logrus"
)

type NoiseMaker interface {
	MakeSound()
}

type Whistle string

func (w Whistle) MakeSound() {
	log.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	log.Println("Honk!")
}

type Bell string

func (b Bell) MakeSound() {
	log.Println("Ring!")
}

func doubleSound(item NoiseMaker) {
	for i := 0; i < 2; i++ {
		item.MakeSound()
	}
}

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("인터페이스")

	var b1 = Bell("")
	doubleSound(b1)
	var b2 = Whistle("")
	doubleSound(b2)
	var b3 = Horn("")
	doubleSound(b3)
}
