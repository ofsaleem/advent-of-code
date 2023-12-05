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
  print(total)
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
  var total int
  for _, c := range line {
    if unicode.IsDigit(c) { 
      total += int(c)
    }
  }
  return total
}
