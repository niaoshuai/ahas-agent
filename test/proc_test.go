/**
 * @Author: niaoshuai
 * @Date: 2020/8/10 4:02 下午
 */
package test

import (
	"ahas-agent/pkg/logger"
	"ahas-agent/pkg/proc"
	"github.com/shirou/gopsutil/process"
	"testing"
)

func TestGetPsData(t *testing.T) {
	_, err := process.Processes()
	if err != nil {
		logger.Fatal(err)
	}
}

func TestGetPsData1(t *testing.T) {
	proc.GetProcessData()
}
