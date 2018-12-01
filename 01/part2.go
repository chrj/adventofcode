package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	deltas := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		deltas = append(deltas, int(i))

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	frequencies := map[int]struct{}{}
	freq := 0

	for i := 0; ; i++ {

		freq += deltas[i%len(deltas)]
		if _, seen := frequencies[freq]; seen {
			break
		}

		frequencies[freq] = struct{}{}

	}

	fmt.Printf("first repeated: %v\n", freq)

}
