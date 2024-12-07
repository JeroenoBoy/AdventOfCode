package main

import (
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(input), "\n")

	yL := len(split) - 1
	xL := utf8.RuneCountInString(split[0])

	worldMap := make([]bool, yL*xL)
	traveledPoints := make([]bool, len(split)*xL)

	var guardPos vector
	var guardDir direction

	for y, l := range split {
		for x, c := range l {
			switch c {
			case '^':
				guardPos = vector{x: x, y: y}
				guardDir = direction{x: 0, y: -1}
				break
			case '#':
				worldMap[y*xL+x] = true
				break
			}
		}
	}

	traveledPoints[guardPos.y*xL+guardPos.x] = true
	uniquePoints := 1
	for true {
		nextPos := guardPos.add(vector(guardDir))

		if nextPos.x < 0 || nextPos.y < 0 || nextPos.x >= xL || nextPos.y >= yL {
			break
		}

		if worldMap[nextPos.y*xL+nextPos.x] {
			guardDir = guardDir.rotate()
			continue
		}

		guardPos = nextPos
		if true && !traveledPoints[nextPos.y*xL+nextPos.x] {
			traveledPoints[nextPos.y*xL+nextPos.x] = true
			uniquePoints++
		}
	}

	println(uniquePoints)
}

type vector struct {
	x int
	y int
}

type direction vector

func (n vector) add(other vector) vector {
	n.x += other.x
	n.y += other.y
	return n
}

func (n direction) rotate() direction { // Map is in reverse so we rotate anti-clockwise
	x := n.x
	n.x = -n.y
	n.y = x
	return n
}
