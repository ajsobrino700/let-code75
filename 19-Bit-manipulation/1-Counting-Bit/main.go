package main

import "fmt"

func main(){

  fmt.Println(countBits(2))
  fmt.Println(countBits(5))

}

func countBits(n int) []int {
  result := []int {}

  for i:=0; i<= n; i++{
    result = append(result,convertToBit(i))
  }

  return result
}


func convertToBit(nums int) int{
  result:=0
  for nums != 0{
    if nums%2==1{
      result++
    }
    nums/=2
  }
  return result
}
