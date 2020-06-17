package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
)

func GetArrayOfSize(n int) []int {
	dir := "D:/Dev/cygwin/work/golang/golang-knowledge/Algorithms/Sort/utils"

	fname := filepath.Join(dir, "IntegerArray.txt")

	f, _ := os.Open(fname)

	defer f.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}

	return numbers[0:n]
}
