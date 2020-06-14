package line

import (
	service "../talkservice"
)

func (self Client) SendMessage(to, text string) (service.Message, error) {
	msg := service.NewMessage()
	msg.To, msg.Text = to, text
	message, err := self.talk.SendMessage(self.ctx, 0, msg)
	return *message, err
}

func (self Client) DeleteOtherFromChat(to string, targetUsers []string) (service.DeleteOtherFromChatResponse, error) {
	req := service.NewDeleteOtherFromChatRequest()
	req.ReqSeq = 0
	req.ChatMid = to
	req.TargetUserMids = targetUsers
	result, err := self.talk.DeleteOtherFromChat(self.ctx, req)
	return *result, err
}

func (self Client) GetProfile() (*service.Profile, error) {
	prof, err := self.talk.GetProfile(self.ctx, 0)
	self.Profile = prof
	return prof, err
}

func (self Client) Noop() {
	self.talk.Noop(self.ctx)
}

//
