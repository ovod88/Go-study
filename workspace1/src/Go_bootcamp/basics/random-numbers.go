package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `Enter possitive number like 2. %d iteration will be done`
)

var (
	win_messages  = []string{"You WON", "You awesome", "You cool"}
	lose_messages = []string{"You LOSE", "You bad", "No worries"}
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
		} else {
			fmt.Println(win_messages[rand.Intn(len(win_messages)-1)])
			return
		}
	}
	fmt.Println(lose_messages[rand.Intn(len(lose_messages)-1)])
}
