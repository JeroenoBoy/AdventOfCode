package main

import (
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

func main() {

	inputs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputs), "\n")

	matches := 0
	for y := 1; y < len(lines)-2; y++ { // Again I don't know why I need 1 more than the other but it works so idc
		for x := 1; x < utf8.RuneCountInString(lines[y])-1; x++ {
			if lines[y][x] != 'A' {
				continue
			}

			if !compareDiagonal(lines, 1, 1, y, x) {
				continue
			}
			if !compareDiagonal(lines, 1, -1, y, x) {
				continue
			}

			matches++
		}
	}

	println(matches)
}

func compareDiagonal(lines []string, dirV, dirH, y, x int) bool {
	a := lines[y+dirV][x+dirH]
	b := lines[y-dirV][x-dirH]

	if a == 'M' && b == 'S' {
		return true
	}
	if b == 'M' && a == 'S' {
		return true
	}

	return false
}
