package main

import (
	"fmt"
	"nested_folder/address"
	"nested_folder/name"
)

func main() {
	fmt.Println("Message: This message is come from main folder")
	name.DeveloperName()
	address.DeveloperOfficeAddress()
	address.DeveloperHomeAddress()

}
