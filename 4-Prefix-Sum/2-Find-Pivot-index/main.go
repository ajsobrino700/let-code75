package main

import "fmt"

func main(){
  fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
}

func pivotIndex(nums []int) int {
  pivotIndex := -1
  length := len(nums)
  sum:= nums[0]
  for i:=1; i < length; i++{
    sum+= nums[i]
    nums[i] = sum
  }

  if nums[0] == nums[length-1]{
    pivotIndex=0
  }

  for i:=1; i < length && pivotIndex == -1; i++{
    if nums[i-1] == nums[length-1]-nums[i] {
      pivotIndex = i
    }
  }

  return pivotIndex
}
