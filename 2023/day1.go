package main

import (
  "os"
  "bufio"
)

func main() {
  input, err := os.Open("testinput.txt")
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


}

func parseFile(filename string) []string {

}

func parseLine(line string) int {

}
