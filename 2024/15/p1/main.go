package main

import (
	"flag"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/JeroenoBoy/AdventOfCode/utils"
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

	s := strings.Split(strings.Trim(string(input), "\n"), "\n\n")

	mapStrs := strings.Split(s[0], "\n")
	worldMap := utils.NewArray2D[NodeType](utf8.RuneCountInString(mapStrs[0]), len(mapStrs))
	playerPosition := utils.Vector{}

	for y, row := range mapStrs {
		for x, rune := range row {
			switch rune {
			case '@':
				playerPosition = utils.Vector{X: x, Y: y}
				break
			case '#':
				worldMap.Set(x, y, NodeTypeWall)
				break
			case 'O':
				worldMap.Set(x, y, NodeTypeBox)
				break
			case '.':
				worldMap.Set(x, y, NodeTypeEmpty)
				break
			}
		}
	}

	move := func(delta utils.Vector) {
		tempPos := playerPosition
		didHitEmptyPos := false
		tilesPassed := 0
		for !didHitEmptyPos {
			tempPos = tempPos.Add(delta)
			tilesPassed++
			rune := worldMap.Get(tempPos.X, tempPos.Y)

			switch rune {
			case NodeTypeWall:
				return
			case NodeTypeEmpty:
				didHitEmptyPos = true
				break
			}
		}

		for i := tilesPassed - 1; i > 0; i-- {
			fromTile := playerPosition.Add(delta.Multiply(i))
			toTile := fromTile.Add(delta)

			x := worldMap.Get(fromTile.X, fromTile.Y)
			worldMap.Set(toTile.X, toTile.Y, x)
			worldMap.Set(fromTile.X, fromTile.Y, NodeTypeEmpty)
		}

		playerPosition = playerPosition.Add(delta)
	}

	for _, command := range s[1] {
		switch command {
		case '>':
			move(utils.Vector{X: 1})
			break
		case '<':
			move(utils.Vector{X: -1})
			break
		case '^':
			move(utils.Vector{Y: -1})
			break
		case 'v':
			move(utils.Vector{Y: 1})
			break
		}
	}

	printWorldMap(worldMap, playerPosition)

	result := 0

	for y := range worldMap.Height {
		for x := range worldMap.Width {
			if worldMap.Get(x, y) != NodeTypeBox {
				continue
			}
			result += 100*y + x
		}
	}

	println(result)
}

type NodeType int

const (
	NodeTypeEmpty NodeType = iota
	NodeTypeWall  NodeType = iota
	NodeTypeBox   NodeType = iota
)

func printWorldMap(worldMap utils.Array2D[NodeType], playerPosition utils.Vector) {
	for y := range worldMap.Height {
		for x := range worldMap.Width {
			if playerPosition.X == x && playerPosition.Y == y {
				print("@")
				continue
			}

			switch worldMap.Get(x, y) {
			case NodeTypeWall:
				print("#")
				break
			case NodeTypeBox:
				print("O")
				break
			case NodeTypeEmpty:
				print(".")
				break
			}
		}
		print("\n")
	}
}
