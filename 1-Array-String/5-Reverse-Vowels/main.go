package main

import (
	"fmt"
	"strings"
)


func main(){
    fmt.Println(reverseVowels("hello"))
}


const vowels string = "aeiou"

func reverseVowels(s string) string {
  buf := []rune(s)
  vowelsOrder := []byte{}

  for i := 0; i< len(s); i++{
    if strings.Contains(vowels,string(s[i])){
      vowelsOrder = append(vowelsOrder,s[i])
    }
  }

  j :=0
   for i:= len(s)-1; i>=0; i-- {

    if strings.Contains(vowels,string(s[i])){
      buf[i] = rune(vowelsOrder[j])
      j++
    }
  }
  return string(buf)
}
