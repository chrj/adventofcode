package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func eq(s1, s2 string) string {

	result := ""

	for i, c := range s1 {

		if i < len(s2) && byte(c) == s2[i] {
			result += string(c)
		}
	}

	return result

}

func main() {

	ids := []string{}
	eqc := ""

	scanner := bufio.NewScanner(os.Stdin)
outerloop:
	for scanner.Scan() {

		id := scanner.Text()
		ids = append(ids, id)

		for _, s := range ids[:len(ids)-1] {

			eqc = eq(id, s)

			if len(eqc) == len(id)-1 {
				break outerloop
			}

		}

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("equal characters: %v\n", eqc)

}
