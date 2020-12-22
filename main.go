package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/RossJ24/connect474/modes"
)

// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	args := os.Args
	// Go through commend line args and determine what the user wants to do
	if len(args) == 1 || args[1] == "-h" || args[1] == "-help" {
		modes.Help()
		return
	}
	if args[1] == "--R" {
		modes.RandomPlay()
		return
	}
	if args[1] == "--PvP" {
		modes.PvP()
		return
	}
	if args[1] == "--PvC" {
		if args[2] == "MCTS" {
			modes.PvC(true)
			return
		} else if args[2] == "MMAB" {
			modes.PvC(false)
			return
		} else {
			modes.Help()
			return
		}
	}
	if args[1] == "--CvC" {
		if len(args) != 4 {
			modes.Help()
			return
		}
		games, _ := strconv.Atoi(args[3])
		if args[2] == "MCTS" {
			modes.CvC(true, games)
		} else if args[2] == "MMAB" {
			modes.CvC(false, games)
		} else if args[2] == "ALGOS" {
			modes.CVCAlgos(games)
		} else {
			modes.Help()
			return
		}
	}
}
