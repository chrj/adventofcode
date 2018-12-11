package main

import (
	"fmt"
	"sync"
)

func grid(serial int) [][]int {

	g := make([][]int, 300)

	for x := 0; x < 300; x++ {

		g[x] = make([]int, 300)

		for y := 0; y < 300; y++ {

			x1 := x + 1 // 1-indexed
			y1 := y + 1 // 1-indexed

			// Find the fuel cell's rack ID, which is its X coordinate plus 10.
			rackID := x1 + 10

			// Begin with a power level of the rack ID times the Y coordinate.
			powerLevel := rackID * y1

			// Increase the power level by the value of the grid serial number (your puzzle input).
			powerLevel += serial

			// Set the power level to itself multiplied by the rack ID.
			powerLevel *= rackID

			// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
			powerLevel = (powerLevel / 100) % 10

			// Subtract 5 from the power level.
			powerLevel -= 5

			// Update grid
			g[x][y] = powerLevel

		}

	}

	return g

}

func search(grid [][]int, size int) (int, int, int) {

	best := 0
	bestX := 0
	bestY := 0

	for x := 0; x < (300 - size); x++ {
		for y := 1; y < (300 - size); y++ {

			total := 0

			for xx := x; xx < x+size; xx++ {
				for yy := y; yy < y+size; yy++ {
					total += grid[xx][yy]
				}
			}

			if total > best {

				best = total
				bestX = x
				bestY = y

			}

		}
	}

	return best, bestX, bestY // 0-indexed

}

type result struct {
	total, x, y, size int
}

func sizes(grid [][]int) (int, int, int) {

	// Search all sizes of the grid concurrently

	results := make(chan result)
	var wg sync.WaitGroup

	wg.Add(300)

	for size := 1; size <= 300; size++ {

		go func(size int) {
			total, x, y := search(grid, size)
			results <- result{total, x, y, size}
			wg.Done()
		}(size)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	best := 0
	bestX := 0
	bestY := 0
	bestSize := 0

	for res := range results {

		if res.total > best {

			best = res.total
			bestX = res.x
			bestY = res.y
			bestSize = res.size

		}

	}

	return bestX, bestY, bestSize

}

func main() {

	serial := 6548

	g := grid(serial)

	_, x, y := search(g, 3)

	fmt.Printf("The X,Y coordinate of the top-left fuel cell of the 3x3 square with the largest total power in the grid with serial %v: %v,%v\n",
		serial, x+1, y+1)

	x, y, size := sizes(g)

	fmt.Printf("The X,Y,size identifier of the square with the largest total power for serial %v: %v,%v,%v\n",
		serial, x+1, y+1, size)

}
