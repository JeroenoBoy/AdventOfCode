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

	findMul := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
    findDigits := regexp.MustCompile("^mul\\((\\d+),(\\d+)\\)$")
        
	muls := findMul.FindAllString(string(inputs), -1)

    count := 0
    for _, mul := range muls {
        m := findDigits.FindStringSubmatch(mul)

        a, _ := strconv.Atoi(m[1])
        b, _ := strconv.Atoi(m[2])

        count += a * b
    }

    println(count)
}
