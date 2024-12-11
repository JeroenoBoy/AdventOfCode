package main

import (
	"flag"
	"io/ioutil"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	fileInput  = flag.String("i", "input.txt", "change the input file")
	iterations = flag.Int("c", 25, "the count of iterations")
)

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.ReplaceAll(string(input), "\n", ""), " ")
	stones := make([]string, 0, len(s))
	for _, input := range s {
		if input == "\n" {
			continue
		}

		stones = append(stones, input)
	}

	for i := 0; i < *iterations; i++ {
		for x := 0; x < len(stones); x++ {
			doStep(&stones, &x)
		}
	}

	println(len(stones))
}

func doStep(stones *[]string, index *int) {
	str := (*stones)[*index]
	num, _ := strconv.Atoi(str)

	if num == 0 {
		(*stones)[*index] = "1"
		return
	}

	rc := utf8.RuneCountInString(str)
	if rc%2 == 0 && rc >= 2 {
		c := utf8.RuneCountInString(str) / 2
		a, _ := strconv.Atoi(str[:c])
		b, _ := strconv.Atoi(str[c:])

		aS := strconv.Itoa(a)
		bS := strconv.Itoa(b)

		s := slices.Insert((*stones), *index+1, bS)
		s[*index] = aS
		*stones = s

		*index++

		return
	}

	(*stones)[*index] = strconv.Itoa(num * 2024)
}
