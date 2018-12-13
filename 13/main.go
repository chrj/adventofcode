package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type direction struct {
	x, y int
}

var verbose = flag.Bool("verbose", false, "Show map for each iteration")

var (
	UP    = direction{0, -1}
	DOWN  = direction{0, 1}
	LEFT  = direction{-1, 0}
	RIGHT = direction{1, 0}
)

var turns = map[byte]map[direction]direction{

	'\\': map[direction]direction{
		UP:    LEFT,
		DOWN:  RIGHT,
		LEFT:  UP,
		RIGHT: DOWN,
	},

	'/': map[direction]direction{
		UP:    RIGHT,
		DOWN:  LEFT,
		LEFT:  DOWN,
		RIGHT: UP,
	},
}

var carts = map[direction]byte{

	UP:    '^',
	DOWN:  'v',
	LEFT:  '<',
	RIGHT: '>',
}

var intersections = [3]map[direction]byte{

	map[direction]byte{
		UP:    '\\',
		DOWN:  '\\',
		LEFT:  '/',
		RIGHT: '/',
	},

	map[direction]byte{
		UP:    '|',
		DOWN:  '|',
		LEFT:  '-',
		RIGHT: '-',
	},

	map[direction]byte{
		UP:    '/',
		DOWN:  '/',
		LEFT:  '\\',
		RIGHT: '\\',
	},
}

type xy struct {
	x, y int
}

func (loc xy) add(d direction) xy {
	return xy{
		loc.x + d.x,
		loc.y + d.y,
	}
}

type train struct {
	dir   direction
	turns int
}

func display(m map[xy]byte, trains map[xy]train, sX, sY int) {

	for y := 0; y <= sY; y++ {
		for x := 0; x <= sX; x++ {

			loc := xy{x, y}

			if t, found := trains[loc]; found {
				fmt.Printf("%s", string(carts[t.dir]))
				continue
			}

			if token, found := m[loc]; found {
				fmt.Printf("%s", string(token))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

// Simulate a step, return coordinates for any collisions
func sim(m map[xy]byte, trains map[xy]train) []xy {

	collisions := []xy{}

	keys := []xy{}
	for k := range trains {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {

		if keys[i].y != keys[j].y {
			return keys[i].y < keys[j].y
		}

		return keys[i].x < keys[j].x

	})

	for _, loc := range keys {

		t, found := trains[loc]
		if !found {
			continue
		}

		delete(trains, loc)

		nextLoc := loc.add(t.dir)
		nextTrack := m[nextLoc]
		nextDir := t.dir

		// Handle intersections

		if nextTrack == '+' {

			turnNo := t.turns % 3
			nextTrack = intersections[turnNo][t.dir]
			t.turns++

		}

		// Handle turns

		if turnDir, turn := turns[nextTrack]; turn {
			nextDir = turnDir[t.dir]

		}

		// Check if there's already a train in the new location

		if _, found := trains[nextLoc]; found {
			collisions = append(collisions, xy{nextLoc.x, nextLoc.y})
			delete(trains, nextLoc)
		} else {
			trains[nextLoc] = train{
				dir:   nextDir,
				turns: t.turns,
			}
		}

	}

	return collisions

}

func main() {

	flag.Parse()

	m := map[xy]byte{}  // system map
	t := map[xy]train{} // train locations

	var x, y int

	var sX, sY int // map size

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {

		if y > sY {
			sY = y
		}

		if x > sX {
			sX = x
		}

		token := scanner.Text()[0]

		if token == '\n' {

			x = 0
			y++

			continue

		}

		if idx := strings.IndexByte("<>^v", token); idx > -1 {

			t[xy{x, y}] = train{
				dir:   []direction{LEFT, RIGHT, UP, DOWN}[idx],
				turns: 0,
			}

			if idx < 2 {
				token = '-'
			} else {
				token = '|'
			}

		}

		if token == '-' || token == '|' || token == '/' || token == '\\' || token == '+' {
			m[xy{x, y}] = token
		}

		x++

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	if *verbose {
		display(m, t, sX, sY)
		fmt.Println()
	}

	for {

		collisions := sim(m, t)
		if len(collisions) > 0 {
			fmt.Printf("collision at: %v,%v (%v train(s) remain)\n", collisions[0].x, collisions[0].y, len(t))
		}

		if len(t) == 1 {
			for loc := range t {
				fmt.Printf("last train at: %v,%v\n", loc.x, loc.y)
			}
			break
		}

		if *verbose {
			display(m, t, sX, sY)
			fmt.Println()
		}

	}

}
