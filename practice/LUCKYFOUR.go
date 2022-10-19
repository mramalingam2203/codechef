//https://www.codechef.com/submit/LUCKYFR

package main

import(
  "fmt"
  "os"
  "strconv"
  "strings"
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

    if k < 1 || k > 1e5 {
      fmt.Println("No.of tests invalid")
      os.Exit(0)
    }

  	a := make([]int, k)

  	for row := 0; row < k; row++ {
  			fmt.Scanf("%d", &a[row])
  		}

    return k,a
}

func ItoA(n int) string{
    s := strconv.Itoa(n)
    return s
}

func main(){
  t, a := readInput()
  for i:=0; i <t ;i++{
    int_to_string := ItoA(a[i])
    fmt.Println(strings.Count(int_to_string, "4"))
  }
}
