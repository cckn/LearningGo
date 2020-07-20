package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")

	fmt.Println("Target is ", target)

	for true {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Make Guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess == target {
			fmt.Println("Correct!")
			return
		} else if guess > target {
			fmt.Println("less")
		} else if guess < target {
			fmt.Println("more")
		}
	}

	fmt.Println(target)

}
