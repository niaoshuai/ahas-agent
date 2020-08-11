package main

import (
	"ahas-agent/pkg/logger"
	"ahas-agent/pkg/proc"
	client "github.com/influxdata/influxdb1-client/v2"
	"os"
	"time"
)

var influxURL = "http://localhost:8086"

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

	for {
		// 开始采集进程数据
		go reportProcess(c)

		time.Sleep(10 * time.Second)
	}
}

func reportProcess(c client.Client) {
	hostName, err := os.Hostname()
	if err != nil {
		logger.Warn(err)
	}
	// 批量插入配置
	batchPoint, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "ahas",
	})
	if err != nil {
		logger.Warn(err)
	}
	// 获取 进程 信息
	processData, connectData := proc.GetProcessData()
	for i := processData.Front(); i != nil; i = i.Next() {
		// 上报数据
		ps := i.Value.(proc.PData)
		tags := map[string]string{"host": hostName, "pid": string(ps.Pid)}
		fields := map[string]interface{}{
			"ppid": ps.Ppid,
			"path": ps.Path,
			"exec": ps.Exec,
			"cpu":  ps.Cpu,
			"mem":  ps.Mem,
		}
		// 时序 ahas_processes
		pt, err := client.NewPoint("ahas_processes", tags, fields, time.Now())
		if err != nil {
			logger.Warn(err)
		}
		batchPoint.AddPoint(pt)
	}

	for i := connectData.Front(); i != nil; i = i.Next() {
		connect := i.Value.(proc.CData)
		tags := map[string]string{"host": hostName, "pid": string(connect.Pid)}
		fields := map[string]interface{}{
			"localAddr":  connect.LocalAddr,
			"remoteAddr": connect.RemoteAddr,
			"status":     connect.Status,
		}
		// 时序 ahas_processes
		pt, err := client.NewPoint("ahas_netstat", tags, fields, time.Now())
		if err != nil {
			logger.Warn(err)
		}
		batchPoint.AddPoint(pt)

	}

	err = c.Write(batchPoint)
	if err != nil {
		logger.Warn(err)
	}
	logger.Info("data record")
}
