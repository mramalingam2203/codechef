//https://www.codechef.com/submit/EXPERT
//
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


		for row := 0; row < k; row++ {
            if !(a[row][0] >= 1 || a[row][0] <= a[row][1] || a[row][1] <= 1e6) {
              fmt.Println("Invalid values!")
              os.Exit(0)
          }
			}


    return k, a
}


func main()  {
  t,a := readInput()
  fmt.Print(t,a)
  var half float64 = 0.5
  for i:=0; i< t; i++{
    if float64(a[i][1]) >= half*float64(a[i][0]){
      fmt.Println("YES")
    }else{
      fmt.Println("NO")
    }

  }
}
