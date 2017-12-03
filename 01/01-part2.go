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

  // To have a ordered map, create an index array
  // which is as long as the map I created above
  var keys []int

  for k := range m {
    keys = append(keys, k)
  }

  // Sort the newly created index
  sort.Ints(keys)

  // Calcualte the steps ahead (based on specification in advent codeing #1)
  step := count / 2

  // Go through map and look if the current number matches the number which is "step" ahead
  // of the current one
  for _, k := range keys {
    fmt.Println("Current number ", m[k])
    if num := k + step; num < count {
      // if it matches, add the number to the sum
      if m[k] == m[k + step] {
        sum = sum + m[k]
      }
      // If the current index + step is larger then the array,
      // start counting from the beginning (circular array)
    } else {
      n := step - (count - k)

      if m[k] == m[n] {
        sum = sum + m[k]
      }
    }
  }

  fmt.Println(sum)
}
