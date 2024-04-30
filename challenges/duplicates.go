package challenges

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	// TODO Implement
	devMap := make(map[string]bool)
	res := make([]string, 0)
	for _, v := range developers {
		if devMap[v.Name] {
			continue
		}
		devMap[v.Name] = true
		res = append(res, v.Name)
	}
	return res
}

func FilterUniqueMain() {
	fmt.Println("Filter Unique Challenge")
	devs := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}
	fmt.Println(FilterUnique(devs))
}
