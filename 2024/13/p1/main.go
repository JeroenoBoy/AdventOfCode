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

			cost := clawMachine.CalculateCheapestOption(Vector{})
			if cost == -1 {
				return
			}

			mtx.Lock()
			result += cost
			println("finished", i, "out of", count)
			mtx.Unlock()
		}()

		wg.Wait()
	}

	wg.Wait()

	println(result)
}

type ClawMachine struct {
	ButtonA Vector
	ButtonB Vector
	Prize   Vector
}

func (m *ClawMachine) CalculateCheapestOption(currentPosition Vector) int {
	if currentPosition.X > m.Prize.X || currentPosition.Y > m.Prize.Y {
		return -1
	}

	if currentPosition.X == m.Prize.X && currentPosition.Y == m.Prize.Y {
		return 0
	}

	costB := m.CalculateCheapestOption(currentPosition.Add(m.ButtonB))
	costA := m.CalculateCheapestOption(currentPosition.Add(m.ButtonA))

	if costA == -1 {
		if costB == -1 {
			return -1
		}
		return costB + COST_B
	} else if costB == -1 {
		return costA + COST_A
	}

	costA += COST_A
	costB += COST_B

	if costA > costB {
		return costB
	}

	return costA
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
