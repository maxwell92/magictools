package merror

import (
	//"mlogger"
	"fmt"
	"log"
	"testing"
)

type abcController struct {
	//*iris.Context
	Err *YceError
	//Logger *Mlogger
	//...
}

func (a *abcController) LogAndResponse() {
	//a.Logger.Printf()
	//a.Write(string(a.Err.EncodeSelf()))
	log.Println(a.Err.Code)
	log.Println(a.Err.Message)

	fmt.Println(string(a.Err.EncodeSelf()))
}
func Test_LogAndRespose(t *testing.T) {
	abc := &abcController{
		Err: New(1, "This is error", nil),
		//Logger: mlogger.New(4),
	}

	abc.LogAndResponse()
}
