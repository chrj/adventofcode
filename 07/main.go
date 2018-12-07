package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type depgraph struct {
	m        map[string][]string
	resolved map[string]struct{}
}

func newdepgraph() *depgraph {
	return &depgraph{
		m:        make(map[string][]string),
		resolved: make(map[string]struct{}),
	}
}

func (d *depgraph) add(from, to string) {
	d.m[to] = append(d.m[to], from)
	d.m[from] = d.m[from]
}

func (d *depgraph) resolve(dep string) {
	d.resolved[dep] = struct{}{}
}

func (d *depgraph) list() []string {
	result := []string{}
	for k := range d.m {
		result = append(result, k)
	}
	return result
}

func (d *depgraph) reset() {
	d.resolved = make(map[string]struct{})
}

func (d *depgraph) unresolved(dep string) int {
	unresolved := 0
	for _, ddep := range d.m[dep] {
		if _, found := d.resolved[ddep]; found {
			continue
		}
		unresolved++
	}
	return unresolved
}

func main() {

	deps := newdepgraph()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		a := string(line[5])
		b := string(line[36])

		deps.add(a, b)

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	steps := deps.list()
	order := []string{}

	for len(steps) > 0 {

		sort.Slice(steps, func(i, j int) bool {

			di := deps.unresolved(steps[i])
			dj := deps.unresolved(steps[j])

			if di == dj {
				return steps[i] < steps[j]
			}

			return di < dj

		})

		next := steps[0]

		order = append(order, next)
		steps = steps[1:]

		deps.resolve(next)

	}

	fmt.Printf("instruction order: %s\n", strings.Join(order, ""))

	// Part 2

	deps.reset()

	workers := [5]int{}      // worker-id -> seconds left of work
	workingOn := [5]string{} // worker-id -> what are they working on

	second := 0

	for {

		// Check if any work has finished

		for wid, assignment := range workingOn {

			if assignment != "" && workers[wid] == 0 {
				deps.resolve(assignment)
				workingOn[wid] = ""
			}

		}

		// Check for idle workers and assign jobs

		for wid, workleft := range workers {

			if workleft > 0 {
				continue
			}

			// try to find job

			for i := 0; i < len(order); i++ {

				// Check if theres any unresolved dependencies

				if deps.unresolved(order[i]) == 0 {

					// Start work and remove from order queue

					workers[wid] = int(order[i][0]) - 64 + 60
					workingOn[wid] = order[i]

					order = append(order[:i], order[i+1:]...)

					break

				}

			}

		}

		// Simulate a second

		allWorkersIdle := true

		for wid, workleft := range workers {

			if workleft > 0 {
				workers[wid]--
			}

			if workleft > 1 {
				allWorkersIdle = false
			}

		}

		second++

		if len(order) == 0 && allWorkersIdle {
			break
		}

	}

	fmt.Printf("all work done in %v seconds\n", second)

}
