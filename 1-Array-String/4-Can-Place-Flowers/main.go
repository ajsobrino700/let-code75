package main

import "fmt"

func main(){
  fmt.Println(canPlaceFlowers([]int{1,0,0,0,1},1))
}


func canPlaceFlowers(flowerbed []int, n int) bool {
  length := len(flowerbed)-1
  previous, next := 0,0

  for i:=0; i<length; i++{
    next = flowerbed[i+1]
    if previous == 0 && flowerbed[i] == 0 && next == 0 {
      flowerbed[i] = 1
      n--
    }
    previous = flowerbed[i]
  }

  if previous == 0 && flowerbed[length] == 0 {
   n--
  }

  return n <= 0
}
