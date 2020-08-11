/**
 * @Author: niaoshuai
 * @Date: 2020/8/10 3:47 下午
 */
package proc

import (
	"ahas-agent/pkg/logger"
	"container/list"
	"github.com/shirou/gopsutil/process"
	"strconv"
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
			continue
		}
		// 排除
		if len(path) == 0 {
			continue
		}
		if !strings.Contains(path, "java") && !strings.Contains(path, "ssh") {
			continue
		}

		ppid, err := process.Ppid()
		if err != nil {
			logger.Warn(err)
			continue
		}
		exec, err := process.Cmdline()
		if err != nil {
			logger.Warn(err)
			continue
		}
		cpu, err := process.CPUPercent()
		if err != nil {
			logger.Warn(err)
			continue
		}
		mem, err := process.MemoryPercent()
		if err != nil {
			logger.Warn(err)
			continue
		}
		//获取网络链接状态
		connStatList, err := process.Connections()
		if err != nil {
			logger.Warn(err)
			continue
		}

		pid := process.Pid

		pData := PData{
			Pid:  pid,
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
				LocalAddr:  connStat.Laddr.IP,
				RemoteAddr: connStat.Raddr.IP,
				Status:     connStat.Status,
			}
			connectData.PushBack(cData)
		}
	}

	logger.Info(strconv.Itoa(processData.Len()))
	return processData, connectData
}
