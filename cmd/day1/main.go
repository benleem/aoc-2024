package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	leftSorted := sortList(left)
	rightSorted := sortList(right)
	sum := findSum(leftSorted, rightSorted)
	fmt.Printf("sum: %v\n", sum)
	similarity := findSimilarity(left, right)
	fmt.Printf("similarity: %v\n", similarity)
}

func readInput() ([]uint, []uint, error) {
	var left []uint
	var right []uint

	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, "   ")
		intLeft, err := strconv.Atoi(lineArray[0])
		if err != nil {
			return nil, nil, err
		}
		intRight, err := strconv.Atoi(lineArray[1])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, uint(intLeft))
		right = append(right, uint(intRight))
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(left) != len(right) {
		return nil, nil, fmt.Errorf("sorted list are not the same length bub")
	}

	return left, right, err
}

func sortList(slice []uint) []uint {
	var sorted []uint
	sort.Slice(slice, func(i int, j int) bool {
		return slice[i] < slice[j]
	})
	for _, v := range slice {
		sorted = append(sorted, v)
	}
	return sorted
}

func findSum(leftSorted []uint, rightSorted []uint) uint {
	var sum uint
	for i, v := range leftSorted {
		if leftSorted[i] > rightSorted[i] {
			sum += v - rightSorted[i]
		} else {
			sum += rightSorted[i] - v
		}
	}
	return sum
}

func findSimilarity(left []uint, right []uint) uint {
	var similarity uint
	similarityMap := map[uint]uint{}
	for _, vl := range left {
		for _, vr := range right {
			if vl == vr {
				similarityMap[vl] = similarityMap[vl] + 1
			}
		}
	}
	for num, reps := range similarityMap {
		similarity += num * reps
	}
	return similarity
}
