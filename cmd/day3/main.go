package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	parseInput(input)
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
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return input, nil
}

func parseInput(input string) []string {
	fmt.Println(input)
	for i, v := range input {
		fmt.Println(i, v)
	}
	// input = strings.FieldsFunc(line, func(r rune) bool {
	// 	if r == '(' || r == ')' {
	// 		return true
	// 	}
	// 	return false
	// })
	return nil
}
