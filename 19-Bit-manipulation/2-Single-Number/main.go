package main

import "fmt"

func main(){

  fmt.Println(singleNumber([]int{2,2,1}))

}


func singleNumber(nums []int) int {
  numberMap := map[int]int{}
  for i:=0; i < len(nums); i++{
    currentValue := nums[i]
    value,ok :=numberMap[currentValue]
    if ok {
      value++
      numberMap[currentValue]= value
    }else {
      numberMap[currentValue] = 1
    }
  }

  for k,v:= range numberMap{
    if v == 1{
      return k
    }
  }

  return 0
}
