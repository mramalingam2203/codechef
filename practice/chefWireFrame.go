//https: //www.codechef.com/submit/CWIREFRAME

package main

import (
	"fmt"
	"os"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {

	var T int
	_, err := fmt.Scanf("%d", &T)
	//	fmt.Println(k, err)
	Use(err)

	a := make([][]int, T)
	for j := range a {
		a[j] = make([]int, 3)
	}

	for row := 0; row < T; row++ {
		for col := 0; col < 3; col++ {
			fmt.Scanf("%d", &a[row][col])
		}
	}

	if T < 1 || T > 1000 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	for i := 0; i < T; i++ {
		fmt.Println(2 * (a[i][0] + a[i][1]) * a[i][2])
	}
}
