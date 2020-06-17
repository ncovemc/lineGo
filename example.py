package main

import (
	"./line"
)

func main() {
	cl = line.Line("token", "IOS")
	fetchOp(cl)
}

func fetchOp(cl line.Clinet) {
	for {
		ops := cl.fetchOperations(cl.ctx, cl.revision, count)
		for op := range ops {
			bot(op)
			if cl.revision < op.Revision {
				cl.revision = op.Revision
			}
		}
	}
}

func bot(op line.Operation) {
	switch op.Type {
	case 26:
		msg := op.message
		if msg.Text == "test" {
			cl.SendMessage(msg.to, "ok")
		}
	}
}
