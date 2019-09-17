package common

import "fmt"

type Err struct {
	ErrCode int `json:"err_code"`
	Msg string `json:"msg"`
}

func (err Err) Error() string {
	return fmt.Sprintf("Error code:%s, error msg:%s", err.ErrCode, err.Msg)
}

type Response struct {
	Err
	Data interface{} `json:"data"`
}

