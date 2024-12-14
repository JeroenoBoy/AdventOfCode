package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	fileInput  = flag.String("i", "input.txt", "change the input file")
	iterations = flag.Int("c", 75, "the count of iterations")
)

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.ReplaceAll(string(input), "\n", ""), " ")

	stones := make(map[string]int)
	for _, input := range s {
		if input == "\n" {
			continue
		}
		stones[input] += 1
	}

	for i := 0; i < *iterations; i++ {
		newStones := make(map[string]int)
		for key, value := range stones {
			if key == "0" {
				newStones["1"] += value
				continue
			}

			rc := utf8.RuneCountInString(key)
			if rc%2 == 0 {
				a, _ := strconv.Atoi(key[:rc/2])
				b, _ := strconv.Atoi(key[rc/2:])

				aS := strconv.Itoa(a)
				bS := strconv.Itoa(b)

				newStones[aS] += value
				newStones[bS] += value
				continue
			}

			num, _ := strconv.Atoi(key)
			newStones[strconv.Itoa(num*2024)] += value
		}

		stones = newStones
		count := 0
		for _, val := range stones {
			count += val
		}
	}

	count := 0
	for _, val := range stones {
		count += val
	}
	println(count)
}
