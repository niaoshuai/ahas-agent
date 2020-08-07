package main

import (
	"ahas-agent/pkg/logger"
	client "github.com/influxdata/influxdb1-client/v2"
)

var influxURL = "http://localhost:8083"

func main() {
	logger.InitLog("ahas-agent.log")
	// 初始化数据库连接
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxURL,
	})
	if err != nil {
		logger.Fatal(err)
	}
	defer c.Close()

	// 开始采集进程数据
	// 开始采集网络数据
	// 开始采集
}