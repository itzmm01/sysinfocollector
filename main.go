package main

import (
	"flag"
	"fmt"
	"monitor/config"
	"monitor/logger"
	"monitor/monitorcmd"

	"github.com/robfig/cron"
)

func main() {
	configInfo := config.Config{}

	configInfo.LogSaveDay = flag.Int("s", 5, "Log retention days")
	configInfo.RunFreq = flag.Int("m", 1, "Run frequency: 1(default) minute run once")

	logDir := "/var/log/monitor"
	configInfo.LogDir = &logDir
	flag.Parse()

	logger.Logger.Infof("Start monitor...")

	cmdList := map[string]string{
		"iostat": "iostat -x",
		"ps":     "ps",
		"sar":    "sar",
		"vmstat": "vmstat",
		"load":   "uptime",
		"top":    "top -b -n 1",
	}

	monitorcmd.RunClean(cmdList, *configInfo.LogDir, *configInfo.LogSaveDay)
	//创建一个cron实例
	cron2 := cron.New()
	err := cron2.AddFunc(fmt.Sprintf("00 */%v * * * *", *configInfo.RunFreq), func() {
		monitorcmd.Run(configInfo, cmdList)
		// monitorcmd.Test()
	})
	if err != nil {
		logger.Logger.Errorf(err.Error())
	}

	err1 := cron2.AddFunc("00 00,12 * * * *", func() {
		monitorcmd.RunClean(cmdList, *configInfo.LogDir, *configInfo.LogSaveDay)
	})
	if err1 != nil {
		logger.Logger.Errorf(err.Error())
	}
	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}

}
