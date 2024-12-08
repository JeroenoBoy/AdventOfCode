package main

import (
	"flag"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

var (
	fileInput = flag.String("i", "input.txt", "change the input file")
)

func main() {
	flag.Parse()

	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(input), "\n")

	antennaMap := make(map[rune][]Vector)
	antinodesMap := newArray2D[bool](utf8.RuneCountInString(s[0]), len(s)-1)
	antiNodesCount := 0

	for y, row := range s {
		for x, rune := range row {
			if rune == '.' {
				continue
			}

			p, ok := antennaMap[rune]
			if !ok {
				p = make([]Vector, 0)
			}

			p = append(p, Vector{X: x, Y: y})
			antennaMap[rune] = p
		}
	}

	for _, antennas := range antennaMap {
		for aI, a := range antennas {
			for bI, b := range antennas {
				if aI == bI {
					continue
				}

				for i := 0; true; i++ {
					np := a.Plus(a.Minus(b).Multiply(i))
					if !antinodesMap.contains(np.X, np.Y) {
						break
					}
					if antinodesMap.get(np.X, np.Y) {
						continue
					}

					antinodesMap.set(np.X, np.Y, true)
					antiNodesCount++
				}
			}
		}
	}

	println(antiNodesCount)
}

type Array2D[T any] struct {
	array  []T
	Width  int
	Height int
}

func newArray2D[T any](width, height int) Array2D[T] {
	return Array2D[T]{
		array:  make([]T, width*height),
		Width:  width,
		Height: height,
	}
}

func (arr *Array2D[T]) contains(x, y int) bool {
	return x >= 0 && x < arr.Width && y >= 0 && y < arr.Height
}

func (arr *Array2D[T]) get(x, y int) T {
	return arr.array[y*arr.Width+x]
}

func (arr *Array2D[T]) set(x, y int, nv T) {
	arr.array[y*arr.Width+x] = nv
}

type Vector struct {
	X int
	Y int
}

func (n Vector) Plus(other Vector) Vector {
	n.X += other.X
	n.Y += other.Y
	return n
}

func (n Vector) Minus(other Vector) Vector {
	n.X -= other.X
	n.Y -= other.Y
	return n
}

func (n Vector) Multiply(scalar int) Vector {
	n.X *= scalar
	n.Y *= scalar
	return n
}
