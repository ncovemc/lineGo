package main

import (
	"fmt"

	. "./line"
)

var client = NewClient("ESOapJ4bSgBsdu2lRPj6.P0E/6h9KZjye3300Y/LnHG.LyfM0smymGNxs+6K7WJhsmwSR+0Ilg9wvKy83EnW0Rs=", "MAC")

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
