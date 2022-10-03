//https://www.codechef.com/LP2TO301/problems/MAKEDIV3

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func readInput() (int, []int) {
	var k int
	_, err := fmt.Scanf("%d", &k)
	Use(err)

	if k < 1 || k > 500 {
		fmt.Println("No.of tests invalid")
		os.Exit(0)
	}

	a := make([]int, k)

	for row := 0; row < k; row++ {
		fmt.Scanf("%d", &a[row])
	}

	for row := 0; row < k; row++ {
		if a[row] < 1 || a[row] > 1e4 {
			fmt.Println("Huge numbers")
			os.Exit(0)
		}
	}

	sum := 0
	for t := 0; t < k; t++ {
		sum += a[t]
	}

	if sum > 1e5 {
		fmt.Println("Summing criterion not met")
		os.Exit(0)
	}

	fmt.Println(sum)

	return k, a
}

func filterNumber(s string) (bool, []rune) {
	r := []rune(s)
	const rune_offset int32 = 48
	for i := 0; i < len(r); i++ {
		r[i] -= rune_offset
	}

	var sumResult int32
	for i := 0; i < len(r); i++ {
		sumResult += r[i]
	}

	if sumResult%3 == 2 {
		r[len(r)-1] += 1
		sumResult += 1
	} else if sumResult%3 == 1 {
		r[len(r)-1] -= 1
		sumResult -= 1
	}

	if sumResult%3 == 0 && sumResult%9 != 0 {
		return true, r
	} else {
		return false, r
	}

}

const charset = "123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	t, a := readInput()
	var number string

	Use(a, t, number)

	fmt.Println(t)
	for i := 0; i < t; i++ {
		var decide bool = false
		var nrune []rune
		Use(nrune)
		for !decide {
			number = StringWithCharset(a[i], charset)
			decide, nrune = filterNumber(number)
		}
		for j := 0; j < len(nrune); j++ {
			fmt.Print(nrune[j])
		}
		fmt.Println()
	}

	return
}
