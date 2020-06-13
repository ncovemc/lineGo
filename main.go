package main

import (
	."./line"
	"fmt"
)

func main() {
	client := NewClient("", "MAC")
	profile, err := client.GetProfile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile)
}
