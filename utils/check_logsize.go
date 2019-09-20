package utils

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"os"
	"strconv"
)

func CheckLogSize(fileObj *os.File)  {
	for {
		maxSize, err := strconv.Atoi(config.ConfObj.LogMaxSize)
		if err != nil{
			fmt.Println("convert logMaxSize error:" ,err)
		}
		fileInfo, err := fileObj.Stat()
		if err != nil{
			fmt.Printf("acquire %s info error: %s", fileObj.Name(), err)
		}

		fileSize := fileInfo.Size()
		fileSiz := int(fileSize)

		if fileSiz >= maxSize{
			oldName := fileObj.Name()
			fileObj.Close()
			newFileObj := config.LogFileHandle()
			fileObj = newFileObj
			fmt.Printf("%s log file is fulled, create new log %s...", oldName, fileObj.Name())
		}
	}
}
