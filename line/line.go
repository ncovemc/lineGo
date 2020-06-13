package line

import (
	"context"

	talk "./talkservice"

	"./login"
)

type Clinet struct {
	talk     *talk.TalkServiceClient
	poll     *talk.TalkServiceClient
	revision int64
	ctx      context.Context
}

func Line(token, appName) line {
	ctx := context.Background()
	talk := login.Talk(token, appName)
	poll := login.Poll(token, appName)
	return Clinet{talk: talk, poll: poll, revision: -1, ctx: ctx}
}
