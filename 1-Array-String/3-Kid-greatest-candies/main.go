package main

import "fmt"

func main(){
  fmt.Println(kidsWithCandies([]int{2,3,5,1,3},3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
  maxCandies := 0
  kidNumber := len(candies)

  for i := 0; i < kidNumber; i++ {
    if candies[i] > maxCandies  {
      maxCandies = candies[i]
    }
  }

  result := []bool{}


  for i := 0; i < kidNumber; i++{
    if maxCandies <= candies[i]+extraCandies {
      result = append(result, true)
    }else {
      result =append(result, false)
    }
  }

  return result
}
