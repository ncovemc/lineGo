package line

import (
	service "../talkservice"
)

func (self Client) FetchOperations() ([]*service.Operation, error) {
	ops, err := self.poll.FetchOperations(self.ctx, self.Revision, 50)
	return ops, err
}
