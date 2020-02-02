package protocol

import "fmt"

func (rr *RegularReply) HandleRegularReply() error {
	if rr.Error != "" {
		return fmt.Errorf(rr.Error)
	}

	for _, message := range rr.Messages {
		fmt.Println(message)
	}
	return nil
}

func (rr *RegularReply) AddError(err error) *RegularReply {
	rr.Error = err.Error()
	return rr
}

func (rr *RegularReply) AddMessage(msg string) *RegularReply {
	rr.Messages = append(rr.Messages, msg)
	return rr
}
