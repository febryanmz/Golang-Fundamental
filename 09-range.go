package main

import "fmt"

func main() {

	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for index, value := range nums {
		sum += value
		fmt.Println(index, value)
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself. See Strings and Runes for more details.
	for i, c := range "go" {
		fmt.Println(i, string(c))
	}

	fmt.Println("--------------------")
	for _, v := range []int{1, 2, 3} {
		if v%2 == 0 {
			fmt.Println("genap")
		} else {
			fmt.Println("ganjil")
		}

	}
}

//febryan
