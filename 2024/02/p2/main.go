package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	inputs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	saveReports := 0
	for _, report := range strings.Split(string(inputs), "\n") {
		if len(strings.Trim(report, " ")) == 0 {
			continue
		}

		lStr := strings.Split(report, " ")

		levels := make([]int, len(lStr))
		t := 0
		for i, s := range lStr {
			c, _ := strconv.Atoi(s)
			levels[i] = c

			if i == 0 {
				continue
			}

			t += sign(c - levels[i-1])
		}

		if t == 0 {
			continue
		}

		dir := sign(t)
		p := levels[0]

		m := 0

		for i := 1; i < len(levels); i++ {
			c := levels[i]
			d := c - p
			s := sign(d)

			if d == 0 || s != dir || math.Abs(float64(d)) > 3 {
				m++
				continue
			}

			p = c
		}

		if m < 2 {
			saveReports++
		}
	}

	println(saveReports)
}

func sign(i int) int {
	return i / max(int(math.Abs(float64(i))), 1)
}
