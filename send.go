package main

import (
	"github.com/zgfzgf/rabbitmq/mqengine"
)

type Send struct {
	productId string
	sendChan  chan *mqengine.Message
}

func NewSend(productId string) *Send {
	process := &Send{
		productId: productId,
		sendChan:  make(chan *mqengine.Message, config.ChanNum.Store),
	}
	return process
}

func (p *Send) SendChan() <-chan *mqengine.Message {
	return p.sendChan
}

func (p *Send) Start() {
	p.sendChan <- NewMessage("body1", "a1")
	p.sendChan <- NewMessage("body2", "a2")
	p.sendChan <- NewMessage("body3", "a3")
	p.sendChan <- NewMessage("body4", "a4")
	p.sendChan <- NewMessage("body5", "a5")
	p.sendChan <- NewMessage("body6", "a6")
}

func NewMessage(body, corrId string) *mqengine.Message {
	return &mqengine.Message{
		Body:          []byte(body),
		Status:        mqengine.MessageStatusProcess,
		CorrelationId: corrId,
	}
}
