package main

import (
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(input), "\n")

	match := regexp.MustCompile("^(\\d+) +(\\d+)$")

	aArray := make([]int, len(inputs))
	bArray := make([]int, len(inputs))

	for _, line := range inputs {
		if len(line) == 0 {
			continue
		}

		matches := match.FindStringSubmatch(line)

		a, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		aArray = append(aArray, a)
		bArray = append(bArray, b)
	}

	sort.Ints(aArray)
	sort.Ints(bArray)

	result := 0

	for i := range aArray {
		a := aArray[i]
		b := bArray[i]

		result = result + int(math.Abs(float64(a-b)))
	}

	log.Println(result)
}
