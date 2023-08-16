package main

import "fmt"


func main(){
  fmt.Println(maxOperations([]int{1,2,3,4},5))

}

func maxOperations(nums []int, k int) int {
  operations :=0
  length := len(nums)
  for i:=0; i<length; i++{
    currentValue := nums[i]
    for j:=i+1; j < length; j++ {
      if currentValue + nums[j] == k {
        nums[j] = nums[length-1]
        length--
        operations++
        break
      }
    }
  }



  return operations
}
