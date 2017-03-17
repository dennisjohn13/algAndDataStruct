// For any 3 elements a,b and c in an array, if the fifth root of (a - b + c) == 0 and given the value of 'a',
// can you find the orhter 2 elements
package main

import "fmt"
import "math"

func main() {
	var z float64
	var y float64
	z = 0
	for n := 0; n < 5; n++ {
		fmt.Println(y, z)
		y = math.Sqrt(z)
		z = y
		fmt.Println(y, z)
	}
	fmt.Println(math.Sqrt(1.4953487812212205))
}
