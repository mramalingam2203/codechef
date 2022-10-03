//https://www.codechef.com/submit/BIRYANI


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

func readInput()(int, [][]int){
  	var k int
  	_, err := fmt.Scanf("%d", &k)
  	Use(err)

  	a := make([][]int, k)
  	for j := range a {
  		a[j] = make([]int,2)
  	}

  	for row := 0; row < k; row++ {
  		for col := 0; col < 2; col++ {
  			fmt.Scanf("%d", &a[row][col])
  		}
  	}

  	if k < 1 || k > 1e4 {
  		fmt.Println("No.of tests invalid")
  		os.Exit(0)
  	}

  	for i := 0; i < k; i++ {
  		for j := 0; j < 2; j++ {
  			if a[i][j] < 1 || a[i][j] > 100 {
  				fmt.Println("Distance not valid")
  				os.Exit(0)
  			}
  		}
  	}

    return k, a
}

func  main()  {
  T, A := readInput()
  for i:=0; i < T; i++{
    fmt.Println(A[i][0]*A[i][1])
  }


}
