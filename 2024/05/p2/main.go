package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rules := make([][]int, 99)
	pages := make([][]int, 0)

	parseType := 0
	for _, input := range strings.Split(string(inputs), "\n") {
		input = strings.Trim(input, " ")
		if input == "" {
			parseType = 1
			continue
		}

		if parseType == 0 {
			s := strings.Split(input, "|")
			a, _ := strconv.Atoi(s[0])
			b, _ := strconv.Atoi(s[1])

			if rules[a] == nil {
				rules[a] = make([]int, 99)
			}

			rules[a][b] = 1
			continue
		}

		if parseType == 1 {
			s := strings.Split(input, ",")
			a := make([]int, len(s))
			pages = append(pages, a)

			for i := range a {
				a[i], _ = strconv.Atoi(s[i])
			}
		}
	}

	total := 0
	for pI := 0; pI < len(pages); pI++ {
		correctlyOrdered := true
	reset:
		page := pages[pI]
		i := len(page) - 1

		for ; i >= 1; i-- {
			rule := rules[page[i]]
			for y := i - 1; y >= 0; y-- {
				if rule[page[y]] == 1 {
					shiftEnd(&pages[pI], y)
					correctlyOrdered = false
					goto reset
				}
			}
		}

		if !correctlyOrdered {
            total += page[(len(page)+1)/2-1]
		}
	}

	println(total)
}

func shiftEnd(s *[]int, x int) {
	tmp := (*s)[x]
	*s = append((*s)[:x], (*s)[x+1:]...)
	*s = append((*s), tmp)
}
