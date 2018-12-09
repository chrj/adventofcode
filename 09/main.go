package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"sort"
)

func sim(players, lastMarble int) int {

	scores := make([]int, players)

	marbles := ring.New(1)
	marbles.Value = 0

	for m := 1; m <= lastMarble; m++ {

		// calculate player number from marble value
		player := (m - 1) % players

		if m%23 == 0 {

			// add the marble to be placed to players score
			scores[player] += m

			// move pointer counter-clockwise
			marbles = marbles.Prev().Prev().Prev().Prev().Prev().Prev().Prev()

			// remove marble
			removed := marbles.Prev().Prev().Unlink(1)

			// add to score
			scores[player] += removed.Value.(int)

			// move pointer forwards
			marbles = marbles.Next()

		} else {

			// insert marble

			r := ring.New(1)
			r.Value = m

			marbles = marbles.Link(r)

		}

	}

	sort.Ints(scores)

	return scores[players-1]

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		var players, lastMarble int

		_, err := fmt.Sscanf(
			line,
			"%d players; last marble is worth %d points",
			&players, &lastMarble,
		)

		if err != nil {
			log.Fatal(err)
		}

		// Part 1 simulation

		highscore := sim(players, lastMarble)

		fmt.Printf("with players: %v, last marble worth: %v. Got highscore: %v\n",
			players, lastMarble, highscore)

		// Part 2 simulation

		lastMarble *= 100

		highscore = sim(players, lastMarble)

		fmt.Printf("with players: %v, last marble worth: %v. Got highscore: %v\n",
			players, lastMarble, highscore)

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
