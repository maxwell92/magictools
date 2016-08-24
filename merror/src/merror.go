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
	return New(100, "error code not found", nil).Error()
}

type yerror interface {
	Error() string
}

type YceError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    []byte `json:"data,omitempty"`
}

func New(code int32, text string, data []byte) *YceError {
	return &YceError{Code: code, Message: text, Data: data}
}

func (e *YceError) Error() string {
	return e.Message
}

func (e *YceError) EncodeSelf() []byte {
	log.Println(e)
	errJSON, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(errJSON)
	log.Printf("%s\n", string(errJSON))
	return errJSON
}
