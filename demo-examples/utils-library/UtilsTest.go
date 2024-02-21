package main

import (
	"fmt"
	"utilsLibrary/utils"
)

func main() {
	add := utils.Add(5, 5)
	fmt.Println("Adding 5 to 5 : ", add)
	sub := utils.Substract(10, 8)
	fmt.Println("Subtracting 8 from 10 : ", sub)

	upperStr := utils.ToUpperCase("Test")
	fmt.Println("Uppercasting Test : ", upperStr)
}
