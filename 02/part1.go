package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(counts map[byte]int) (int, int) {

	two := 0
	three := 0

	for _, c := range counts {
		switch c {
		case 2:
			two = 1
		case 3:
			three = 1
		}
	}

	return two, three

}

func main() {

	twos := 0
	threes := 0
	counts := map[byte]int{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {

		ord := uint8(byte(scanner.Text()[0]))

		switch {

		case ord == 10: // newline

			two, three := check(counts)
			twos += two
			threes += three

			counts = map[byte]int{}

		case ord >= 97 && ord <= 122:

			counts[ord] += 1

		}

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	two, three := check(counts)
	twos += two
	threes += three

	fmt.Printf("checksum: %v", twos*threes)

}
