package talk

import (
	service "/talkservice"

	"./line"
)

//  継承 どうすればええんやｗ
// 勉強し直す

var self = line.Clinet

func SendMessage(to, text string) (Message, error) {
	msg := service.NewMessage()
	msg.To, msg.Text = to, text
	message, err := self.talk.SendMessage(self.ctx, 0, msg)
	return
}

func DeleteOtherFromChat(to string, targetUsers []string) DeleteOtherFromChatResponse {
	req = service.NewDeleteOtherFromChatRequest()
	req.ReqSeq = 0
	req.ChatMid = to
	req.TargetUserMids = targetUsers
	return self.talk.DeleteOtherFromChat(self.ctx, 0, msg)
}

func Noop() {
	self.talk.Noop(self.ctx)
}
