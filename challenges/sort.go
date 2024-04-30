package challenges

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Flights []Flight

func (f Flights) Len() int {
	return len(f)
}

func (f Flights) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Flights) Less(i, j int) bool {
	return f[i].Price < f[j].Price
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	//sort.Sort(flights)
	sort.Sort(Flights(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func SortByPriceMain() {

	flights := []Flight{
		{
			Origin:      "Hyd",
			Destination: "BLR",
			Price:       200,
		},
		{
			Origin:      "Hyd",
			Destination: "BZA",
			Price:       100,
		},
	}
	//var flights []Flight

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
