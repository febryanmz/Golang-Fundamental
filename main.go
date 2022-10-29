package main

import "fmt"

// Assignment no 1
// Get only Odd Number
// Ex:
// input: []int{1,3,4,44,56,51}
// output: []int{1,3,51}
func GetOddNumber(numbers []int) []int {
	// insert your code here

	var arr []int

	// for loop number
	// number%2 --> ganjil/genap nya
	// append number ke array
	// return array nya

	return arr
}

// Assignment no 2
// Given a positive integer A, return an array of strings with all the integers from 1 to N.
// But for multiples of 3 the array should have “Fizz” instead of the number.
// For the multiples of 5, the array should have “Buzz” instead of the number.
// For numbers which are multiple of 3 and 5 both, the array should have “FizzBuzz” instead of the
// Ex:
// input: 15
// output: [1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, 10, 11, Fizz, 13, 14, FizzBuzz]
func FizzBuzz(num int) []int {
	// insert your code here

	return []int{}
}

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())

	fmt.Println(GetOddNumber([]int{1, 3, 4, 44, 56, 51}))
	fmt.Println(FizzBuzz(15))
}

//febryan
