package main

import (
	"AdventOfCode24/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func calcMul(mul string) int {
	i := strings.Index(mul, "(")
	j := strings.Index(mul, ",")
	k := strings.Index(mul, ")")

	lhs, _ := strconv.Atoi(mul[i+1 : j])
	rhs, _ := strconv.Atoi(mul[j+1 : k])
	return lhs * rhs
}

func mulRegex(input string) []string {
	re := regexp.MustCompile(`mul\(\d*,\d*\)`)
	return re.FindAllString(input, -1)
}

func mulDoDontRegex(input string) [][]string {
	allPat := `(?s:(?:^|do\(\))(.*?)(?:don't\(\)|$))`
	allRe := regexp.MustCompile(allPat)
	allMatch := allRe.FindAllStringSubmatch(input, -1)
	return allMatch
}

func solvePuzzleOne(input string) int {
	sum := 0
	for _, match := range mulRegex(input) {
		sum += calcMul(match)
	}

	return sum
}

func calcMulFromRes(results [][]string) int {
	sum := 0

	for _, match := range results {
		sum += solvePuzzleOne(match[0])
	}
	return sum
}

func solvePuzzleTwo(input string) int {
	mulRes := mulDoDontRegex(input)
	return calcMulFromRes(mulRes)
}

func main() {
	input := utils.ReadText("input.txt")
	puzzleOne := solvePuzzleOne(input)
	fmt.Println("Sum of all multiplications is:", puzzleOne)

	puzzleTwo := solvePuzzleTwo(input)
	fmt.Println("Sum of all multiplications not preceded by don't is:", puzzleTwo)
}
