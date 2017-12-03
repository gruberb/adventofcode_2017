package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanRunes)

	count := 0
	var first int
	var previous int
	var current int
	var sum int

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())

		if count == 0 {
			first = number
			count++
		}

		current = number

		if current == previous {
			sum = sum + current
		}

		previous = current
	}

	if first == current {
		sum = sum + current
	}

	fmt.Println(sum)

}
