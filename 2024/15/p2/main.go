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
	worldMap := utils.NewArray2D[NodeType](utf8.RuneCountInString(mapStrs[0])*2, len(mapStrs))
	playerPosition := utils.Vector{}

	for y, row := range mapStrs {
		for x, rune := range row {
			switch rune {
			case '@':
				playerPosition = utils.Vector{X: x * 2, Y: y}
				break
			case '#':
				worldMap.Set(x*2, y, NodeTypeWall)
				worldMap.Set(x*2+1, y, NodeTypeWall)
				break
			case 'O':
				worldMap.Set(x*2, y, NodeTypeBoxL)
				worldMap.Set(x*2+1, y, NodeTypeBoxR)
				break
			case '.':
				worldMap.Set(x*2, y, NodeTypeEmpty)
				worldMap.Set(x*2+1, y, NodeTypeEmpty)
				break
			}
		}
	}

	var canMove func(utils.Vector, utils.Vector, *[]utils.Vector) bool
	canMove = func(position utils.Vector, delta utils.Vector, interactedTiles *[]utils.Vector) bool {
		point := position.Add(delta)
		nodeType := worldMap.Get(point.X, point.Y)

		if nodeType == NodeTypeEmpty {
			return true
		}
		if nodeType == NodeTypeWall {
			return false
		}

		*interactedTiles = append(*interactedTiles, point)

		if nodeType == NodeTypeBoxL && delta.Y != 0 {
			*interactedTiles = append(*interactedTiles, point.Add(utils.Vector{X: 1}))
			a := canMove(point, delta, interactedTiles)
			b := canMove(point.Add(utils.Vector{X: 1}), delta, interactedTiles)
			return a && b
		}
		if nodeType == NodeTypeBoxR && delta.Y != 0 {
			*interactedTiles = append(*interactedTiles, point.Add(utils.Vector{X: -1}))
			a := canMove(point, delta, interactedTiles)
			b := canMove(point.Add(utils.Vector{X: -1}), delta, interactedTiles)
			return a && b
		}
		if nodeType == NodeTypeBoxL && delta.X == 1 {
			*interactedTiles = append(*interactedTiles, point.Add(delta))
			return canMove(point.Add(delta), delta, interactedTiles)
		}
		if nodeType == NodeTypeBoxR && delta.X == -1 {
			*interactedTiles = append(*interactedTiles, point.Add(delta))
			return canMove(point.Add(delta), delta, interactedTiles)
		}
		panic("HI!?")
	}

	move := func(delta utils.Vector) {

		interactedTiles := make([]utils.Vector, 0)
		if !canMove(playerPosition, delta, &interactedTiles) {
			return
		}

		movedTiles := make(map[utils.Vector]bool)
		for i := len(interactedTiles) - 1; i >= 0; i-- {
			fromTile := interactedTiles[i]
			if movedTiles[fromTile] {
				continue
			}
			movedTiles[fromTile] = true
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
			if worldMap.Get(x, y) != NodeTypeBoxL {
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
	NodeTypeBoxL  NodeType = iota
	NodeTypeBoxR  NodeType = iota
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
			case NodeTypeBoxL:
				print("[")
				break
			case NodeTypeBoxR:
				print("]")
				break
			case NodeTypeEmpty:
				print(".")
				break
			}
		}
		print("\n")
	}
}
