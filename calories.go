package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getCalorieCounts() []int {
	dat, err := os.ReadFile("Calories.txt")
	check(err)

	elves := strings.Split(string(dat), "\n\n")

	var calorieCounts []int
	for _, elf := range elves {
		snacks := strings.Split(elf, "\n")

		totalSnackage := 0
		for _, snack := range snacks {
			calories, err := strconv.Atoi(snack)
			check(err)
			totalSnackage += calories
		}
		calorieCounts = append(calorieCounts, totalSnackage)
	}

	sort.Ints(calorieCounts)

	return calorieCounts
}

func day1() {
	calories := getCalorieCounts()
	fmt.Printf("The swolest elf has %d calories\n", calories[len(calories)-1])

	sum := 0
	for _, cal := range calories[len(calories)-3:] {
		sum += cal
	}

	fmt.Printf("The 3 swolests elfses have %d calories\n", sum)
}
