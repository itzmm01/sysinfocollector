package logger

import (
	"fmt"
	"time"

	go_logger "github.com/phachon/go-logger"
)

var (
	// logger
	Logger *go_logger.Logger
)

func init() {
	NowDay := time.Now().Format("20060102")
	logger := go_logger.NewLogger()
	logger.SetAsync()
	logger.Detach("console")

	format := "%timestamp_format% [%level_string%] [%file%:%line%] %body%"
	// console adapter config
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,   // Does the text display the color
		JsonFormat: false,  // Whether or not formatted into a JSON string
		Format:     format, // JsonFormat is false, logger message output to console format string
	}
	// add output to the console
	logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)
	// file adapter config
	fileConfig := &go_logger.FileConfig{
		Filename: fmt.Sprintf("/var/log/monitor/monitor.log.%v", NowDay), // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		// LevelFileName: map[int]string{
		// 	logger.LoggerLevel("error"): "./error.log", // The error level log is written to the error.log file.
		// 	logger.LoggerLevel("info"):  "./info.log",  // The info level log is written to the info.log file.
		// 	logger.LoggerLevel("debug"): "./debug.log", // The debug level log is written to the debug.log file.
		// },
		MaxSize:    0,      // File maximum (KB), default 0 is not limited
		MaxLine:    0,      // The maximum number of lines in the file, the default 0 is not limited
		DateSlice:  "d",    // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: false,  // Whether the file data is written to JSON formatting
		Format:     format, // JsonFormat is false, logger message written to file format string
	}
	// add output to the file
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	Logger = logger
}
