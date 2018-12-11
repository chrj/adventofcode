package main

import "testing"

func TestGrid(t *testing.T) {

	tests := []struct {
		serial             int
		x, y               int
		expectedPowerLevel int
	}{
		{57, 122, 79, -5},
		{39, 217, 196, 0},
		{71, 101, 153, 4},
	}

	for _, test := range tests {

		g := grid(test.serial)
		if g[test.x-1][test.y-1] != test.expectedPowerLevel {
			t.Errorf("expected fuel cell at %v,%v, grid serial number %v, to have power level %v. Got: %v.",
				test.x, test.y, test.serial, test.expectedPowerLevel, g[test.x-1][test.y-1])

		}

	}

}

func TestSearch(t *testing.T) {

	tests := []struct {
		serial       int
		bestX, bestY int
	}{
		{18, 33, 45},
		{42, 21, 61},
	}

	for _, test := range tests {

		_, x, y := search(grid(test.serial), 3)

		if x+1 != test.bestX || y+1 != test.bestY {
			t.Errorf("for serial %v; expected best 3x3 square at %v,%v. Got: %v,%v",
				test.serial, test.bestX, test.bestY, x+1, y+1)

		}

	}

}

func TestSizes(t *testing.T) {

	tests := []struct {
		serial                 int
		bestX, bestY, bestSize int
	}{
		{18, 90, 269, 16},
		//		{42, 232, 251, 12},
	}

	for _, test := range tests {

		x, y, size := sizes(grid(test.serial))

		if x+1 != test.bestX || y+1 != test.bestY || size != test.bestSize {
			t.Errorf("for serial %v; expected size %v square at %v,%v. Got: %v,%v,%v",
				test.serial, test.bestSize, test.bestX, test.bestY, x+1, y+1, size)

		}

	}

}
