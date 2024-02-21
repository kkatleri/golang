package main

import "fmt"

func main() {

	var slice1 = []int{1, 2, 3}
	/*
		OR
		var slice1 = make([]int, 3) - this would initiize empty slice of length and capacity 3
	*/
	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5, 6, 7)

	fmt.Println(slice1)

	//Slicing
	newSlice := slice1[1:3]
	fmt.Println("After slicing...")
	fmt.Println("Original : ", slice1)
	fmt.Println("New Slice : ", newSlice)

	//Copying
	destSlice := make([]int, 4)
	copy(destSlice, slice1)
	fmt.Println("Copied slice - ", destSlice)

	//Iterating over slice
	fmt.Println("Iterating over the slice..")
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

}
