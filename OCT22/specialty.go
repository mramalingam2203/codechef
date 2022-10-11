//https://www.codechef.com/submit/SPECIALITY
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

func readInput() (int, [][]int) {
	var k int
	_, err := fmt.Scanf("%d", &k)
	Use(err)

	if k < 1 || k > 144 {
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

	for i := 0; i < k; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] < 1 || a[i][j] > 100 {
				fmt.Println("Distance not valid")
				os.Exit(0)
			}
		}
	}
	return k, a
}

func max(a int, b int, c int) int{
  if a > b && a > c{
    return a
  }else if c > b {
    return c
  }else{
    return b
  }

}

func  main()  {
  trials, speciality := readInput()
  post := [3]string{"Setter", "Tester", "Editorialist"}
  for i:=0; i<trials; i++{
    most := max(speciality[i][0], speciality[i][1], speciality[i][2])
      fmt.Println(most)
    if most == speciality[i][0]{
      fmt.Println(post[0])
    } else if most == speciality[i][1]{
      fmt.Println(post[1])
    } else {
      fmt.Println(post[2])
    }
  }

}
