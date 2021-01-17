package main

import "fmt"

type Element struct {
  Value string
  Next  *Element
}

type SinglyLinkedList struct {
  Count int
  Head  *Element
  Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	// create a new ode
	node := Element{value, nil}

	// Condition for the first node
	if l.Head == nil {
		l.Head = &node
		l.Tail = &node
		l.Count++
		return
	}
	// Iterate to go to last node
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	// reached the last node
	current.Next = &node
	// update the node count
	l.Count++
	// upate the tail node
	l.Tail = &node
}

// You will have to ensure when you add new elements
// that this method still returns the correct value
func (l *SinglyLinkedList) Size() int {
  return l.Count
}

func (l *SinglyLinkedList) Print() {
  current := l.Head
  fmt.Printf("%+v\n", current.Value)
  for current.Next != nil {
    fmt.Printf("%+v\n", current.Value)
    current = current.Next
  }
}

func main() {
  fmt.Println("Singly Linked List Challenge")

  var llist SinglyLinkedList

  values := []string{"First", "Second", "Third"}
  for _, value := range values {
    llist.Append(value)
  }
  llist.Print()
}
