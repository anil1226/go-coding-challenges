package goroutines

import "fmt"

func MatrixProc() {
	mat1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	mat2 := [3][3]int{
		{11, 12, 13},
		{14, 15, 16},
		{17, 18, 19},
	}

	fmt.Println(len(mat1), mat2)

	// for _, v := range mat1 {
	// 	for i := 0; i < len(mat2); i++ {
	// 		v[0] * mat2[][]
	// 	}
	// }
}
