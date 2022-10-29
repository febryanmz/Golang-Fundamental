package main

import "fmt"

func main() {
	//INTERGER
	fmt.Println("this is integer", 1)
	fmt.Println("this is floating point", 1.2)

	var angka int
	angka = -129

	fmt.Println(angka)

	//BOOLEAN
	fmt.Println(true)
	fmt.Println(false)

	//STRING
	fmt.Println(`this is back quote string`)
	fmt.Println("this is double quote string")

	// [el1, el2, el3]
	// el1 -> 0
	// el2 -> 1
	// el3 -> 2

	// built in function in string type
	fmt.Println(len("hello"))       // count characters
	fmt.Println(string("hello"[1])) // get characters using index position, will return raw byte

}

//febryan
