package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Settings struct {
	FileName string
	FilePath string
	PrefixUrl string
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

func InitConfigInfo(fileName string, settings interface{}) (error) {
	typ := reflect.TypeOf(settings)
	val := reflect.ValueOf(settings)

	// grammar check
	if typ.Kind() != reflect.Ptr{
		err := errors.New("a point object is required")
		return err
	}

	if typ.Elem().Kind() != reflect.Struct{
		err := errors.New("the point object's item need to be struct obj")
		return err
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		err_info := fmt.Sprintf("open %s failed, %s", fileName,err)
		err = errors.New(err_info)
		return err
	}
	lineSlice := strings.Split(string(data), "\n")

	for index, line := range lineSlice{
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#"){
			continue
		}

		equalIndex := strings.Index(line,"=")
		if equalIndex == -1{
			err = fmt.Errorf("%d line grammar error", index+1)
			return err
		}

		key := line[:equalIndex]
		value := line[equalIndex+1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if len(key) == 0{
			err = fmt.Errorf("%d line grammar error", index+1)
			return err
		}

		// assign value
		for i :=0;i<typ.Elem().NumField(); i++{
			filed := typ.Elem().Field(i)
			// TODO
			fmt.Println(filed)
		}
	}
	return err
}

func InitLogger(filename, filepath string)  {
	filename = filename +time.Now().Format("20060102")
	file, err :=os.Open(filepath+filename)
	if err !=nil{
		fmt.Printf("open %s error: %s", filepath+filename,err)
	}
	logger =log.New(file,"",log.LstdFlags)
}

func Debug(v ...interface{})  {
	setPrefix(DEBUG)
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