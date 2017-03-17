// For any 3 elements a,b and c in an array, if the fifth root of (a - b + c) == 0 and given the value of 'a',
// can you find the orhter 2 elements
// https://play.golang.org/p/e6cnpqaIDT
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 16, 19, 13}
	givenValue := 3
	fmt.Println("Array = ", a, ", A = ", givenValue)
	for i := 0; i < len(a); i++ {
		if a[i] == givenValue {
			continue
		}
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] || a[i] == a[j] || a[j] == givenValue || givenValue != (a[i]-a[j]) {
				continue
			}
			fmt.Println("B = ", a[i], ", C = ", a[j])
		}
	}
}
