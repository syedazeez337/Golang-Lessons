package main

import (
	"fmt"
	"maps"
)

func main() {
	var aMap map[string]int

	fmt.Println(aMap, len(aMap))

	names := map[string]int{
		// map[keyType]valueType
		"Hello": 1,
		"world": 2,
	}
	names["world"]++
	fmt.Println("value", names["Hello"])
	fmt.Println(names["world"])

	v, ok := names["Hello"]
	fmt.Println(v, ok)

	delete(names, "Hello")

	fmt.Println(names)

	clear(names)
	fmt.Println(names)

	m := map[string]int{
		"hello": 5,
		"world": 6,
	}

	n := map[string]int{
		"world": 6,
		"hello": 5,
	}

	fmt.Println(maps.Equal(n, m))

	/*
		for key, val := range names {
			fmt.Println(key, val)
		}

		s := make(map[string]string, 10)
		fmt.Println(s, len(s))
	*/
}
