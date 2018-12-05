package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func react(p []byte) []byte {

	i := 0

	for {

		if p[i] > p[i+1] && p[i]-p[i+1] == 32 ||
			p[i] < p[i+1] && p[i+1]-p[i] == 32 {

			p = append(p[:i], p[i+2:]...)

			if i > 0 {
				i--
			}

		} else {
			i++
		}

		if i >= len(p)-1 {
			break
		}

	}

	return p

}

func remove(p []byte, u byte) []byte {

	i := 0

	for {

		if p[i] == u || p[i] == u+32 {

			p = append(p[:i], p[i+1:]...)

		} else {
			i++
		}

		if i >= len(p)-1 {
			break
		}

	}

	return p

}

func main() {

	polymer, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	polymer = bytes.TrimSpace(polymer)

	// Part 1

	reacted := make([]byte, len(polymer))
	copy(reacted, polymer)
	reacted = react(reacted)

	fmt.Printf("Units remaining after reactions: %v\n", len(reacted))

	// Part 2

	best := len(reacted)
	bestunit := 0

	for i := 32; i <= 90; i++ {

		cleaned := make([]byte, len(polymer))
		copy(cleaned, polymer)
		cleaned = react(remove(cleaned, byte(i)))

		if len(cleaned) < best {
			best = len(cleaned)
			bestunit = i
		}

	}

	fmt.Printf("Unit to be removed: %v, new length: %v\n", string(bestunit), best)

}
