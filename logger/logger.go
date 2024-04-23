package logger

import (
	"log/slog"
	"os"
	"path"
)

const logDirName = "logs"
const logFileName = "zhulong.log"
const loggerTag = "[logger]"

func createDir(dirPath string) {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				panic(loggerTag + err.Error())
			}
		} else {
			panic(loggerTag + err.Error())
		}
	}
}

func New() *slog.Logger {
	wd, _ := os.Getwd()
	logDir := path.Join(wd, logDirName)
	logFilePath := path.Join(logDir, logFileName)

	createDir(logDir)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(loggerTag + err.Error())
	}

	return slog.New(slog.NewJSONHandler(file, nil))
}
