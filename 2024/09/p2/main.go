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

	for i := len(mem) - 1; i >= 0; i-- {
		id := mem[i]
		if id == -1 {
			continue
		}

		blockSize := 1
		for ; i-blockSize >= 0 && mem[i-blockSize] == id; blockSize++ {
			// no body needed
		}

        println(i, id, blockSize)

		fI := 0
		for ; fI < i-blockSize; fI++ {
			if mem[i] != -1 {
				continue
			}

			fBlockSize := 0
			for ; fI+fBlockSize < i-blockSize && mem[fI+fBlockSize] == -1; fBlockSize++ {
				// no body needed
			}

			if fBlockSize >= blockSize {
				goto moveBlock
			}
		}
		i -= blockSize - 1
		continue
	moveBlock:
		for blockIndex := 0; blockIndex < blockSize; blockIndex++ {
			mem[fI+blockIndex] = id
			mem[i-blockIndex] = -1
		}
		i += blockSize - 1
	}

	sum := 0
	for i := 0; i < len(mem); i++ {
		print(mem[i])
		if mem[i] == -1 {
			continue
		}
		sum += i * mem[i]
	}

	print("\n")
	println(sum)
}
