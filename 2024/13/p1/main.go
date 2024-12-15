package main

import (
	"flag"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	fileInput = flag.String("i", "input.txt", "change the input file")

	buttonARegex = regexp.MustCompile("(?m)Button A: X\\+(\\d+), Y\\+(\\d+)")
	buttonBRegex = regexp.MustCompile("(?m)Button B: X\\+(\\d+), Y\\+(\\d+)")
	prizeRegex   = regexp.MustCompile("(?m)Prize: X=(\\d+), Y=(\\d+)")

	COST_A = 3
	COST_B = 1
)

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	result := 0
	var wg sync.WaitGroup
	var mtx sync.Mutex

	s := strings.Split(string(input), "\n\n")
	count := len(s)
	for i, machineInput := range s {
		wg.Add(1)

		go func() {
			defer wg.Done()
			buttonA := buttonARegex.FindStringSubmatch(machineInput)
			buttonB := buttonBRegex.FindStringSubmatch(machineInput)
			prize := prizeRegex.FindStringSubmatch(machineInput)

			bAX, _ := strconv.Atoi(buttonA[1])
			bAY, _ := strconv.Atoi(buttonA[2])
			bBX, _ := strconv.Atoi(buttonB[1])
			bBY, _ := strconv.Atoi(buttonB[2])
			pX, _ := strconv.Atoi(prize[1])
			pY, _ := strconv.Atoi(prize[2])

			clawMachine := ClawMachine{
				ButtonA: Vector{X: bAX, Y: bAY},
				ButtonB: Vector{X: bBX, Y: bBY},
				Prize:   Vector{X: pX, Y: pY},
			}

			cost := clawMachine.CalculateCheapestOption()
			println("finished", i, "out of", count, "with value", cost)
			if cost == -1 {
				return
			}

			mtx.Lock()
			result += cost
			mtx.Unlock()
		}()
	}

	wg.Wait()

	println(result)
}

type ClawMachine struct {
	ButtonA Vector
	ButtonB Vector
	Prize   Vector
}

func (m *ClawMachine) CalculateCheapestOption() int {
	costs := make(map[Vector]int)
	costs[Vector{}] = 0

	hasUnresolvedPositions := true
	lowestCost := 999999999999999999
	for hasUnresolvedPositions {
		newCosts := make(map[Vector]int)

		hasUnresolvedPositions = false
		for position, value := range costs {
			if position.X >= m.Prize.X || position.Y >= m.Prize.Y {
				if value < lowestCost && position.X == m.Prize.X && position.Y == m.Prize.Y {
					lowestCost = value
				}
				continue
			}

			hasUnresolvedPositions = true

			newPosA := position.Add(m.ButtonA)
			newPosB := position.Add(m.ButtonB)

			if cost, ok := newCosts[newPosA]; !ok {
				newCosts[newPosA] = value + COST_A
			} else if value+COST_A < cost {
				newCosts[newPosA] = value + COST_A
			}

			if cost, ok := newCosts[newPosB]; !ok {
				newCosts[newPosB] = value + COST_B
			} else if value+COST_A < cost {
				newCosts[newPosB] = value + COST_B
			}

		}

		costs = newCosts
	}

	if lowestCost == 999999999999999999 {
		return -1
	}

	return lowestCost
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
