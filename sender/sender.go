package sender

import (
	"fmt"
)

type Sender struct {
	ch       chan Mail
	doneChan chan struct{}
}

type Mail struct {
	Email     string
	TypeOfMsg string
	Values    map[string]string
}

func NewSender() *Sender {
	s := Sender{
		ch:       make(chan Mail),
		doneChan: make(chan struct{}),
	}

	go s.bg()

	return &s
}

func (s Sender) Send(login string, typeOfMsg string, values map[string]string) error {
	s.ch <- Mail{
		Email:     login,
		TypeOfMsg: typeOfMsg,
		Values:    values,
	}
	return nil
}

func (s Sender) bg() {
	for {
		select {
		case <-s.doneChan:
			return
		case m := <-s.ch:
			fmt.Printf("send %#v\n", m)
			// TODO: send email (go)
		}
	}
}
