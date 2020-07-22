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
	var second = time.Now().Unix()
	rand.Seed(second)

	var target = rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	var reader = bufio.NewReader(os.Stdin)

	var success = false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println(fmt.Sprintf("You have %d guesses left", 10-guesses))
		fmt.Printf("Make a guess : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("GJ! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println(fmt.Sprintf("Sorry, you didn't guess my number. It was: %d", target))
	}
}
