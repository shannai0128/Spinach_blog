package common

import "fmt"

type Err struct {
	err error
	ErrCode int `json:"err_code"`
	Msg string `json:"msg"`
}

func (err Err) CustomError() string {
	return fmt.Sprintf("Error code:%s, error msg:%s", err.ErrCode, err.Msg)
}

func (err Err) Error() string {
	return err.err.Error() + err.Msg
}

type Response struct {
	Err
	Data interface{} `json:"data"`
}

