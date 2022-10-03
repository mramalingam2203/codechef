//https://www.codechef.com/submit/BOBBANK

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

	var k int
	_, err := fmt.Scanf("%d", &k)
	//	fmt.Println(k, err)
	Use(err)

	a := make([][]int, k)
	for j := range a {
		a[j] = make([]int, 4)
	}

	for row := 0; row < k; row++ {
		for col := 0; col < 4; col++ {
			fmt.Scanf("%d", &a[row][col])
		}
	}

	fmt.Println("Printing 2D Arrays", a)

	if k < 1 || k > 1000 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	for i := 0; i < k; i++ {
		for j := 0; j < 4; j++ {
			if a[i][j] < 0 || a[i][j] > 1e4 {
				fmt.Println("Amounts invalid")
				os.Exit(0)
			}
		}
	}

  for i := 0; i < k; i++ {
    fmt.Println( a[i][0] + (a[i][1] - a[i][2]) * a[i][3])
  }
  return
}
