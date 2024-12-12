package main

import (
	"flag"
	"io/ioutil"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

var (
	fileInput         = flag.String("i", "input.txt", "change the input file")
	iterations        = flag.Int("c", 75, "the count of iterations")
	routineCount      = flag.Int("r", 100, "the count of routines to run")
	routineStartDepth = flag.Int("d", 25, "the depht when the multi threading starts")
)

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.ReplaceAll(string(input), "\n", ""), " ")

	stones := make([]string, 0, len(s))
	for _, input := range s {
		if input == "\n" {
			continue
		}

		stones = append(stones, input)
	}

	for i := 0; i < *routineStartDepth; i++ {
		for x := 0; x < len(stones); x++ {
			str := stones[x]
			if str == "0" {
				stones[x] = "1"
				continue
			}

			strLen := utf8.RuneCountInString(str)
			if strLen >= 2 && strLen%2 == 0 {
				c := utf8.RuneCountInString(str) / 2
				a, _ := strconv.Atoi(str[:c])
				b, _ := strconv.Atoi(str[c:])

				aS := strconv.Itoa(a)
				bS := strconv.Itoa(b)

				stones = slices.Insert(stones, x+1, bS)
				stones[x] = aS

				x++
				continue
			}

			num, _ := strconv.Atoi(stones[x])
			num *= 2024
			stones[x] = strconv.Itoa(num)
		}
	}

	count := 0
	var wg sync.WaitGroup
	var mtx sync.Mutex
	sem := make(chan struct{}, *routineCount)

	for _, input := range stones {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			itCount := iterate(input, *iterations-*routineStartDepth)
			mtx.Lock()
			count += itCount
			mtx.Unlock()
			<-sem
			wg.Done()
		}()
	}

	wg.Wait()

	println(count)
}

func iterate(input string, loopDepth int) int {
	str := input
	num, _ := strconv.Atoi(str)

	count := 1
	for i := 0; i < loopDepth; i++ {
		if num == 0 {
			num = 1
			str = "1"
			continue
		}

		rc := utf8.RuneCountInString(str)
		if rc%2 == 0 && rc >= 2 {
			c := utf8.RuneCountInString(str) / 2
			a, _ := strconv.Atoi(str[:c])
			b, _ := strconv.Atoi(str[c:])

			aS := strconv.Itoa(a)
			bS := strconv.Itoa(b)

			num = a
			str = aS

			count += iterate(bS, loopDepth-i-1)
			continue
		}
		num = num * 2024
		str = strconv.Itoa(num)
	}

	return count
}
