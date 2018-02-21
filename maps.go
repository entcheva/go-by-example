package main

import "fmt"

func main() {
	
	// maps aka hashes
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

}
