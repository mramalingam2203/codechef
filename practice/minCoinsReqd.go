//https: //www.codechef.com/submit/MINCOINSREQ

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

	if T < 1 || T > 1000 {
		fmt.Println("No. of trials invalid")
		os.Exit(0)
	}

	money := make([]int, T)

	for row := 0; row < T; row++ {
		fmt.Scanf("%d", &money[row])
	}

	for i := 0; i < T; i++ {
		if seat[i] < 0 || money[i] > 1000 {
			fmt.Println("money invalid")
			os.Exit(0)
		}
	}
