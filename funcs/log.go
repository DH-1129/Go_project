package funcs

import (
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	date := time.Now().Format("2006-01-02")
	current_dir, _ := os.Getwd()
	path := current_dir + "/logs"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	dir := path + "/" + date + ".log"
	// 创建每天的日志文件，如果不存在就新创建，存在则追加
	file, err := os.OpenFile(dir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
}

// 重构对应等级的日志文件
func Info(args ...interface{}) {
	logger.SetPrefix("[INFO] ")
	logger.Println(args...)

}
func Warning(args ...interface{}) {
	logger.SetPrefix("[Waring] ")
	logger.Println(args...)
}

// 避免和error重名而使用了Danger
func Danger(args ...interface{}) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(args...)
}
