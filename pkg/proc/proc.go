/**
 * @Author: niaoshuai
 * @Date: 2020/8/10 3:47 下午
 */
package proc

import (
	"ahas-agent/pkg/logger"
	"container/list"
	"github.com/shirou/gopsutil/process"
	"strings"
)

type PData struct {
	Pid  int32
	Ppid int32
	Path string
	Exec string
	Cpu  float64
	Mem  float32
}

type CData struct {
	Pid        int32
	LocalAddr  string
	RemoteAddr string
	Status     string
}

// 获取进程信息
func GetProcessData() (*list.List, *list.List) {
	// 获取进程列表
	processes, err := process.Processes()
	if err != nil {
		logger.Fatal(err)
	}
	// 定义返回结果
	processData := list.New()
	// connection
	connectData := list.New()

	for _, process := range processes {
		path, err := process.Exe()
		if err != nil {
			logger.Warn(err)
		}
		// 排除
		if len(path) == 0 {
			continue
		}
		if !strings.Contains(path, "java") {
			continue
		}

		ppid, err := process.Ppid()
		if err != nil {
			logger.Warn(err)
		}
		exec, err := process.Cmdline()
		if err != nil {
			logger.Warn(err)
		}
		cpu, err := process.CPUPercent()
		if err != nil {
			logger.Warn(err)
		}
		mem, err := process.MemoryPercent()
		if err != nil {
			logger.Warn(err)
		}
		//获取网络链接状态
		connStatList, err := process.Connections()
		if err != nil {
			logger.Warn(err)
		}

		pData := PData{
			Pid:  process.Pid,
			Ppid: ppid,
			Path: path,
			Exec: exec,
			Cpu:  cpu,
			Mem:  mem,
		}
		processData.PushBack(pData)

		for _, connStat := range connStatList {
			cData := CData{
				Pid:        process.Pid,
				LocalAddr:  connStat.Laddr.String(),
				RemoteAddr: connStat.Raddr.String(),
				Status:     connStat.Status,
			}
			connectData.PushBack(cData)
		}
	}

	logger.Info("processData:")
	return processData, connectData
}
