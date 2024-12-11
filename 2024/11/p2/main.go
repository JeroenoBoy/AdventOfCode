package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

var (
	fileInput  = flag.String("i", "input.txt", "change the input file")
	iterations = flag.Int("c", 25, "the count of iterations")
)

func main() {
	flag.Parse()
	input, err := ioutil.ReadFile(*fileInput)
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.ReplaceAll(string(input), "\n", ""), " ")

	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	count := 0

	var spawn func(string, int)

	spawn = func(num string, iteration int) {
		wg.Add(1)
		mtx.Lock()
		count++
		mtx.Unlock()
		go func() {
            iterate(num, iteration, spawn)
            wg.Done()
        }()
	}

	for _, input := range s {
		if input == "\n" {
			continue
		}

		wg.Add(1)
		count++
		go func() {
			iterate(input, *iterations, spawn)
			wg.Done()
		}()
	}

	wg.Wait()

	println(count)
}

func iterate(input string, count int, spawnNew func(string, int)) {
	str := input
	num, _ := strconv.Atoi(str)
	for i := 0; i < count; i++ {
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

			spawnNew(bS, count-i-1)
			continue
		}
		num = num * 2024
		str = strconv.Itoa(num)
	}
}
