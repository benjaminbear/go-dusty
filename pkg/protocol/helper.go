package protocol

import "fmt"

func (rr *RegularReply) HandleRegularReply() error {
	if rr.Error {
		return fmt.Errorf(rr.Message)
	}

	fmt.Println(rr.Message)
	return nil
}

func CreateRegularErrorReply(err error) *RegularReply {
	return &RegularReply{Error: true, Message: err.Error()}
}
