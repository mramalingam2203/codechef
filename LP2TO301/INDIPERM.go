//https://www.codechef.com/LP2TO301/problems/INDIPERM

package main

import "fmt"

//function to swap the variables
func swap(a *int32, b *int32) {
	var temp int32
	temp = *a
	*a = *b
	*b = temp
}

//function to print the array
func printarray(arr []int32) {
	//  if isAbsolutePermut(arr,2) == true{
	for i := 1; i < len(arr); i++ {
		//if arr[i]%int32(i+1) != 0 {
		fmt.Print(arr[i])
		//}
	}
	fmt.Println()
	//  }
}

func permutation(arr []int32, start int32, end int32) {
	if start == end {
		printarray(arr)
	}
	var i int32
	for i = start; i <= end; i++ {
		swap(&arr[i], &arr[start])
		permutation(arr, start+1, end)
		swap(&arr[i], &arr[start])
	}
}

func main() {
	array := []int32{1, 2, 3, 4}
	permutation(array, 0, int32(len(array)-1))

}
