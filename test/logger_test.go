package test

import (
	"ahas-agent/pkg/logger"
	"testing"
)

func TestLogger(t *testing.T)  {
	logger.InitLog("tests.log")
	logger.Info("testInfo")
	//logger.Fatal(nil)
}
