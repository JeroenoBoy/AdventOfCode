package main

import (
	"flag"
	"io/ioutil"
	"strconv"
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

	mem := make([]int, 0)
	for i, rune := range input {
		if rune == '\n' {
			continue
		}
		id := i / 2

		num, err := strconv.Atoi(string(rune))
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			for i := 0; i < num; i++ {
				mem = append(mem, id)
			}
		} else {
			for i := 0; i < num; i++ {
				mem = append(mem, -1)
			}
		}
	}
	maxI := len(mem)
	sum := 0
	for i := 0; i < maxI; i++ {
		if mem[i] == -1 {
			maxI--
			for mem[maxI] == -1 {
				maxI--
			}
			if i >= maxI {
				break
			}
			sum += i * mem[maxI]
		} else {
			sum += i * mem[i]
		}
	}

	println(sum)
}
