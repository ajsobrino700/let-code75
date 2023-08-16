package main

import "fmt"

func main(){

  numbers:=[]int{0,1,0,3,12}
  moveZeroes(numbers)
  fmt.Println(numbers)
}


func moveZeroes(nums []int){

  length := len(nums)
  i := 0;
  upZeros := 0;
  for i < length && upZeros < length{
    if nums[i] == 0 {
      zeroUp(nums, i)
    }else {
      i++
    }
    upZeros++
  }

}

func zeroUp(nums []int, index int){

  for i:=index; i < len(nums)-1; i++{
    aux := nums[i]
    nums[i] = nums[i+1]
    nums[i+1] = aux
  }
}

