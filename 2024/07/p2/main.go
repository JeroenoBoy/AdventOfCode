package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := 0
	for _, calculation := range strings.Split(string(input), "\n") {
		if calculation == "" {
			continue
		}

		s := strings.Split(calculation, ": ")
		expectedResult, _ := strconv.Atoi(s[0])

		ns := strings.Split(s[1], " ")
		nums := make([]int, len(ns))
		for i, v := range ns {
			nums[i], _ = strconv.Atoi(v)
		}

		if hasValidOption(nums, 1, nums[0], expectedResult) {
			result += expectedResult
		}
	}

	println(result)
}

func hasValidOption(nums []int, i int, currentValue int, expectedResult int) bool {
	if i == len(nums) {
		return currentValue == expectedResult
	}

	if hasValidOption(nums, i+1, currentValue+nums[i], expectedResult) {
		return true
	}

	if hasValidOption(nums, i+1, currentValue*nums[i], expectedResult) {
		return true
	}

	v, _ := strconv.Atoi(strconv.Itoa(currentValue) + strconv.Itoa(nums[i]))
	if hasValidOption(nums, i+1, v, expectedResult) {
		return true
	}

	return false
}
