package line

import (
	"context"
	talk "../talkservice"
	"../login"
)

type Client struct {
	talk     *talk.TalkServiceClient
	poll     *talk.TalkServiceClient
	revision int64
	ctx      context.Context
}

func Line(token, appName string) Client {
	ctx := context.Background()
	talk := login.Talk(token, appName)
	poll := login.Poll(token, appName)
	return Client{talk: talk, poll: poll, revision: -1, ctx: ctx}
}
