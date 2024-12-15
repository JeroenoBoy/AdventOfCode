package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	fileInput   = flag.String("i", "input.txt", "change the input file")
	amountInput = flag.Int("a", 10_000_000_000_000, "amount to add to the price input")

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
				ButtonAPosition: Vector{X: bAX, Y: bAY},
				ButtonBPosition: Vector{X: bBX, Y: bBY},
				PricePosition:   Vector{X: pX + *amountInput, Y: pY + *amountInput},
			}

			cost := clawMachine.CalculateCheapestOption()
			log.Println("finished", i, "out of", count, "with value", cost)
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
	ButtonAPosition Vector
	ButtonBPosition Vector
	PricePosition   Vector
}

func (m *ClawMachine) CalculateCheapestOption() int {
	cx := m.PricePosition.X / m.ButtonAPosition.X
	cy := m.PricePosition.Y / m.ButtonAPosition.Y
	c := int(math.Ceil(min(float64(cx), float64(cy))))

	lowestCost := 999999999999999999

	for i := 0; i <= c; i++ {

		if i%1_000_000_000 == 0 {
			log.Println(int(math.Round(float64(i)/float64(c)*100)), "%")
		}

		cost := i * COST_A
		position := m.ButtonAPosition.Multiply(i)

		posLeft := m.PricePosition.Remove(position)
		cx := posLeft.X / m.ButtonBPosition.X
		cy := posLeft.Y / m.ButtonBPosition.Y
		c := int(math.Ceil(min(float64(cx), float64(cy))))

		position = position.Add(m.ButtonBPosition.Multiply(c))
		cost += c * COST_B

		if cost < lowestCost && position == m.PricePosition {
			lowestCost = cost
		}
	}

	if lowestCost == 999999999999999999 {
		return -1
	}

	return lowestCost
}

func (m *ClawMachine) CalculateOption(position Vector, aOnly bool) int {

	if position == m.PricePosition {
		return 0
	}

	if position.X > m.PricePosition.X || position.Y > m.PricePosition.Y {
		return -1
	}

	costB := -1
	if !aOnly {
		costB = m.CalculateOption(position.Add(m.ButtonBPosition), false) + COST_B
	}

	costA := m.CalculateOption(position.Add(m.ButtonAPosition), true) + COST_A

	if costA == COST_A-1 {
		if costB == COST_B-1 {
			return -1
		}
		return costB
	}
	if costB == COST_B-1 {
		return costA
	}

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

func (n Vector) Multiply(multi int) Vector {
	n.X *= multi
	n.Y *= multi
	return n
}

func (n Vector) isAnyAxisGreaterOrEqual(other Vector) bool {
	return n.X >= other.X || n.Y >= other.Y
}
