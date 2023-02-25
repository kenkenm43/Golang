package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	inputYourname = "[your name]"
	feelingSyntax = `%s feels %s`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(inputYourname)
		return
	}
	name := args[0]
	moods := [...]string{
		"sad 😂",
		"happy 🙂",
		"terrible ",
		"good 👍",
		"awesome",
	}
	n := rand.Intn(len(moods))

	fmt.Printf(feelingSyntax, name, moods[n])
}
