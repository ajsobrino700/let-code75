package main

import (
	"fmt"
	"strconv"
)

func main(){

  //printlnAndCompress([]byte{'a','a','a','b','b'})
  printlnAndCompress([]byte{'a','b','b','b','b','b','b','b','b'})
}



func compress(chars []byte)int{
  count := 1
  auxIndex :=0
  newIndex := 0;
  for i := 1 ; i <= len(chars); i++{
    fmt.Println(string(chars[auxIndex]))
    fmt.Println(string(chars[i]))
    if i != len(chars) && chars[auxIndex] == chars[i] {
      count++
    }else{
      chars[newIndex] = chars[auxIndex]
      newIndex++
      fillCount(chars,&newIndex,count)
      count = 1
      auxIndex = i
    }
  }
  return newIndex
}

func fillCount(chars []byte, newIndex *int, count int){

  s:= strconv.Itoa(count)
  //fmt.Println(s)
    for i:=0; i< len(s);i++{
      chars[*newIndex] = byte(s[i])
      *newIndex++
    }

}

func printlnAndCompress(chars []byte){
    length := compress(chars)
    for i := 0; i < length ; i++{
      fmt.Println(string(chars[i]))
    }

}



