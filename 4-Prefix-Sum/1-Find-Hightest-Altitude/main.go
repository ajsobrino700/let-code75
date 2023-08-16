package main

import "fmt"

func main(){
 fmt.Println(largestAltitude([]int{-5,1,5,0,-7}))
}

func largestAltitude(gain []int) int {
  highest := 0
  sum:= 0

  for i:=0;i<len(gain);i++{
    sum += gain[i]
    if sum > highest {
      highest = sum
    }
  }


  return highest
}
