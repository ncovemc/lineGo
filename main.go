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
	profile, _ := client.GetProfile()
	fmt.Println(profile.DisplayName + ": login success")
	for {
		ops, err := client.FetchOperations()
		for _, op := range ops {
			if op.Type != 0 {
				if op.Revision > client.Revision {
					client.Revision = op.Revision
				}
				fmt.Println(op)
			}
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
