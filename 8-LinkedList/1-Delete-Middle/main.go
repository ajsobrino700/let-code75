package main

import "fmt"

func main(){

  linkedList := ListNode{
    Val:1,
    Next : &ListNode{
      Val: 3,
      Next: &ListNode{
        Val: 4,
        Next: &ListNode{
          Val: 7,
          Next: &ListNode{
            Val: 1,
            Next: &ListNode{
              Val: 2,
              Next: &ListNode{
                Val: 6,
                  Next: nil,
              },
            },
          },
        },
      },
    },
  }

  print(deleteMiddle(&linkedList))

  print(deleteMiddle(&ListNode{Val:1}))
}

type ListNode struct{
  Val int
  Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
  size := 0
  current := head
  for current != nil {
    current = current.Next
    size++
  }

  if size == 1 {
    return nil
  }

  removePosition:= size/2
  current = head
  for i:=0; i<removePosition; i++{
     if i == removePosition-1 {
       previous := current
       actual := current.Next
       previous.Next = actual.Next
     }
    current = current.Next
  }

   return head
}

func print(head *ListNode) {
    current := head
    for current != nil {
      fmt.Print(current.Val," ")
      current = current.Next
    }
  fmt.Println()
}
