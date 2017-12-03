package main

import (
	"bufio"
	"fmt"
  "sort"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanRunes)

  var m map[int]int
  count := 0
  sum := 0

  m = make(map[int]int)

	for scanner.Scan() {
		number,_ := strconv.Atoi(scanner.Text())

    m[count] = number
    count++
	}

  var keys []int

  for k := range m {
    keys = append(keys, k)
  }

  sort.Ints(keys)

  step := count / 2

  for _, k := range keys {
    fmt.Println("Current number ", m[k])
    if num := k + step; num < count {
      if m[k] == m[k + step] {
        sum = sum + m[k]
      }
    } else {
      n := step - (count - k)

      if m[k] == m[n] {
        sum = sum + m[k]
      }
    }
  }

  fmt.Println(sum)
}
