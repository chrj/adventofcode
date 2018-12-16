package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sim(target string) (string, int) {

	targetLength, _ := strconv.Atoi(target)

	var recipes strings.Builder

	recipes.Grow(targetLength)
	recipes.WriteString("37")

	elf1 := 0
	elf2 := 1

	var last7 string

	for recipes.Len() < (targetLength+10) || strings.Index(last7, target) == -1 {

		// Compute new recipes

		recipes.WriteString(fmt.Sprintf("%v", int(recipes.String()[elf1]-48)+int(recipes.String()[elf2]-48)))

		// Move elfs

		elf1 = (elf1 + 1 + int(recipes.String()[elf1]-48)) % recipes.Len()
		elf2 = (elf2 + 1 + int(recipes.String()[elf2]-48)) % recipes.Len()

		if recipes.Len() >= 7 {
			last7 = recipes.String()[recipes.Len()-7:]
		}

	}

	return recipes.String()[targetLength : targetLength+10], strings.Index(recipes.String(), target)

}

func main() {
	fmt.Println(sim("540561"))
}
