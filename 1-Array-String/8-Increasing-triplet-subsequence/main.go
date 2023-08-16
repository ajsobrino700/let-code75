package main

import "fmt"

func main(){
  fmt.Println(increasingTriplet([]int{1,2,3,4,5}))
}

func increasingTriplet(nums []int) bool {
  length := len(nums)
  if length<3 {
    return false
  }

  x := nums[0]
  y := nums[1]
  z := nums[2]

  for i:=2;i < length; i++{
    if x<y && y<z {
      return true
    }


  }

  return false
}


func calculate(x *int, y *int, z *int, num int){

  if *x >= *y {
    *x = *y
  }

}
