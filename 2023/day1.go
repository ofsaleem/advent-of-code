package main

import (
  "os"
  "bufio"
  "unicode"
)

func main() {
  lines := parseFile("./testinput.txt")
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
      current = append(current, int(c))
    }
  }
  if len(current) >= 2 {
    return 10 * current[0] + current[len(current) - 1]
  } else {
    return 0
  }
}
