package main

import "fmt"

func main(){


  fmt.Println(minFlips(8,3,5))
  fmt.Println(minFlips(2,6,5))
  fmt.Println(minFlips(4,2,7))
  fmt.Println(minFlips(1,2,3))
  fmt.Println(minFlips(238812,179925,927481))
  fmt.Println(minFlips(8280154,4116539,8849163))
}

func minFlips(a int, b int, c int) int {
  bitsC := convertToBit(c)
  bitsB := convertToBit(b)
  bitsA := convertToBit(a)

  length := getMax(len(bitsA),len(bitsB),len(bitsC))

  result:=0
  for i:=0; i<length; i++{
    if ( i < len(bitsC) && bitsC[i] == 1 ) &&  ( ( i >= len(bitsB) || bitsB[i]!=1) && ( i >= len(bitsA) || bitsA[i]!=1) ){
      result++
    }else if  i >= len(bitsC) || bitsC[i] == 0 {
      if i < len(bitsA) && bitsA[i] == 1 {
        result++
      }
      if i< len(bitsB) && bitsB[i] == 1{
        result++
      }
    }
  }

   return result
}

func getMax(a, b, c int) int{
  result := a
  if a<= b && b>=c{
    result = b
  }

  if a>=b && a>=c {
    result = a
  }

  if c>= a && c>=b {
    result = c
  }


  return result
}


func convertToBit(nums int)[]int8{
  var result []int8=[]int8{}
  for nums != 0 {
    result = append(result,int8(nums%2))
    nums/=2
  }
  return result
}
