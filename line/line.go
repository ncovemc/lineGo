package line

import (
	"context"
	"log"

	talk "../talkservice"

	"../login"
)

type Client struct {
	talk     *talk.TalkServiceClient
	poll     *talk.TalkServiceClient
	Revision int64
	ctx      context.Context
	Profile  *talk.Profile
}

func NewClient(token, appName string) Client {
	ctx := context.Background()
	talk := login.Talk(token, appName)
	poll := login.Poll(token, appName)
	profile, err := talk.GetProfile(ctx, 0)
  log.Println(profile.DisplayName + ": login success")
	if err != nil {
		log.Fatal(err)
	}
	var rev int64
	for {
		ops, err := poll.FetchOperations(ctx, -1, 50)
		if err != nil {
			log.Fatal(err)
		}
		for _, op := range ops {
			if op.Revision != -1 {
				rev = op.Revision
				break
			}
		}
		if rev != -1 {
			break
		}
	}
	return Client{talk: talk, poll: poll, Revision: rev, ctx: ctx, Profile: profile}
}
