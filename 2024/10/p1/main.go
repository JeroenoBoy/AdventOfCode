package main

import (
	"flag"
	"io/ioutil"
	"strconv"
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
	worldMap := newArray2D[int](utf8.RuneCountInString(s[0]), len(s)-1)
	pairs := make(map[Vector]map[Vector]bool)
	ninePositions := make([]Vector, 0)

	for y := 0; y < worldMap.Height; y++ {
		for x := 0; x < worldMap.Width; x++ {
			rune := s[y][x]

			if rune == '.' {
				worldMap.set(x, y, -1)
			} else {
				i, _ := strconv.Atoi(string(rune))
				worldMap.set(x, y, i)

				if i == 9 {
					ninePositions = append(ninePositions, Vector{X: x, Y: y})
				}
			}
		}
	}

	sum := 0
	for _, pos := range ninePositions {
		sum += getCheckSum(worldMap, pairs, pos, pos, 9)
	}

	println(sum)
}

func getCheckSum(worldMap Array2D[int], pairs map[Vector]map[Vector]bool, startPosition, position Vector, next int) int {

	if !worldMap.contains(position.X, position.Y) {
		return 0
	}

	if worldMap.get(position.X, position.Y) != next {
		return 0
	}

	if next == 0 {
		m, ok := pairs[startPosition]
		if !ok {
			m = make(map[Vector]bool)
			pairs[startPosition] = m
		}

		if m[position] {
			return 0
		}

		m[position] = true

		return 1
	}

	sum := 0

	sum += getCheckSum(worldMap, pairs, startPosition, position.Plus(Vector{X: 0, Y: 1}), next-1)
	sum += getCheckSum(worldMap, pairs, startPosition, position.Plus(Vector{X: 0, Y: -1}), next-1)
	sum += getCheckSum(worldMap, pairs, startPosition, position.Plus(Vector{X: 1, Y: 0}), next-1)
	sum += getCheckSum(worldMap, pairs, startPosition, position.Plus(Vector{X: -1, Y: 0}), next-1)

	return sum
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
