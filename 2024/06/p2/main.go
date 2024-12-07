package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
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
				guardPos = vector{X: x, Y: y}
				guardDir = direction{X: 0, Y: -1}
				break
			case '#':
				worldMap[y*xL+x] = true
				break
			}
		}
	}

	traveledPoints[guardPos.Y*xL+guardPos.X] = true
	lastObstacles := make([]vector, 0, 4)
	uniquePoints := 0
	for true {
		nextPos := guardPos.add(vector(guardDir))

		if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= xL || nextPos.Y >= yL {
			break
		}

		if worldMap[nextPos.Y*xL+nextPos.X] {
			for len(lastObstacles) >= 3 {
				lastObstacles = lastObstacles[len(lastObstacles)-2:]
			}
			lastObstacles = append(lastObstacles, guardPos)
			if isLoopable(lastObstacles, worldMap, xL, guardDir) {
				uniquePoints++
			}
			guardDir = guardDir.rotate()
			continue
		}

		guardPos = nextPos
	}

	println(uniquePoints)
}

type vector struct {
	X int
	Y int
}

func (v *vector) String() string {
	return fmt.Sprintf("{%v,%v}", v.X, v.Y)
}

type direction vector

func (n vector) add(other vector) vector {
	n.X += other.X
	n.Y += other.Y
	return n
}

func (n vector) minus(other vector) vector {
	n.X -= other.X
	n.Y -= other.Y
	return n
}

func (n vector) multiply(magnitude int) vector {
	n.X *= magnitude
	n.Y *= magnitude
	return n
}

func (n vector) dot(other direction) int {
	return n.X*other.X + n.Y*other.Y
}

func (n vector) toDirection() direction {

	if n.X != 0 && n.Y != 0 {
		panic("One of X or Y must be 0, found: " + strconv.Itoa(n.X) + "," + strconv.Itoa(n.Y))
	}

	if n.X != 0 {
		n.X /= int(math.Abs(float64(n.X)))
		return direction(n)
	}

	if n.Y != 0 {
		n.Y /= int(math.Abs(float64(n.Y)))
		return direction(n)
	}

	panic("One of X or Y must not be 0, found: " + strconv.Itoa(n.X) + "," + strconv.Itoa(n.Y))
}

func (n direction) rotate() direction { // Map is in reverse so we rotate anti-clockwise
	x := n.X
	n.X = -n.Y
	n.Y = x
	return n
}

func (n direction) rotateReverse() direction {
	x := n.X
	n.X = n.Y
	n.Y = -x
	return n
}

func isLoopable(lastObstacles []vector, obstacles []bool, xL int, dir direction) bool {
	size := len(lastObstacles)
	if size != 3 {
		return false
	}

	lastPos := lastObstacles[0]
	midPos := lastObstacles[1]
	firstPos := lastObstacles[2]

	var targetPos vector
	targetPos = targetPos.add(vector(dir.rotateReverse()).multiply(lastPos.dot(dir.rotateReverse())))
	targetPos = targetPos.add(vector(dir).multiply(firstPos.dot(dir)))

    targetObtacle := targetPos.add(vector(dir.rotate()))
    if obstacles[targetObtacle.Y*xL+targetObtacle.X] {
        return false
    }

	fmt.Println("dir=", dir, "tp=", targetPos, "fp=", firstPos, "mp=", midPos, "lp=", lastPos)

	difA := lastPos.minus(targetPos)
	difB := firstPos.minus(targetPos)
	dirA := difA.toDirection()
	dirB := difB.toDirection()

	for a := 0; a < int(math.Abs(float64(difA.dot(dirA)))); a++ {
		pos := targetPos.add(vector(dirA).multiply(a))
		if obstacles[pos.Y*xL+pos.X] {
			fmt.Println("ha=", pos, a, math.Abs(float64(difA.dot(dirA))), "\n")
			return false
		}
	}

	for b := 0; b < int(math.Abs(float64(difB.dot(dirB)))); b++ {
		pos := targetPos.add(vector(dirB).multiply(b))
		if obstacles[pos.Y*xL+pos.X] {
			fmt.Println("hb=", pos, b, math.Abs(float64(difB.dot(dirB))), "\n")
			return false
		}
	}

	return true
}
