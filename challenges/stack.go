package challenges

import "fmt"

type Stack struct {
	flights []Flight
}

func (s *Stack) Pop() Flight {
	if s.IsEmpty() {
		return Flight{}
	}
	firstFlight := s.flights[0]
	s.flights = s.flights[1:]
	return firstFlight
}

func (s *Stack) Push(f Flight) {
	s.flights = append([]Flight{f}, s.flights...)
}

func (s *Stack) Peek() Flight {
	if s.IsEmpty() {
		return Flight{}
	}
	return s.flights[0]
}

func (s *Stack) IsEmpty() bool {
	return len(s.flights) == 0
}

func StackMain() {
	fmt.Println("Go Stack Implementation")
	s := []Flight{
		{Origin: "", Destination: "", Price: 100},
		{Origin: "", Destination: "", Price: 200},
		{Origin: "", Destination: "", Price: 300},
	}

	f := Flight{
		Origin: "", Destination: "", Price: 500,
	}

	s2 := Stack{
		flights: s,
	}

	fmt.Println("Current Stack", s2)
	s2.Push(f)
	fmt.Println("Pushed a new Flight. Current Stack", s2)
	fmt.Println("Peek the Top of Stack", s2.Peek())
	fmt.Println("Pop the Stack", s2.Pop())
	fmt.Println("Final Stack", s2)
}
