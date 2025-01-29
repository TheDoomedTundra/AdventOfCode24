package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadData(fileName string) [][]string {
	file, err := os.Open(fileName)
	CheckErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	CheckErr(err)

	return records
}

func ReadText(fileName string) string {
	text, err := os.ReadFile(fileName)
	CheckErr(err)
	return string(text)
}

func StringToIntArr(arr []string) []int {
	intArr := make([]int, len(arr))
	for i, s := range arr {
		num, err := strconv.Atoi(s)
		CheckErr(err)
		intArr[i] = num
	}
	return intArr
}
