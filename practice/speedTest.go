//https://www.codechef.com/submit/SPEEDTEST

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
  		a[j] = make([]int, 4)
  	}

  	for row := 0; row < k; row++ {
  		for col := 0; col < 4; col++ {
  			fmt.Scanf("%d", &a[row][col])
  		}
  	}

  	if k < 1 || k > 1000 {
  		fmt.Println("No.of tests invalid")
  		os.Exit(0)
  	}

  	for i := 0; i < k; i++ {
  		for j := 0; j < 4; j++ {
  			if a[i][j] < 1 || a[i][j] > 1000 {
  				fmt.Println("Distance not valid")
  				os.Exit(0)
  			}
  		}
  	}

    return k, a
}

func main(){
  T, A := readInput()
  for i:=0 ; i < T; i++{
    if float64(A[i][0]/A[i][1]) <  float64(A[i][2]/A[i][3]){
      fmt.Println("Bob")
    }else if float64(A[i][0]/A[i][1]) >  float64(A[i][2]/A[i][3]){
      fmt.Println("Alice")
    }else{
      fmt.Println("Equal")
    }
  }
}
