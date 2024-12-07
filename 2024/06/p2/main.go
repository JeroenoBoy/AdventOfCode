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

	var startPos vector
	var startDir direction

	for y, l := range split {
		for x, c := range l {
			switch c {
			case '^':
				startPos = vector{X: x, Y: y}
				startDir = direction{X: 0, Y: -1}
				break
			case '#':
				worldMap[y*xL+x] = true
				break
			}
		}
	}

	guardPos := startPos
	guardDir := startDir

	traveledMap := make([]bool, len(worldMap))
	traveledMap[guardPos.Y*xL+guardPos.X] = true
	traveledPoints := make([]vector, 0)
	for true {
		nextPos := guardPos.add(vector(guardDir))

		if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= xL || nextPos.Y >= yL {
			break
		}

		if worldMap[nextPos.Y*xL+nextPos.X] {
			guardDir = guardDir.rotate()
			continue
		}

		guardPos = nextPos
		if !traveledMap[guardPos.Y*xL+guardPos.X] {
			traveledMap[guardPos.Y*xL+guardPos.X] = true
			traveledPoints = append(traveledPoints, guardPos)
		}
	}

	result := 0
	for _, extraPoint := range traveledPoints {

		guardPos := startPos
		guardDir := startDir

		worldMap[extraPoint.Y*xL+extraPoint.X] = true

		for i := 0; i < 10_000_000; i++ { // I know of a more efficient way but this is just easier rn
			nextPos := guardPos.add(vector(guardDir))

			if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= xL || nextPos.Y >= yL {
				goto next
			}

			if worldMap[nextPos.Y*xL+nextPos.X] {
				guardDir = guardDir.rotate()
				continue
			}

			guardPos = nextPos
		}

		result++
	next:
		worldMap[extraPoint.Y*xL+extraPoint.X] = false
	}

	println(result)
}

type vector struct {
	X int
	Y int
}

type direction vector

func (n vector) add(other vector) vector {
	n.X += other.X
	n.Y += other.Y
	return n
}

func (n direction) rotate() direction { // Map is in reverse so we rotate anti-clockwise
	x := n.X
	n.X = -n.Y
	n.Y = x
	return n
}
