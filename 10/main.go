package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

type point struct {
	x, y   int
	vx, vy int
}

func (p *point) add() {
	p.x = p.x + p.vx
	p.y = p.y + p.vy
}

func (p *point) sub() {
	p.x = p.x - p.vx
	p.y = p.y - p.vy
}

func (p *point) xy() xy {
	return xy{p.x, p.y}
}

type xy struct {
	x, y int
}

func bounds(points []*point) image.Rectangle {

	maxX, maxY := 0, 0
	minX, minY := 1<<50, 1<<50

	for _, p := range points {

		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}

	}

	return image.Rect(minX, minY, maxX, maxY)

}

func main() {

	points := []*point{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		var x, y, vx, vy int

		_, err := fmt.Sscanf(
			line,
			"position=<%d,  %d> velocity=<%d, %d>",
			&x, &y, &vx, &vy,
		)

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, &point{x, y, vx, vy})

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Loop through seconds until the minimum bounding box
	// has been found

	lastSize := 1 << 50
	seconds := 0

	for {

		b := bounds(points)
		size := b.Size().X * b.Size().Y

		if size > lastSize {
			seconds--
			break
		}

		for _, p := range points {
			p.add()
		}

		lastSize = size
		seconds++

	}

	bitmap := make(map[xy]bool)

	for _, p := range points {
		p.sub()
		bitmap[xy{p.x, p.y}] = true
	}

	b := bounds(points)

	for y := 0; y <= b.Size().Y; y++ {
		for x := 0; x <= b.Size().X; x++ {

			if bitmap[xy{x + b.Min.X, y + b.Min.Y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}

		}

		fmt.Printf("\n")

	}

	fmt.Printf("\n")
	fmt.Printf("Waited %v seconds\n", seconds)

}
