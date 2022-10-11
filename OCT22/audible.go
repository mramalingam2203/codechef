//https://www.codechef.com/submit/AUDIBLE

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

func readInput() (uint32, []uint32) {
	var k uint32
	_, err := fmt.Scanf("%d", &k)
	Use(err)

	if k < 1 || k > 1e4   {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	a := make([]uint32, k)
  var row,i uint32
	for row = 0; row < k; row++ {
			fmt.Scanf("%d", &a[row])
	}

	for i = 0; i < k; i++ {
			if a[i] < 1 || a[i]> 1e6 {
				fmt.Println("Distance not valid")
				os.Exit(0)
		}
	}
	return k, a
}

func main(){
  var i, trials uint32
  frequency := make([]uint32, 0)
  trials, frequency = readInput()
  const low  uint32 = 67
  const high  uint32 = 45e3

  for i =0; i < trials; i++{
    if frequency[i] < low || frequency[i] > high{
      fmt.Println("NO")
    }else{
      fmt.Println("YES")
    }
  }

}
