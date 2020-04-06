package main

import (
	"encoding/json"
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
	"os"
)

var config *mqengine.GbeConfig
var logger *zap.Logger

func main() {
	config = mqengine.GetConfig("./send.json")
	byte, _ := json.Marshal(config)
	logger = mqengine.GetLog()
	defer logger.Sync()
	logger.Info("log 初始化成功")
	logger.Info("see:",
		zap.ByteString("conf", byte))
	StartClient()
	//for i:=0; i<1000000; i++{
	//	logger.Info("log 初始化成功")
	//}
	c := make(chan os.Signal)
	<-c
}
