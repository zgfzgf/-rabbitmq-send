package main

import "github.com/zgfzgf/rabbitmq/mqengine"

func StartClient() {
	productId := "aaa"
	storeMq := mqengine.NewStoreMQ(productId)
	send := NewClient(productId, storeMq)
	send.Start()
}
