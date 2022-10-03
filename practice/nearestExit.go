//https: //www.codechef.com/submit/NEARESTEXI

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

	if T < 1 || T > 100 {
		fmt.Println("No. of trials invalid")
		os.Exit(0)
	}

	seat := make([]int, T)

	for row := 0; row < T; row++ {
		fmt.Scanf("%d", &seat[row])
	}

	for i := 0; i < T; i++ {
		if seat[i] < 0 || seat[i] > 100 {
			fmt.Println("Seat no. invalid")
			os.Exit(0)
		}
	}

	for i := 0; i < T; i++ {
		if seat[i] <= 50 {
			fmt.Println("LEFT")
		} else {
			fmt.Println("RIGHT")
		}
	}

	return
}
