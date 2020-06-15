package line

import (
	service "../talkservice"
)

func (self Client) FetchOperations() ([]*service.Operation, error) {
	return self.poll.FetchOperations(self.ctx, self.Revision, 50)
}
