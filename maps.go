package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
  "fmt"
)

func WordCount(s string) map[string]int {
  words := strings.Fields(s)
  words_map := make(map[string]int)

  for _, word := range words {
    words_map[word]++
  }
  fmt.Println(words_map)
  return words_map
}

func main() {
    wc.Test(WordCount)
}
