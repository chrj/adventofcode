package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	freq := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		freq += int(i)

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("frequency: %v\n", freq)

}
