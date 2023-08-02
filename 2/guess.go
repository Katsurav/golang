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

	//** Gererate a Seed
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	//** Generate a Targe number
	target := rand.Intn(100) + 1
	//** fmt.Println(target)

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	//** Get Integer from the keyboard
	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//** Comparison
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
