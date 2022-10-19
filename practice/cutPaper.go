//https://www.codechef.com/submit/CUTPAPER

package main

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

  	for row := 0; row < k; row++ {
  		for col := 0; col < 2; col++ {
  			fmt.Scanf("%d", &a[row][col])
  		}
  	}
/*
		for row := 0; row < k; row++ {
            if 1 <= a[row][0] && a[row][0] <= a[row][1] && a[row][1] <= 1000 {
              fmt.Println("Invalid values!")
              os.Exit(0)
          }
			}
*/

    return k, a
}


func main(){
  //t, a := readInput()
  //fmt.Print(t,a)
  count := 0

  n :=  5
  k := 2
  //for l:=0;l<t; l++{
  for i:=0; i < n; i++{
    for j:=0; j < n ; j++{
      if i%k == 0 && j % k == 0{
        count += 1
        }
      }
    //}
    fmt.Println(count)
  }

}
