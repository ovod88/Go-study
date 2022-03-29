package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------
const (
	maxTurns = 5
	usage    = `Enter possitive number like 2. %d iteration will be done`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me number")
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if guess <= 0 {
		fmt.Println("You gave negative number")
		return
	}
	if err != nil {
		fmt.Printf(usage, maxTurns)
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("You got from game %d\n", n)
		if n != guess {
			continue
		}

		if guess == n && turn == 0 {
			fmt.Println("You WON on first round!")
			return
		}
		fmt.Println("You WON!")
	}
	fmt.Println("You LOST! Try again!")
}
