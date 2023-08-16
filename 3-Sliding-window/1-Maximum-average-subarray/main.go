package main

import "fmt"

func main(){
  fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3},4))
}

func findMaxAverage(nums []int, k int) float64 {
  maximum:=0
  sum:=0
  for i:=0;i<len(nums);i++{
    if i<k {
      sum+= nums[i]
      maximum = sum
    }else{
      sum-=nums[i-k]
      sum+=nums[i]
      if sum > maximum {
        maximum = sum
      }
    }
  }

  return float64(maximum)/float64(k)
}
