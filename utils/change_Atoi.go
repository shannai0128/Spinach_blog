package utils

import (
	"github.com/c479096292/Spinach_blog/config"
	"strconv"
)

func ChangeAtoi(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil{
		config.Error(err)
	}
	return result
}
