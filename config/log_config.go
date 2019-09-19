package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Settings struct {
	FileName string
	FilePath string
	PrefixUrl string
	LogSaveName string
	LogSavePath string
	LogMaxSize int
}

type Level int

var (
	logPrefix = ""
	logger  *log.Logger
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func InitLogger(logSaveName, logSavePath string) (err error) {
	logSaveName = logSaveName +time.Now().Format("20060102")+".log"
	_, err =os.Stat(logSavePath+logSaveName)
	if err == nil{
		fmt.Println("Log directory is normal")
	}
	if os.IsNotExist(err){
		fmt.Println("Log directory is not exist, start create...")
		err = os.Mkdir(logSavePath,os.ModePerm)
		if err != nil{
			fmt.Printf("create log directory error: %s",err)
		}
		err = ioutil.WriteFile(logSavePath+logSaveName,nil,os.ModePerm)
		if err != nil{
			fmt.Printf("create log file error: %s",err)
		}
	}
	file, err :=os.Open(logSavePath+logSaveName)
	if err !=nil{
		fmt.Printf("open %s error: %s", logSavePath+logSaveName,err)
	}
	logger =log.New(file,"",log.LstdFlags)
	return
}

func ParseConfigInfo() error { // "conf/"+"app.ini"
	conf, err :=config.NewConfig("ini","conf/app.ini")
	if err!=nil{
		err = fmt.Errorf("open config file failed, error:%s", err)
		return err
	}
	logSavePath := conf.String("app::LogSavePath")

	logSaveName := conf.String("app::LogSaveName")

	err = InitLogger(logSaveName, logSavePath)
	return err
}


func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}else {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}
	logger.SetPrefix(logPrefix)
}