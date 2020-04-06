package main

import (
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
)

type Client struct {
	// productId是一个engine的唯一标识，每个product都会对应一个engine
	productId string
	// engine持有的处理
	proccess *Send

	// 用于保存回复消息
	storeHandle *mqengine.RabbitMq
	storeChan   chan *mqengine.Message
	ackChan     chan *mqengine.Message
}

func NewClient(productId string, store *mqengine.RabbitMq) *Client {
	e := &Client{
		productId:   productId,
		proccess:    NewSend(productId),
		storeHandle: store,
		storeChan:   make(chan *mqengine.Message, config.ChanNum.Store),
		ackChan:     make(chan *mqengine.Message, config.ChanNum.Ack),
	}
	return e
}

func (e *Client) Start() {
	if err := recover(); err != nil {
		logger.Error("recover", zap.Error(err.(error)))
	}
	e.storeHandle.RegisterProducer(e)
	go e.storeHandle.Start()
	go e.runSend()
	e.proccess.Start()

}

func (e *Client) runSend() {
	for {
		select {
		case message := <-e.proccess.SendChan():
			e.storeChan <- message
		}
	}
}

func (e *Client) Store() (<-chan *mqengine.Message, chan<- *mqengine.Message) {
	return e.storeChan, e.ackChan
}
