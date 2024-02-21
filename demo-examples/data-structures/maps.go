package main

import "fmt"

func main() {

	ages := map[string]int{}
	/*
		ages:=make(map[string] int)
	*/

	ages["Optimus"] = 1000
	ages["BumbleBee"] = 500
	ages["Megatron"] = 5000

	fmt.Println(ages)

	//update
	fmt.Println("Updating Bumblebee's age..")
	ages["BumbleBee"] = 800
	fmt.Println(ages)

	//Checking for existence
	fmt.Println("Checking for Megatron's existence and get his age..")
	age, present := ages["Megatron"]
	fmt.Println(present, age)

	//Remove
	fmt.Println("Deleting Megatron..")
	delete(ages, "Megatron")
	fmt.Println(ages)

	//Length of a map
	fmt.Println("Length of ages map - ", len(ages))
}
