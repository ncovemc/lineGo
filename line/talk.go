package line

import (
	service "../talkservice"
)

func (cl Client) SendMessage(to, text string) (service.Message, error) {
	msg := service.NewMessage()
	msg.To, msg.Text = to, text
	message, err := cl.talk.SendMessage(cl.ctx, 0, msg)
	return *message, err
}

func (cl Client) DeleteOtherFromChat(to string, targetUsers []string) (service.DeleteOtherFromChatResponse, error) {
	req := service.NewDeleteOtherFromChatRequest()
	req.ReqSeq = 0
	req.ChatMid = to
	req.TargetUserMids = targetUsers
	result, err := cl.talk.DeleteOtherFromChat(cl.ctx, req)
	return *result, err
}

func (cl Client) GetProfile() (*service.Profile, error) {
	return cl.talk.GetProfile(cl.ctx, 0)
}

func (cl Client) Noop() {
	cl.talk.Noop(cl.ctx)
}

//
