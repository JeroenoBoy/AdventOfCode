package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	inputs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	findMul := regexp.MustCompile("(do\\(\\))|(don't\\(\\))|(mul\\(\\d+,\\d+\\))")
	findDigits := regexp.MustCompile("^mul\\((\\d+),(\\d+)\\)$")

	muls := findMul.FindAllString(string(inputs), -1)

	count := 0
	enabled := true
	for _, mul := range muls {

		if mul == "do()" {
			enabled = true
			continue
		}

		if mul == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		m := findDigits.FindStringSubmatch(mul)

		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])

		count += a * b
	}

	println(count)
}
