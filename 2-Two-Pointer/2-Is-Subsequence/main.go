package main

import (
	"fmt"
)

func main(){
  fmt.Println(isSubsequence("abc","ahbgdc"))

  fmt.Println(isSubsequence("abc",""))

  fmt.Println(isSubsequence("","ahbgdc"))
}

func isSubsequence(s string, t string) bool {
    count:=0
    for i := 0; i<len(t) && count != len(s); i++{
      if t[i] == s[count] {
        count++
      }
    }
    return count == len(s)
}
