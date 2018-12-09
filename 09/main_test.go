package main

import "testing"

func TestSim(t *testing.T) {

	tests := []struct {
		players, lastMarble, highscore int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for _, test := range tests {
		h := sim(test.players, test.lastMarble)
		if h != test.highscore {
			t.Errorf("with players: %v, last marble worth: %v. Got highscore: %v, expected: %v", test.players, test.lastMarble, h, test.highscore)
		}
	}

}
