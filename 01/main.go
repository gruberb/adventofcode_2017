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

	var previous int
	var sum int

	for scanner.Scan() {
		
		number,_ := strconv.Atoi(scanner.Text())
		fmt.Println("current number %s", number)

		current := number

		if current == previous {
			sum = sum + current
		}

		fmt.Println("Old previous %s", previous)
		previous = current
		fmt.Println("New previous %s", previous)
		fmt.Println(sum)
	}

}