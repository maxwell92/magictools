package merror

import (
	"encoding/json"
	"log"
)

type Errno uintptr

const (
	ETIMEOUT Errno = 1 /*Operation Time Out*/
)

var errors = [...]string{
	ETIMEOUT: "Operation Time out",
}

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(errors) {
		return errors[e]
	}
	return New(100, "error code not found").Error()
}

type yerror interface {
	Error() string
}

type YceError struct {
	code    int32
	message string
}

func New(code int32, text string) *YceError {
	return &YceError{code: code, message: text}
}

func (e *YceError) Error() string {
	return e.message
}

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func (e *YceError) EncodeSelf() []byte {
	r := &Response{
		Code:    e.code,
		Message: e.message,
		Data:    nil,
	}
	log.Println(r)
	errJSON, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(errJSON)
	log.Printf("%s\n", string(errJSON))
	return errJSON
}
