package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)


type Level int

type  Conf struct {
	FileName string
	FilePath string
	PrefixUrl string
	LogSaveName string
	LogSavePath string
	LogMaxSize string
	IpBlackExpire int64
	SecretKey string
}

var (
	ConfObj Conf
	logPrefix = ""
	Logger  *log.Logger
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// 获取配置信息
func init()  {
	conf, err :=config.NewConfig("ini","conf/app.ini")
	if err!=nil{
		err = fmt.Errorf("open config file failed, error:%s", err)
		fmt.Println(err)
	}
	logSavePath := conf.String("app::LogSavePath")
	logSaveName := conf.String("app::LogSaveName")
	logMaxSize  := conf.String("app::LogMaxSize")
	IpBlackExpire := conf.String("app::IpBlackExpire")
	SecretKey := conf.String("app::SecretKey")
	ConfObj.LogMaxSize = logMaxSize
	ConfObj.LogSaveName = logSaveName
	ConfObj.LogSavePath = logSavePath
	IntIpBlackExpire, err := strconv.ParseInt(IpBlackExpire,10,64)
	ConfObj.IpBlackExpire = IntIpBlackExpire
	ConfObj.SecretKey = SecretKey
}

// 初始化logger
func InitLogger(logSaveName, logSavePath string) (file *os.File,err error) {
	logSaveName = logSaveName +time.Now().Format("2006010215")+".log"
	_, err =os.Stat(logSavePath+logSaveName)
	if err == nil{
		fmt.Println("Log directory is normal")
	}
	if os.IsNotExist(err){
		fmt.Println("Log directory is not exist, start create...")
		err = os.Mkdir(logSavePath,os.ModePerm)
		if err != nil{
			fmt.Println(err)
		}
		err = ioutil.WriteFile(logSavePath+logSaveName,nil,os.ModePerm)
		if err != nil{
			fmt.Printf("create log file error: %s",err)
		}
	}
	file, err =os.OpenFile(logSavePath+logSaveName,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err !=nil{
		fmt.Printf("open %s error: %s", logSavePath+logSaveName,err)
	}
	Logger =log.New(file,"",log.LstdFlags)
	return file ,nil
}

// 日志文件句柄
func LogFileHandle() (file *os.File) {
	logSaveName := ConfObj.LogSaveName +time.Now().Format("2006010215:04")+".log"
	err := ioutil.WriteFile(ConfObj.LogSavePath+logSaveName,nil,os.ModePerm)
	if err != nil{
		fmt.Printf("create log file error: %s",err)
	}
	file, err =os.OpenFile(ConfObj.LogSavePath+logSaveName,os.O_RDWR|os.O_APPEND,os.ModePerm)
	if err !=nil{
		fmt.Printf("open %s error: %s", ConfObj.LogSavePath+logSaveName,err)
	}
	Logger =log.New(file,"",log.LstdFlags)
	fmt.Println("Jubing run....")
	return file
}

// 日志信息前缀
func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}else {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}
	Logger.SetPrefix(logPrefix)
}

// 加载日志
func ParseConfigInfo(level Level) (file *os.File) { // "conf/"+"app.ini"
	file,err := InitLogger(ConfObj.LogSaveName, ConfObj.LogSavePath)
	if err != nil{
		fmt.Println("Init log error: %s", err)
	}

	setPrefix(level)
	return
}


func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	Logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	Logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	Logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	Logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	Logger.Fatalln(v)
}


