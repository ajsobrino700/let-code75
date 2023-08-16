package main

import (
	"fmt"
	"strings"
)

func main(){

  //fmt.Println(reverseWords("the sky is blue"))
  fmt.Println(reverseWords("a good  example"))
}


func reverseWords(s string) string {
  var result string
  words := strings.Fields(s)

  length := len(words)
  for i := length-1; i >= 0 ; i--{
    result += words[i]+" "
  }

  return strings.TrimSpace(result)
}
