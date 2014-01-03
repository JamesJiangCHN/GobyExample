/*
for is Go's only looping construct.Here are three basic
types of for loops.
*/
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	i = 1
	for {
		if i > 3 {
			break
		}
		fmt.Println("loop")
		i += 1
	}

}
