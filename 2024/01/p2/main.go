package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(input), "\n")

	match := regexp.MustCompile("^(\\d+) +(\\d+)$")

	aArray := make([]int, 99999)
	bArray := make([]int, 99999)

	for _, line := range inputs {
		if len(line) == 0 {
			continue
		}

		matches := match.FindStringSubmatch(line)

		a, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		aArray[a] = aArray[a] + 1
		bArray[b] = bArray[b] + 1
	}

    result := 0

    for i := range len(aArray) {
        result += i * aArray[i] * bArray[i]
    }

	log.Println(result)
}

func catch(arr []int, indx *int, to int) {
	for *indx < len(arr) && arr[*indx] < to {
		*indx++
	}
}

func count(arr []int, indx *int, target int) int {
	if arr[*indx] != target {
		return 0
	}
	i := 1

	for *indx < len(arr)-1 && arr[*indx+1] == target {
		*indx++
	}

	return i
}
