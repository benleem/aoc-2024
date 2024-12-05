package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	sum := sumSafe(input)
	fmt.Println(sum)
}

func readInput() ([][]uint, error) {
	var input [][]uint
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var uintArray []uint
		line := scanner.Text()
		stringArray := strings.Split(line, " ")
		for _, v := range stringArray {
			converted, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			uintArray = append(uintArray, uint(converted))
		}
		input = append(input, uintArray)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}

func sumSafe(input [][]uint) uint {
	var sum uint
	for _, v := range input {
		if isSafe(v) {
			sum += 1
		} else if dampenSafe(v) {
			sum += 1
		}
	}
	return sum
}

func isSafe(inputRow []uint) bool {
	var least uint = 1
	var most uint = 3
	var direction bool
	if inputRow[0] > inputRow[1] {
		// [3,2,1]
		// false - row is decreasing
		direction = false
	} else if inputRow[0] < inputRow[1] {
		// [1,2,3]
		// true - row is increasing
		direction = true
	} else {
		return false
	}
	for i := 0; i < len(inputRow)-1; i++ {
		var diff int
		var diffAbs uint
		// positive - row decreasing
		diff += int(inputRow[i]) - int(inputRow[i+1])
		diffAbs = uint(math.Abs(float64(diff)))
		if direction && diff > 0 || !direction && diff < 0 {
			return false
		}
		if diffAbs < least || diffAbs > most {
			return false
		}
	}
	return true
}

func dampenSafe(inputRow []uint) bool {
	for i, _ := range inputRow {
		dampRow := append([]uint{}, inputRow[:i]...)
		dampRow = append(dampRow, inputRow[i+1:]...)
		if isSafe(dampRow) {
			return true
		}
	}
	return false
}
