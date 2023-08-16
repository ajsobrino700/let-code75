package main

import "fmt"

func main(){

  fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))

}

func maxArea(height []int) int {
  maxArea := 0
  for i:=0; i<len(height); i++ {
    for j:=i+1; j<len(height); j++{
      area := (j-i)*getMin(height[i],height[j])
      if area > maxArea {
        maxArea = area
      }
    }
  }
  return maxArea
}


func getMin(x, y int) int {
  result:= y
  if x<y {
    result = x
  }
  return result
}
