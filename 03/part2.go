package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     int
	x0, y0 int
	w, h   int
}

func main() {

	fabric := map[int]map[int]int{}
	claims := []claim{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, " @ ")

		id, err := strconv.ParseInt(parts[0][1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

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

		claims = append(claims, claim{int(id), int(x0), int(y0), int(w), int(h)})

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	var nonoverlapping claim

outerloop:
	for _, c := range claims {
		for x := c.x0; x < c.x0+c.w; x++ {
			for y := c.y0; y < c.y0+c.h; y++ {
				if fabric[x][y] > 1 {
					continue outerloop
				}
			}
		}

		nonoverlapping = c
		break

	}

	fmt.Printf("non overlapping: %v\n", nonoverlapping.id)

}
