package main

import (
	. "./line"
	talk "./talkservice"
	"fmt"
	"strconv"
	"time"
)

var client = NewClient("", "MAC")

func main() {
	operation()
}

func operation() {
	for {
		ops, err := client.FetchOperations()
		for _, op := range ops {
			if op.Type != 0 {
				if op.Revision > client.Revision {
					client.Revision = op.Revision
				}
				bot(op)
			}
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func bot(op *talk.Operation) {
	switch op.Type {
	case 26:
		msg := op.Message
		if msg.Text == ".speed" {
			start := time.Now()
			client.SendMessage(msg.To, "start")
			end := time.Now()
			t := strconv.FormatFloat(end.Sub(start).Seconds(), 'f', 8, 64)
			client.SendMessage(msg.To, t)
		}
	}
}
