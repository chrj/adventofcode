package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) String() string {
	return fmt.Sprintf("%v, %v", p.x, p.y)
}

func (p point) distance(q point) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

func main() {

	points := []point{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), ", ")

		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point{int(x), int(y)})
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Find bounds

	var xMax, yMax, xMin, yMin = 0, 0, 500, 500

	for _, p := range points {

		if p.x < xMin {
			xMin = p.x
		}

		if p.x > xMax {
			xMax = p.x
		}

		if p.y < yMin {
			yMin = p.y
		}

		if p.y > yMax {
			yMax = p.y
		}

	}

	sizes := make(map[point]int, len(points)) // size of the area that point is the closest point for (-1 for infite)

	var withinRegion = 0

	for x := xMin; x <= xMax; x++ {

		for y := yMin; y <= yMax; y++ {

			var closest point
			var distance = 1000
			var count = 0

			var totalDistance = 0

			for _, p := range points {

				d := p.distance(point{x, y})

				totalDistance += d

				if d == distance {
					count++
					continue
				}

				if d > distance {
					continue
				}

				closest = p
				distance = d
				count = 1

			}

			if totalDistance < 10000 {
				withinRegion++
			}

			if count != 1 {
				// Several points have equal distance to this coordinate
				continue
			}

			if x == xMin || x == xMax || y == yMin || y == yMax {
				sizes[closest] = -1 // infinite
			} else if sizes[closest] >= 0 {
				sizes[closest]++
			}

		}

	}

	var maxSize = 0
	var maxPoint point

	for p, size := range sizes {

		if size > maxSize {
			maxSize = size
			maxPoint = p
		}

	}

	fmt.Printf("Size of the largest area that isn't infinite: %v (for point: %s)\n", maxSize, maxPoint)
	fmt.Printf("Size of the region containing all locations which have a total distance to all given coordinates of less than 10000: %v\n", withinRegion)

}
