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
	for i, report := range strings.Split(string(inputs), "\n") {
		if len(strings.Trim(report, " ")) == 0 {
			continue
		}

		r := strings.Split(report, " ")
		allReports := make([][]int, len(r)+1)
		allReports[0] = make([]int, len(r))

		for i, v := range r {
			allReports[0][i], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}

		for i := 1; i < len(allReports); i++ {
			allReports[i] = make([]int, len(allReports[0])-1)
			toSkip := i - 1
			for j := 0; j < len(allReports[i]); j++ {
				d := j
				if j >= toSkip {
					d++
				}

				allReports[i][j] = allReports[0][d]
			}
		}

		for _, levels := range allReports {
			if i == 54 {
				s := ""
				for _, n := range levels {
					s += strconv.Itoa(n) + " "
				}
			}

			a := levels[0]
			b := levels[1]

			v := b
			dir := sign(b - a)

			if dir == 0 || math.Abs(float64(b-a)) > 3 {
				continue
			}

			for i := 2; i < len(levels); i++ {
				c := levels[i]
				d := c - v
				s := sign(d)
				if s != dir || math.Abs(float64(d)) > 3 {
					goto invalid
				}
				v = c
			}
			goto valid
		invalid:
			continue
		}
		continue
	valid:
		saveReports++
	}

	println(saveReports)
}

func sign(i int) int {
	return i / max(int(math.Abs(float64(i))), 1)
}
