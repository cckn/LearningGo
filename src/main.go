package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")

	fmt.Println("Target is ", target)
}
