package main

import "fmt"


func main(){
  fmt.Println(productExceptSelf([]int{1,2,3,4}))
  fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}


func productExceptSelf(nums []int) []int {
  numbersZero := 0;
  totalProduct := 1;
  for i:=0; i< len(nums); i++{
    if nums[i] == 0{
      numbersZero++
    }else{
      totalProduct*=nums[i]
    }
  }
  result := []int{}
  value := 0

  fmt.Println(numbersZero)
  for i:=0;i<len(nums); i++ {
    value = 0
    if nums[i] == 0 && numbersZero == 1{
      value = totalProduct
    }else if numbersZero == 0{
      value = totalProduct/nums[i]
    }
    result = append(result, value)
  }
  return result
}
