//https://www.codechef.com/submit/AGELIMIT

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

    if k < 1 || k > 1000 {
      fmt.Println("No.of tests invalid")
      os.Exit(0)
    }

  	a := make([][]int, k)
  	for j := range a {
  		a[j] = make([]int, 3)
  	}

  	for row := 0; row < k; row++ {
  		for col := 0; col < 3; col++ {
  			fmt.Scanf("%d", &a[row][col])
  		}
  	}

		for row := 0; row < k; row++ {
			if a[row][0] < 20 || a[row][1] > 20 || a[row][0] > 40 || a[row][1] > 40 || a[row][2] < 10 || a[row][2] > 50{
				fmt.Println("Ages invalid")
				os.Exit(0)
			}
		}

    return k, a
}

func main(){
  T, A := readInput()
  for i:=0 ; i < T; i++{
  	if A[i][2] >= A[i][0] && A[i][2]  <  A[i][1]{
			fmt.Println("YES")
			}	else{
				fmt.Println("NO")
			}
  }

}
