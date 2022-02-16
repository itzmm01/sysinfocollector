package monitorcmd

import (
	"fmt"
	"monitor/config"
	"monitor/logger"
	"monitor/tools"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// 清理文件
func cleanHistoryFile(match, path string, saveNum int) {
	// var fileList []string

	fileList, _ := filepath.Glob(filepath.Join(path, match))
	// if err != nil {
	// 	logger.Logger.Errorf(err.Error())

	// }
	// for i := range filepathNames {
	// 	fileList = append(fileList, filepathNames[i])
	// }
	sort.Strings(fileList)
	countNum := len(fileList) - saveNum
	if countNum > 0 {
		for _, v := range fileList[0:countNum] {
			err := os.Remove(v)
			if err != nil {
				logger.Logger.Errorf(fmt.Sprintf("clean file fail %v", err.Error()))
			}
		}
	}
}

// run clean
func RunClean(fileNameList map[string]string, logDir string, logSaveDay int) {
	for name := range fileNameList {
		ArchiveHistoryFile(fmt.Sprintf("%v*.log", name), logDir, 1)
		cleanHistoryFile(fmt.Sprintf("%v*.tar.gz", name), logDir, logSaveDay)
	}
}

// 压缩文件
func ArchiveHistoryFile(match, path string, saveNum int) {
	fileList, _ := filepath.Glob(filepath.Join(path, match))
	sort.Strings(fileList)
	countNum := len(fileList) - saveNum
	if countNum > 0 {
		for _, v := range fileList[0:countNum] {
			tools.ArchiveTar(true, path, v, v+".tar.gz")
		}
	}
}

// Run
func Run(config config.Config, cmdList map[string]string) {
	logDir := *config.LogDir

	nowTime := time.Now().Format("20060102")
	logger.Logger.Infof("collection info start")

	if _, err := os.Stat(logDir); err != nil {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			logger.Logger.Errorf(err.Error())
		}
	}
	for name, cmd := range cmdList {
		logFile := fmt.Sprintf("%v/%v.%v.log", logDir, name, nowTime)
		nowTime1 := time.Now().Format("2006-01-02 15:04:05")
		tools.CmdRun(fmt.Sprintf("echo '%v: -------------------------' >> %v", nowTime1, logFile))
		cmdStr := fmt.Sprintf("%v >> %v 2>&1", cmd, logFile)
		tools.CmdRun(cmdStr)
	}
	logger.Logger.Infof("collection info end")
}

// CheckFileIsExist
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
