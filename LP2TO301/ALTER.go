//https://www.codechef.com/LP2TO301/problems/ALTER

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


func readInput() (int, [][]int) {
	var k int
	_, err := fmt.Scanf("%d", &k)
	Use(err)

	if k < 1 || k > 1e3 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	a := make([][]int, k)
	for j := range a {
		a[j] = make([]int, 4)
	}

	for row := 0; row < k; row++ {
		for col := 0; col < 4; col++ {
			fmt.Scanf("%d", &a[row][col])
		}
	}

	for i := 0; i < k; i++ {
		for j := 0; j < 4; j++ {
			if a[i][j] < 1 || a[i][j] > 1e9 {
				fmt.Println("Distance not valid")
				os.Exit(0)
			}
		}
	}
	return k, a
}



func main()  {
  tests, array := readInput()
  Use(tests,array)
}
