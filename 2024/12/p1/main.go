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
	areaSize := make([]int, 1)
	areaBorder := make([]int, 1)

	regionIdx := 1
	for y, row := range farm {
		for x := range row {
			if regionMap.get(x, y) == 0 {
				makeRegion(farm, regionMap, regionIdx, x, y)
				areaSize = append(areaSize, 0)
				areaBorder = append(areaBorder, 0)
				regionIdx++
			}

			region := regionMap.get(x, y)
			areaSize[region] = areaSize[region] + 1
			border := areaBorder[region]
			if !regionMap.contains(x+1, y) {
				border++
			} else if regionMap.get(x+1, y) != region {
				border++
            }

			if !regionMap.contains(x-1, y) {
				border++
			} else if regionMap.get(x-1, y) != region {
				border++
            }

			if !regionMap.contains(x, y+1) {
				border++
			} else if regionMap.get(x, y+1) != region {
				border++
            }

			if !regionMap.contains(x, y-1) {
				border++
			} else if regionMap.get(x, y-1) != region {
				border++
			}
			areaBorder[region] = border
		}
	}

	result := 0
	for region, size := range areaSize {
		border := areaBorder[region]
		result += size * border
	}

	println(result)
}

func makeRegion(farm []string, regionMap Array2D[int], region, x, y int) {
	tile := farm[y][x]
	if regionMap.get(x, y) != 0 {
		return
	}

	regionMap.set(x, y, region)

	if regionMap.contains(x+1, y) && farm[y][x+1] == tile {
		makeRegion(farm, regionMap, region, x+1, y)
	}
	if regionMap.contains(x-1, y) && farm[y][x-1] == tile {
		makeRegion(farm, regionMap, region, x-1, y)
	}
	if regionMap.contains(x, y+1) && farm[y+1][x] == tile {
		makeRegion(farm, regionMap, region, x, y+1)
	}
	if regionMap.contains(x, y-1) && farm[y-1][x] == tile {
		makeRegion(farm, regionMap, region, x, y-1)
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
