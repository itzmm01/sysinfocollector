package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"monitor/logger"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// Run
func CmdRun(command string) error {
	var result []byte
	var err error

	sysType := runtime.GOOS
	if sysType == "windows" {
		result, err = exec.Command("cmd", "/c", command).CombinedOutput()
		// logger.Error("no support system: ", sysType)
	} else if sysType == "linux" {
		result, err = exec.Command("/bin/sh", "-c", command).CombinedOutput()
	} else {
		logger.Logger.Errorf(fmt.Sprintf("no support system: %v", sysType))
	}

	if err != nil {
		logger.Logger.Errorf(fmt.Sprintf(
			"run cmd failed: %v, %v", err.Error(), ConvertByte2String(result, "GB18030"),
		))

	}
	return err
}

// ConvertByte2String
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// CheckFileIsExist
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// WriteFileA
func WriteFile(filePath, content string) {
	if !CheckFileIsExist(filePath) {
		os.Create(filePath)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		logger.Logger.Errorf("open file failed ", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(content)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

// ArchiveTar
func ArchiveTar(remove bool, directory, srcName, destName string) {
	cmdStr := ""
	if remove {
		cmdStr = fmt.Sprintf("cd %v && tar zcf %v %v --remove-files", directory, destName, srcName)
	} else {
		cmdStr = fmt.Sprintf("cd %v && tar zcf %v %v ", directory, destName, srcName)
	}
	CmdRun(cmdStr)

}
