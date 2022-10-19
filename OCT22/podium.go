//https://www.codechef.com/submit/PODIUM

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

	if k < 1 || k > 100 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	a := make([][]int, k)
	for j := range a {
		a[j] = make([]int, 2)
	}

	for row := 0; row < k; row++ {
		for col := 0; col < 2; col++ {
			fmt.Scanf("%d", &a[row][col])
		}
	}

	for i := 0; i < k; i++ {
		for j := 0; j < 2; j++ {
			if a[i][j] < 1 || a[i][j] > 10 {
				fmt.Println("Distance not valid")
				os.Exit(0)
			}
		}
	}
	return k, a
}

func main(){
	trials, time := readInput()
	for i:=0; i < trials; i++{
		fmt.Println(time[i][0] + time[i][1])
	}
}
