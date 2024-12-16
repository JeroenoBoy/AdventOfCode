package main

import (
	"flag"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	fileInput      = flag.String("i", "input.txt", "change the input file")
	iterationCount = flag.Int("c", 100, "iterations of the robots")
	gridSize       = flag.String("s", "101,103", "iterations of the robots")

	robotRegex = regexp.MustCompile("p=(\\d*),(\\d*) v=(-{0,1}\\d*),(-{0,1}\\d*)")
)

func main() {
	flag.Parse()

	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	gridSize := getGridSize()
	halfGrid := gridSize.Remove(Vector{X: 1, Y: 1}).Divide(2)

	s := strings.Split(strings.Trim(string(input), "\n"), "\n")

	grid1Sum := 0
	grid2Sum := 0
	grid3Sum := 0
	grid4Sum := 0

	for _, robotData := range s {
		rr := robotRegex.FindStringSubmatch(robotData)

		pX, err := strconv.Atoi(rr[1])
		panicIfErr(err)
		pY, err := strconv.Atoi(rr[2])
		panicIfErr(err)
		vX, err := strconv.Atoi(rr[3])
		panicIfErr(err)
		vY, err := strconv.Atoi(rr[4])
		panicIfErr(err)

		robot := Robot{
			Position: Vector{X: pX, Y: pY},
			Velocity: Vector{X: vX, Y: vY},
		}

		robot.EndPosition = robot.Position.Add(robot.Velocity.Multiply(*iterationCount)).Modulo(gridSize)

		x := robot.EndPosition.X
		y := robot.EndPosition.Y

		if x < halfGrid.X && y < halfGrid.Y {
			grid1Sum++
		} else if x < halfGrid.X && y > halfGrid.Y {
			grid2Sum++
		} else if x > halfGrid.X && y < halfGrid.Y {
			grid3Sum++
		} else if x > halfGrid.X && y > halfGrid.Y {
			grid4Sum++
		}
	}

	println(grid1Sum * grid2Sum * grid3Sum * grid4Sum)
}

type Robot struct {
	Position    Vector
	Velocity    Vector
	EndPosition Vector
}

type Vector struct {
	X int
	Y int
}

func (n Vector) Add(other Vector) Vector {
	n.X += other.X
	n.Y += other.Y
	return n
}

func (n Vector) Remove(other Vector) Vector {
	n.X -= other.X
	n.Y -= other.Y
	return n
}

func (n Vector) Multiply(multi int) Vector {
	n.X *= multi
	n.Y *= multi
	return n
}

func (n Vector) Divide(div int) Vector {
	n.X /= div
	n.Y /= div
	return n
}

func (n Vector) Modulo(other Vector) Vector {
	n.X %= other.X
	n.Y %= other.Y
	if n.X < 0 {
		n.X = other.X + n.X
	}
	if n.Y < 0 {
		n.Y = other.Y + n.Y
	}
	return n
}

func getGridSize() Vector {
	sizeStr := strings.Split(*gridSize, ",")
	x, err := strconv.Atoi(sizeStr[0])
	panicIfErr(err)
	y, err := strconv.Atoi(sizeStr[1])
	panicIfErr(err)
	return Vector{
		X: x,
		Y: y,
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
