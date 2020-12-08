package lib

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Check(e error) {
	check(e)
}

func GetIntArrayInput(filename string) []int {
	result := make([]int, 0) // len(a)=5
	absFilePath, absFilePathErr := filepath.Abs(filename)
	check(absFilePathErr)

	file, err := os.Open(absFilePath)
	closeFile := func() {
		closeErr := file.Close()
		check(closeErr)
	}
	defer closeFile()

	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)

		result = append(result, i)
	}
	check(scanner.Err())

	return result
}

func ParseLines(filename string) []string {
	lines := make([]string, 0)

	absFilePath, absFilePathErr := filepath.Abs(filename)
	Check(absFilePathErr)

	file, err := os.Open(absFilePath)
	Check(err)
	closeFile := func() {
		closeErr := file.Close()
		check(closeErr)
	}
	defer closeFile()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}
