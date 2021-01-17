package main

import (
  "fmt"
  "errors"
)

type Queue struct{
  Flights []Flight
}

type Flight struct {
  Origin      string
  Destination string
  Price       int
}

func (q *Queue) Pop() (Flight,error) {
  if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	} else {
    firstFlight := q.Flights[0]
		q.Flights = q.Flights[1:]
		return firstFlight, nil
	}
}

func (q *Queue) Push(f Flight) {
  q.Flights = append(q.Flights,f)
}

func (q *Queue) Peek() (Flight,error) {
  if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	} else {
		return q.Flights[0], nil
	}
}

func (q *Queue) IsEmpty() bool {
  return len(q.Flights) == 0
}

func main() {
  flight1 := Flight{Origin: "Pune", Destination: "Nashik", Price: 120,}
  flight2 := Flight{Origin: "Goa", Destination: "Mumbai", Price: 150,}
  
  queue := Queue{}
  // To check all methods when queue is empty
  fmt.Println("queue",queue)
  fmt.Println()
  fmt.Println("queue.IsEmpty()",queue.IsEmpty())
  fmt.Println()
  fmt.Printf("queue.Peek() ")
  fmt.Println(queue.Peek())
  fmt.Println()
  fmt.Printf("queue.Pop() ")
  fmt.Println(queue.Pop())
  fmt.Println()

  // To check all methods when we add flight in the queue
  queue.Push(flight1)
  fmt.Println("queue",queue)
  fmt.Println()
  fmt.Println("queue.IsEmpty()",queue.IsEmpty())
  fmt.Println()
  fmt.Printf("queue.Peek() ")
  fmt.Println(queue.Peek())
  fmt.Println()
  fmt.Printf("queue.Pop() ")
  fmt.Println(queue.Pop())
  fmt.Println()
  fmt.Println("queue",queue)
  fmt.Println()
  fmt.Printf("queue.Pop() ")
  fmt.Println(queue.Pop())
  fmt.Println()

  // To check all methods when we add multiple flight in the queue
  queue.Push(flight1)
  queue.Push(flight2)
  fmt.Println("queue",queue)
  fmt.Println()
  fmt.Println("queue.IsEmpty()",queue.IsEmpty())
  fmt.Println()
  fmt.Printf("queue.Peek() ")
  fmt.Println(queue.Peek())
  fmt.Println()

}
