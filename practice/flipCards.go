//https://www.codechef.com/submit/FLIPCARDS

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
		for col := 0; col < 2; col++ {
			fmt.Scanf("%d", &a[row][col])
		}
	}


	if k < 1 || k > 5000 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

  for i := 0; i < k; i++{
      if a[i][0] < 1 || a[i][0] > 100 || a[i][1] <0 || a[i][1] > a[i][0]{
        fmt.Println("Invalid input")
        os.Exit(0)
      }
  }

  for i := 0; i < k; i++ {
    if a[i][0] == a[i][1] || a[i][1] == 0{
        fmt.Println("0" )
      }else if a[i][0] <= 2*a[i][1]{
        fmt.Println(a[i][0]-a[i][1])
      }else if a[i][0] > 2*a[i][1]{
        fmt.Println(a[i][1])
      }
    }

  return
}
