package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type shift struct {
	guard  int
	asleep []sleepcycle
}

func (s shift) sleepMinutes() int {

	minutes := 0

	for _, cycle := range s.asleep {
		minutes += cycle.end - cycle.start
	}

	return minutes

}

type sleepcycle struct {
	start, end int
}

var dateFmt = "2006-01-02 15:04"

func main() {

	shifts := map[time.Time]shift{}
	lines := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	sort.Strings(lines)

	for _, line := range lines {

		timestamp := line[1:17]
		action := line[19:]

		t, err := time.Parse(dateFmt, timestamp)
		if err != nil {
			log.Fatal(err)
		}

		d := t.Add(time.Hour).Truncate(time.Hour * 24)
		s := shifts[d]
		f := strings.Fields(action)

		switch f[0] {

		case "Guard":

			id, err := strconv.ParseInt(f[1][1:], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			s.guard = int(id)
			s.asleep = []sleepcycle{}

		case "falls":

			s.asleep = append(s.asleep, sleepcycle{start: t.Minute()})

		case "wakes":

			s.asleep[len(s.asleep)-1].end = t.Minute()

		}

		shifts[d] = s

	}

	// Part 1

	guardSleep := map[int]int{} // Total sleep per guard

	for _, shift := range shifts {
		guardSleep[shift.guard] += shift.sleepMinutes()
	}

	// Find most sleepy guard

	sleepyGuard := 0
	sleepyGuardMinutes := 0

	for guard, sleep := range guardSleep {
		if sleep > sleepyGuardMinutes {
			sleepyGuard = guard
			sleepyGuardMinutes = sleep
		}
	}

	// Find frequency per minute for the sleepy guard
	// while keeping track of the minute with the highest
	// frequency

	minutes := [60]int{}
	mostLikelyMinute := 0
	mostLikelyMinuteFreq := 0

	for _, shift := range shifts {

		if shift.guard != sleepyGuard {
			continue
		}

		for _, cycle := range shift.asleep {

			for i := cycle.start; i < cycle.end; i++ {

				minutes[i]++

				if minutes[i] > mostLikelyMinuteFreq {
					mostLikelyMinute = i
					mostLikelyMinuteFreq = minutes[i]
				}

			}

		}
	}

	fmt.Printf("Most sleepy guard: %v (%v minutes). Most likely asleep on minute %v: %v\n", sleepyGuard, sleepyGuardMinutes, mostLikelyMinute, sleepyGuard*mostLikelyMinute)

	// Part 2

	freq := make(map[int][60]int) // guard -> minute -> freq

	maxFreq := 0
	maxFreqGuard := 0
	maxFreqMinute := 0

	for _, shift := range shifts {

		f := freq[shift.guard]

		for _, cycle := range shift.asleep {

			for i := cycle.start; i < cycle.end; i++ {

				f[i]++

				if f[i] > maxFreq {
					maxFreq = f[i]
					maxFreqGuard = shift.guard
					maxFreqMinute = i
				}

			}
		}

		freq[shift.guard] = f

	}

	fmt.Printf("Most frequently sleeping guard: %v (minute %v, %v times): %v\n", maxFreqGuard, maxFreqMinute, maxFreq, maxFreqGuard*maxFreqMinute)

}
