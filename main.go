package main

import (
	."./line"
	"fmt"
)
var client = NewClient("", "MAC")
func main() {
	fmt.Println(client.Revision)
	operation()
}

func operation() {
	ops, err := client.FetchOperations()
	fmt.Println(ops)
	if err != nil {
		fmt.Println(err)
	}
}
