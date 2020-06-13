package main

import (
	."./line"
)

func main() {
	client := NewClient("", "MAC")
        client.Noop()
}
