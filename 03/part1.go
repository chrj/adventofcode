package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func eq(s1, s2 string) string {

	result := ""

	for i, c := range s1 {

		if i < len(s2) && byte(c) == s2[i] {
			result += string(c)
		}
	}

	return result

}

func main() {

	fabric := map[int]map[int]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, " @ ")
		parts = strings.Split(parts[1], ": ")

		coordinate := parts[0]
		size := parts[1]

		parts = strings.Split(coordinate, ",")

		x0, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		y0, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		parts = strings.Split(size, "x")

		w, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		h, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		for x := x0; x < x0+w; x++ {

			if fabric[int(x)] == nil {
				fabric[int(x)] = map[int]int{}
			}

			for y := y0; y < y0+h; y++ {
				fabric[int(x)][int(y)]++
			}
		}

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	overlapping := 0

	for _, ys := range fabric {
		for _, count := range ys {
			if count > 1 {
				overlapping++
			}
		}
	}

	fmt.Printf("square inches overlapping: %v\n", overlapping)

}
