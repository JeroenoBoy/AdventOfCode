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

	input, _ := ioutil.ReadFile(*fileInput)

	farm := strings.Split(strings.Trim(string(input), "\n"), "\n")
	regionMap := newArray2D[int](len(farm), utf8.RuneCountInString(farm[0]))

	regionIdx := 1
	for y, row := range farm {
		for x := range row {
			if regionMap.get(x, y) == 0 {
				getRegionTiles(farm, regionMap, regionIdx, x, y)
				regionIdx++
			}
		}
	}

	result := 0

	println(result)
}

type Region struct {
	Size  int
	Area  int
	Tiles []Vector
}

func newRegion(farm []string, regionMap Array2D[int], region, x, y int) {
}

func getRegionTiles(farm []string, regionMap Array2D[int], region, x, y int) {
	tile := farm[y][x]
	if regionMap.get(x, y) != 0 {
		return
	}

	regionMap.set(x, y, region)

	if regionMap.contains(x+1, y) && farm[y][x+1] == tile {
		getRegionTiles(farm, regionMap, region, x+1, y)
	}
	if regionMap.contains(x-1, y) && farm[y][x-1] == tile {
		getRegionTiles(farm, regionMap, region, x-1, y)
	}
	if regionMap.contains(x, y+1) && farm[y+1][x] == tile {
		getRegionTiles(farm, regionMap, region, x, y+1)
	}
	if regionMap.contains(x, y-1) && farm[y-1][x] == tile {
		getRegionTiles(farm, regionMap, region, x, y-1)
	}
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
