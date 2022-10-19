//https://www.codechef.com/submit/TAXES

package main()

import(
  "os"
  "fmt"
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

    if k < 1 || k > 1e3 {
      fmt.Println("No.of tests invalid")
      os.Exit(0)
    }

  	a := make([][]int, k)
  	for j := range a {
  		a[j] = make([]int, 2)
  	}
