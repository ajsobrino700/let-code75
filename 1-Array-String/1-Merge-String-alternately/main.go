package main

import "fmt"


func main(){
  fmt.Println(mergeAlternately("abc","qpd"))
  fmt.Println(mergeAlternately("ab","pqrs"))
}


func mergeAlternately(wordOne string, wordTwo string) string{
  lengthOne := len(wordOne)
  lengthTwo := len(wordTwo)


  var totalLength int = lengthOne
  if lengthOne < lengthTwo {
    totalLength = lengthTwo
  }

  var result = []byte {}
  for i := 0 ; i < totalLength; i++ {
    result = append(result,getLetter(wordOne, i),getLetter(wordTwo,i))
  }

  return string(result)
}

func getLetter(word string, index int) byte{
  var result byte = byte(0)
  if len(word) > index{
    result= word[index]
  }

  return result
}
