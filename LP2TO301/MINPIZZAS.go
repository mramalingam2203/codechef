//https://www.codechef.com/LP2TO301/problems/MINPIZZAS

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

	if k < 1 || k > 2e5 {
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
			if a[i][j] < 1 || a[i][j] > 1e9 {
				fmt.Println("Distance not valid")
				os.Exit(0)
			}
		}
	}
	return k, a
}

func buyPizza(t int, a [][]int) int {
	for i := 0; i < t; i++ {
		if a[i][0] == a[i][1] {
			fmt.Println(1)
		} else if a[i][0]%a[i][1] != 0 {
			fmt.Println(a[i][0] % a[i][1])

		} else if a[i][0]%a[i][1] == 0 {
			fmt.Println(a[i][0] / a[i][1])
		}
	}
	return 0
}

func main() {
	T, A := readInput()
	Use(T, A)
	buyPizza(T, A)

}
