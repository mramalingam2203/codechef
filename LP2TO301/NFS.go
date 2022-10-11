//https://www.codechef.com/LP2TO301/problems/NFS

package main

import (
	"fmt"
	"math"
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

	if k < 1 || k > 1e5 {
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

	return k, a
}

func speed(t int, a [][]int) {
	//u = sqrt(v*v-2*a*s)
	//order -- u,v,a,s
	fmt.Println(t)

	for i := 0; i < t; i++ {
		//exp_1 := math.Sqrt(
		exp_1 := float64(a[i][1]*a[i][1]-2*a[i][2]*a[i][3])
		exp_2 := math.Sqrt(float64(a[i][0]))

		fmt.Println(exp_1, exp_2)

		if exp_1 == exp_2 {
			fmt.Println("YES")
		}

		if exp_1 <= exp_2 ||  exp_1 == exp_2 {
			fmt.Println("YES")
		} else if exp_1 > exp_2 {
 			fmt.Println("NO")
		}
	}
}

func main() {
	T, A := readInput()
	//Use(T, A)
	speed(T, A)

}
