//https://www.codechef.com/submit/FIT

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

func readInput()(int, []int){
  	var k int
  	_, err := fmt.Scanf("%d", &k)
  	Use(err)

  	a := make([]int, k)


  	for row := 0; row < k; row++ {
  			fmt.Scanf("%d", &a[row])
  	}

  	if k < 1 || k > 10 {
  		fmt.Println("No.of tests invalid")
  		os.Exit(0)
  	}

  	for i := 0; i < k; i++ {
  			if a[i]< 1 || a[i] > 10 {
  				fmt.Println("Distance not valid")
  				os.Exit(0)
  		}
  	}

    return k, a
}

func  main()  {
  T, A := readInput()
  for i:=0; i < T; i++{
    fmt.Println(10*A[i])
  }


}
