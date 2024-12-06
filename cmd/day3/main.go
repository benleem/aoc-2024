package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	inputSlice, err := parseInput(input)
	if err != nil {
		log.Fatalln(err)
	}
	sum := findSum(inputSlice)
	fmt.Println(sum)
}

func readInput() (string, error) {
	var input string
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return input, nil
}

func parseInput(input string) ([][]uint, error) {
	re, err := regexp.Compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	if err != nil {
		return nil, err
	}
	parsed := re.FindAllString(input, -1)

	var inputSlice [][]uint
	var enabled bool = true
	for _, v := range parsed {
		var uintSlice []uint
		var cutString string
		if v == "do()" {
			enabled = true
		}
		if v == "don't()" {
			enabled = false
		}
		if v != "do()" && v != "don't()" {
			cutString, _ = strings.CutPrefix(v, "mul(")
			cutString, _ = strings.CutSuffix(cutString, ")")
			stringSlice := strings.Split(cutString, ",")
			for _, v := range stringSlice {
				converted, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				uintSlice = append(uintSlice, uint(converted))
			}
			if enabled == true {
				inputSlice = append(inputSlice, uintSlice)
			}
		}
	}
	return inputSlice, nil
}

func findSum(input [][]uint) uint {
	var sum uint
	for _, v := range input {
		product := v[0] * v[1]
		sum += product
	}
	return sum
}
