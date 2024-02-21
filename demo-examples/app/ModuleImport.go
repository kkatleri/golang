package main

import (
	"fmt"
	"utilsLibrary/utils"
)

func main() {
	fmt.Println("Adding 5, 8 : ", utils.Add(5, 8))
	fmt.Println("autobots to Uppoer Case : ", utils.ToUpperCase("autobots"))
	fmt.Println("Bumblebee to Upper Case using Alias method : ", utils.AliasToUpper("Bumblebee"))
}
