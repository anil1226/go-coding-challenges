package challenges

import (
	"fmt"
)

func GetMinMax(flights []Flight) (int, int, error) {
	min := flights[0].Price
	max := flights[0].Price

	for i := 1; i < len(flights); i++ {
		if min > flights[i].Price {
			min = flights[i].Price
		}
		if max < flights[i].Price {
			max = flights[i].Price
		}
	}
	return min, max, nil

}

func GetMinMaxMain() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
	flights := []Flight{
		{Origin: "", Destination: "", Price: 100},
		{Origin: "", Destination: "", Price: 200},
		{Origin: "", Destination: "", Price: 300},
	}
	fmt.Println(GetMinMax(flights))
}
