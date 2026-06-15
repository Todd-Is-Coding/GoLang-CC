package main

import (
	"fmt"
	"maps"
)


func main(){
	m := make(map[string] int)

	m["k1"] = 5 
	m["k2"] = 7


	fmt.Println("map:", m)

	v1 := m["k1"]

	fmt.Println(v1)

	v3 := m["k3"]

	fmt.Println(v3)

	fmt.Println("len:", len(m))

	delete(m , "k2")
	fmt.Println("map:", m)
	clear(m)
    fmt.Println("map:", m)

	/*
		1. Map lookup can return two values

		When you access a map in Go:

		value, exists := m["k2"]
		value → the value stored at key "k2"
		exists (often called ok or prs) → true if the key exists, false otherwise

		Example:

		m := map[string]int{
			"k1": 7,
			"k2": 13,
		}

		v, prs := m["k2"]

		fmt.Println(v)   // 13
		fmt.Println(prs) // true
		2. The _ (blank identifier)

		The underscore means:

		"I don't care about this value, throw it away."

		So:

		_, prs := m["k2"]

		means:

		value, prs := m["k2"] // get both values
		_ = value             // discard value

		You're only interested in whether the key exists.

		Example
		m := map[string]int{
			"k1": 7,
			"k2": 13,
		}

		_, prs := m["k2"]
		fmt.Println(prs)

		Output:

		true

		Because "k2" exists.

		Missing key
		_, prs := m["k3"]
		fmt.Println(prs)

		Output:

		false

		Because "k3" is not in the map.
			*/
	_ , prs := m["k2"]

	fmt.Println("prs" , prs)

	n := map[string]int {"ahmed" : 1 , "arafa" : 2}

	n2 := map[string]int {"ahmed" : 1 , "arafa" : 2}


	if maps.Equal(n , n2){
		fmt.Println("they are equal")
	}
}