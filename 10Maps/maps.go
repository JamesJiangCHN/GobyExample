/*
Maps are Go's built-in associative data type (sometimes
called hashes or dicts in other languages).
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	v2, prs := m["k2"]
	fmt.Println("prs:", v2, prs)

	n := map[string]int{"n1": 1, "n2": 2}
	fmt.Println("map:", n)
}
