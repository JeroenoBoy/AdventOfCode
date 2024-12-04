package main

import (
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

var runeList = "XMAS"
var runeCount = utf8.RuneCountInString(runeList)

func main() {

	inputs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputs), "\n")

	matches := 0
	for i, l := range lines {
		for j, c := range l {
			if c != 'X' {
				continue
			}

			println(i, j)

			compareDiagonal(lines, 0, 1, i, j, &matches)
			compareDiagonal(lines, 0, -1, i, j, &matches)
			compareDiagonal(lines, 1, 0, i, j, &matches)
			compareDiagonal(lines, -1, 0, i, j, &matches)

			compareDiagonal(lines, 1, 1, i, j, &matches)
			compareDiagonal(lines, -1, 1, i, j, &matches)
			compareDiagonal(lines, 1, -1, i, j, &matches)
			compareDiagonal(lines, -1, -1, i, j, &matches)
		}
	}

	println(matches)
}

func compareDiagonal(lines []string, dirV, dirH, row, col int, matchCount *int) {
	rowLen := len(lines)
	colLen := utf8.RuneCountInString(lines[row])

	if dirV == -1 && row < runeCount-1 {
		return
	}
	if dirV == 1 && row >= rowLen-runeCount {
		return
	}
	if dirH == -1 && col < runeCount-1 {
		return
	}
	if dirH == 1 && col >= colLen-runeCount+1 { // I really don't know why I need a +1 here and not on the line ^6 but it works
		return
	}

	for i := 1; i < runeCount; i++ {
		if lines[row+i*dirV][col+i*dirH] != runeList[i] {
			return
		}
	}

	*matchCount++
}
