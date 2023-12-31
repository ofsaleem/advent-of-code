package main

import (
  "os"
  "bufio"
  "strings"
)

func main() {
  lines := parseFile("./probleminput.txt")
  var total int
  for _, line := range lines {
    total += parseLine(line)
  }
  print(total, "\n")
}

func parseFile(filename string) []string {
  input, err := os.Open(filename)
    if err != nil {
      panic(err)
    }

  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanLines)
  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  input.Close()
  return lines
}

func parseLine(line string) int {
  print(line, "\n")
  numMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, 
    "six": 6, "seven": 7, "eight": 8, "nine": 9, "1": 1, "2": 2, "3": 3, "4": 4,
    "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
  min := [2]int{65535, 0}
  max := [2]int{0, 0}
  for str, int := range numMap {
    if strings.Contains(line, str) {
      if min[0] >= strings.Index(line, str) {
        min[0] = strings.Index(line, str)
        min[1] = int
      }
      if max[0] <= strings.LastIndex(line, str) {
        max[0] = strings.LastIndex(line, str)
        max[1] = int
      }
    }
  }
  if min[1] != 0 && max[1] != 0 {
    val := 10 * min[1] + max[1]
    print(val, "\n")
    return val
  } else {
    return min[1]
  }
}

