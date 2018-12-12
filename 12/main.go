package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type plants struct {
	pots       []bool
	reactions  []reaction
	firstIndex int
}

type reaction struct {
	data  [5]bool
	plant bool
}

func (r reaction) String() string {

	s := ""

	for _, r := range r.data {
		if r {
			s += "#"
		} else {
			s += "."
		}
	}

	s += " => "

	if r.plant {
		s += "#"
	} else {
		s += "."
	}

	return s

}

func NewPlants() *plants {
	return &plants{}
}

func (p *plants) SetState(state string) {

	state = strings.TrimSpace(state)

	if p.pots == nil || len(p.pots) > 0 {
		p.pots = []bool{false, false, false}
		p.firstIndex = -3
	}

	for _, r := range state {
		p.pots = append(p.pots, r == '#')
	}

	p.pots = append(p.pots, false, false, false)

}

func (p *plants) SetReactions(reactions []string) {

	p.reactions = []reaction{}

	for _, r := range reactions {

		r := strings.TrimSpace(r)

		if len(r) == 0 {
			continue
		}

		parts := strings.Split(r, " => ")

		p.reactions = append(p.reactions, reaction{
			data: [5]bool{
				parts[0][0] == '#',
				parts[0][1] == '#',
				parts[0][2] == '#',
				parts[0][3] == '#',
				parts[0][4] == '#',
			},
			plant: parts[1] == "#",
		})

	}

}

func (p *plants) String() string {

	s := ""

	for _, pot := range p.pots {
		if pot {
			s += "#"
		} else {
			s += "."
		}
	}

	return s
}

func (p *plants) React() {

	newpots := make([]bool, len(p.pots))

	for i := 2; i < len(p.pots)-2; i++ {

		reacted := false

		for _, r := range p.reactions {

			if p.pots[i-2] == r.data[0] &&
				p.pots[i-1] == r.data[1] &&
				p.pots[i] == r.data[2] &&
				p.pots[i+1] == r.data[3] &&
				p.pots[i+2] == r.data[4] {

				reacted = r.plant
				break

			}

		}

		newpots[i] = reacted
	}

	p.pots = newpots

	// Pad with extra pots if needed to the right
	if p.pots[len(p.pots)-3] {
		p.pots = append(p.pots, false)
	}

	// Pad with extra pots if needed to the left
	if p.pots[2] {
		p.pots = append([]bool{false}, p.pots...)
		p.firstIndex--
	}

}

func (p *plants) Sum() int {
	sum := 0
	for i, plant := range p.pots {
		if plant {
			sum += p.firstIndex + i
		}
	}
	return sum
}

func main() {

	var p *plants

	reactions := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		if p == nil {

			// First line

			p = NewPlants()
			p.SetState(line[15:])
			continue

		}

		reactions = append(reactions, strings.TrimSpace(line))

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	p.SetReactions(reactions)

	for i := 0; i < 20; i++ {
		p.React()
	}

	fmt.Printf("After 20 generations, the sum of the numbers of all pots which contain a plant: %v\n", p.Sum())

	n := 50000000000

	// Allow the sum to converge on a linear sum

	for i := 20; i < 1000; i++ {
		p.React()
	}

	// Measure the sum growth over 1000 generations

	sum := p.Sum()

	for i := 1000; i < 2000; i++ {
		p.React()
	}

	step := p.Sum() - sum

	fmt.Printf("After %v generations, the sum of the numbers of all pots which contain a plant: %v\n", n, ((n-2000)/1000*step)+p.Sum())

}
