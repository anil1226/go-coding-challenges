package challenges

import "fmt"

type Queue struct {
	flights []Flight
}

func (s *Queue) Pop() Flight {
	if s.IsEmpty() {
		return Flight{}
	}
	firstFlight := s.flights[len(s.flights)-1]
	s.flights = s.flights[:len(s.flights)-1]
	return firstFlight
}

func (s *Queue) Push(f Flight) {
	s.flights = append(s.flights, []Flight{f}...)
}

func (s *Queue) Peek() Flight {
	if s.IsEmpty() {
		return Flight{}
	}
	return s.flights[len(s.flights)-1]
}

func (s *Queue) IsEmpty() bool {
	return len(s.flights) == 0
}

func QueueMain() {
	fmt.Println("Go Queue Implementation")
	s := []Flight{
		{Origin: "", Destination: "", Price: 100},
		{Origin: "", Destination: "", Price: 200},
		{Origin: "", Destination: "", Price: 300},
	}

	f := Flight{
		Origin: "", Destination: "", Price: 500,
	}

	s2 := Queue{
		flights: s,
	}

	fmt.Println("Current Queue", s2)
	s2.Push(f)
	fmt.Println("Pushed a new Flight. Current Queue", s2)
	fmt.Println("Peek the Top of Queue", s2.Peek())
	fmt.Println("Pop the Queue", s2.Pop())
	fmt.Println("Final Queue", s2)
}
