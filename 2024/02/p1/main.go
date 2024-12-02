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

		levels := strings.Split(report, " ")

		a, _ := strconv.Atoi(levels[0])
		b, _ := strconv.Atoi(levels[1])

		v := b
		dir := sign(b - a)

		if dir == 0 {
			continue
		}

		if math.Abs(float64(b-a)) > 3 {
            continue
		}

		saveReports++
		for i := 2; i < len(levels); i++ {
			c, _ := strconv.Atoi(levels[i])
			d := c - v
			if math.Abs(float64(d)) > 3 {
				saveReports--
				break
			}
			s := sign(d)
			if s == 0 {
				saveReports--
				break
			}
			if s == dir {
				v = c
				continue
			}

			saveReports--
			break
		}
	}

	println(saveReports)
}

func sign(i int) int {
	return i / max(int(math.Abs(float64(i))), 1)
}
