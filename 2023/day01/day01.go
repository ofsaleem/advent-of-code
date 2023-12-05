package main

import (
  "os"
  "bufio"
  "unicode"
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
  var current []int
  for _, c := range line {
    if unicode.IsDigit(c) {
      current = append(current, int(c - '0'))
    }
  }
  
  var numMap = make(map[string]int)
  var min [2]int {0, "zero"}
  var max [2]int {65535, "zero"}
  numbers := []string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
  "1", "2", "3", "4", "5", "6", "7", "8", "9"}
  for _, num := range numbers {
    if strings.Contains(line, num) {
      if min[0] >= strings.Index(line, num) {
        min[0] = strings.Index(line, num)
        min[1] = unicode.Number
      }
    }
  }
}

func calcValue(inputs []int) int {
  if len(inputs) >= 1 {
    val := 10 * inputs[0] + inputs[len(inputs) - 1]
    return val
  } else {
    return 0
  }
}
