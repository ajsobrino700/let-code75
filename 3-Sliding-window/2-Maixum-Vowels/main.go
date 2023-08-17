package main

import (
	"fmt"
	"strings"
)


func main(){
  fmt.Println(maxVowels("abciiidef",3))
  fmt.Println(maxVowels("aeiou",2))
  fmt.Println(maxVowels("leetcode",3))
}



func maxVowels(s string, k int) int {

  count := 0
  max := 0


  for i:=0;i< len(s); i++ {
    if i<k{
      if isVowels(s[i]) == 1{
        count++
        max++
      }
    }else {
      count+= isNotVowels(s[i-k])
      if isVowels(s[i]) == 1 {
        count++
      }
      if max < count {
        max = count
        if max == k {
          break
        }
      }
    }
  }

  return max

}

func isVowels(letter byte)int {
  result := -1
  vowels:="aeiouAEIOU"
  if strings.Contains(vowels,string(letter)) {
    result = 1
  }
  return result
}

func isNotVowels(letter byte) int {
    result := 0
    if isVowels(letter) == 1{
        result = -1
  }
  return result
}

