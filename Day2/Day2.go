package main

import (
	"AdventOfCode24/utils"
	"fmt"
	"slices"
	"strings"
)

func isSortedDecr(arr []int) bool {
	decreasing := slices.IsSortedFunc(arr, func(a, b int) int {
		if b == a {
			return 0
		} else if b > a {
			return 1
		} else {
			return -1
		}
	})
	return decreasing
}

func isSortedIncr(arr []int) bool {
	increasing := slices.IsSortedFunc(arr, func(a, b int) int {
		if b == a {
			return 0
		} else if b < a {
			return 1
		} else {
			return -1
		}
	})
	return increasing
}

func isReportSafe(report []string) int {
	intReport := utils.StringToIntArr(strings.Fields(report[0]))
	safe := 1
	sorted := isSortedIncr(intReport) || isSortedDecr(intReport)
	if !sorted {
		safe = 0
	}

	for i := 0; i < len(intReport)-1; i++ {
		diff := utils.Abs(intReport[i+1] - intReport[i])
		if diff > 3 || diff < 1 {
			safe = 0
			break
		}
	}

	return safe
}

func removeIndex(report []int, idx int) []int {
	if idx < 0 || idx >= len(report) {
		return report
	}
	newReport := make([]int, 0, len(report)-1)
	newReport = append(newReport, report[:idx]...)
	newReport = append(newReport, report[idx+1:]...)
	return newReport
}

func oppositeSigns(x int, y int) bool {
	if (x > 0 && y < 0) || (x < 0 && y > 0) {
		return true
	}
	return false
}

func isSafeWithOneRemovalImpl(report []int, level int) int {
	// Bad conditions:
	/***
	1. The difference is 0 or > 3
	3. Decrease more than once with more than one increases
	4. Increase more than once with more than one decreases
	*/
	removed := 0
	prevDiff := 0

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if utils.Abs(diff) > 3 || diff == 0 || (prevDiff != 0 && oppositeSigns(diff, prevDiff)) {
			tmpRemoved := isSafeWithOneRemovalImpl(removeIndex(report, i), level+1)
			if i == 1 {
				tmpRemoved3 := isSafeWithOneRemovalImpl(removeIndex(report, i-1), level-1)
				tmpRemoved = min(tmpRemoved, tmpRemoved3)
			}
			if (i + 1) < len(report) {
				tmpRemoved2 := isSafeWithOneRemovalImpl(removeIndex(report, i+1), level+1)
				tmpRemoved = min(tmpRemoved, tmpRemoved2)
			}
			tmpRemoved++

			if removed != 0 {
				removed = min(removed, tmpRemoved)
			} else {
				removed = tmpRemoved
			}

		}
		prevDiff = diff
	}

	return removed
}

func isSafeWithOneRemoval(report []string) int {
	intReport := utils.StringToIntArr(strings.Fields(report[0]))
	removed := isSafeWithOneRemovalImpl(intReport, 1)

	safe := 1
	if removed > 1 {
		safe = 0
	}
	return safe

}

func solvePuzzle1(reports [][]string) int {
	safeReports := 0

	for _, report := range reports {
		safeReports += isReportSafe(report)
	}

	return safeReports
}

func solvePuzzle2(reports [][]string) int {
	safeReports := 0

	for _, report := range reports {
		safeReports += isSafeWithOneRemoval(report)
	}

	return safeReports
}

func main() {
	reports := utils.ReadData("input.csv")
	puzzleOne := solvePuzzle1(reports)
	fmt.Println("There are", puzzleOne, "safe reports")

	puzzleTwo := solvePuzzle2(reports)
	fmt.Println("There are", puzzleTwo, "safe reports with one removal")

}
