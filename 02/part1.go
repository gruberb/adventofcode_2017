package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    fileHandle, _ := os.Open("input.txt")
    defer fileHandle.Close()

    fileScanner := bufio.NewScanner(fileHandle)

    for fileScanner.Scan() {
        text := fileScanner.Text()
        fmt.Println(text)
        fmt.Println("next line")
    }
}