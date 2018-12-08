package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(v []int) (int, int, int) {

	if len(v) == 0 {
		return 0, 0, 0
	}

	cn := v[0] // child node count
	md := v[1] // child node count
	sum := 0

	cnvalues := []int{}
	consumed := 0

	i := 2 // skip header

	for consumed < cn {

		csum, cvalue, offset := parse(v[i:])

		sum += csum
		i += offset

		consumed++

		cnvalues = append(cnvalues, cvalue)

	}

	metadata := v[i : i+md]
	mdsum := 0
	for _, v := range metadata {
		mdsum += v
	}

	// Calculate root value

	rv := 0
	if cn == 0 {

		rv = mdsum

	} else {

		for _, md := range metadata {

			md-- // 1-indexed

			if md < len(cnvalues) {
				rv += cnvalues[md]
			}

		}
	}

	return sum + mdsum, rv, i + md

}

func main() {

	licenseData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	license := []int{}
	for _, v := range strings.Fields(string(bytes.TrimSpace(licenseData))) {

		x, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		license = append(license, int(x))

	}

	sum, rv, _ := parse(license)

	fmt.Printf("The sum of all metadata entries: %v\n", sum)
	fmt.Printf("The value of the root node is: %v\n", rv)

}
