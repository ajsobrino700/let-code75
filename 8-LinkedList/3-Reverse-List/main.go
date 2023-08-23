package main

import "fmt"

func main(){

  list := ListNode{Val:1}
  list.add(2)
  list.add(3)
  list.add(4)
  list.add(5)

  reverseList(&list).print()

}

type ListNode struct {
  Val int
  Next *ListNode
}

func (t *ListNode) add(num int){
    current:= t
    for current.Next != nil {
        current = current.Next
    }
    current.Next = &ListNode{
      Val: num,
    }
}


func reverseList(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }
  current := head
  var result ListNode = ListNode{ Val: current.Val, Next: nil}
  current = current.Next
  for current != nil {
    aux := result
    result = *current
    result.Next = &aux

    current = current.Next
  }


  return &result
}

func (t *ListNode) print(){
    current := t
    for current != nil {
      fmt.Print(current.Val, " ")
      current = current.Next
    }
    fmt.Println()
}
