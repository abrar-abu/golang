package main

import (
  "fmt"
  "errors"
)

type Stack struct{
  Flights []Flight
}

type Flight struct {
  Origin      string
  Destination string
  Price       int
}

func (s *Stack) Pop() (Flight,error) {
  if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		// lastElemIndex := len(s.Items) - 1
		flight := s.Flights[len(s.Flights)-1]
		s.Flights = s.Flights[:len(s.Flights)-1]
		return flight, nil
	}
}

func (s *Stack) Push(f Flight) {
  s.Flights = append(s.Flights,f)
}

func (s *Stack) Peek() (Flight,error) {
  if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		return s.Flights[len(s.Flights)-1], nil
	}
}

func (s *Stack) IsEmpty() bool {
  return len(s.Flights) == 0
}

func main() {
  flight1 := Flight{Origin: "Pune", Destination: "Nashik", Price: 120,}
  flight2 := Flight{Origin: "Goa", Destination: "Mumbai", Price: 150,}
  
  stack := Stack{}
  // To check all methods when stack is empty
  fmt.Println("stack",stack)
  fmt.Println()
  fmt.Println("stack.IsEmpty()",stack.IsEmpty())
  fmt.Println()
  fmt.Printf("stack.Peek() ")
  fmt.Println(stack.Peek())
  fmt.Println()
  fmt.Printf("stack.Pop() ")
  fmt.Println(stack.Pop())
  fmt.Println()

  // To check all methods when we add flight in the stack
  stack.Push(flight1)
  fmt.Println("stack",stack)
  fmt.Println()
  fmt.Println("stack.IsEmpty()",stack.IsEmpty())
  fmt.Println()
  fmt.Printf("stack.Peek() ")
  fmt.Println(stack.Peek())
  fmt.Println()
  fmt.Printf("stack.Pop() ")
  fmt.Println(stack.Pop())
  fmt.Println()
  fmt.Println("stack",stack)
  fmt.Println()
  fmt.Printf("stack.Pop() ")
  fmt.Println(stack.Pop())
  fmt.Println()

  // To check all methods when we add multiple flight in the stack
  stack.Push(flight1)
  stack.Push(flight2)
  fmt.Println("stack",stack)
  fmt.Println()
  fmt.Println("stack.IsEmpty()",stack.IsEmpty())
  fmt.Println()
  fmt.Printf("stack.Peek() ")
  fmt.Println(stack.Peek())
  fmt.Println()
}
