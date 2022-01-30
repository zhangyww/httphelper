package handler

import (
	"fmt"
	"github.com/zhangyww/httphelper/format"
	"github.com/zhangyww/httphelper/mdl"
	"net/http"
)

type HttpHelperHandler struct {
	formatter format.Formatter
}

func NewHttpHelperHandler() *HttpHelperHandler {
	return &HttpHelperHandler{
		formatter: format.NewJsonFormatter(format.WithIndent(true)),
	}
}

func (this *HttpHelperHandler) SetFormatter(f format.Formatter) {
	this.formatter = f
}

func (this *HttpHelperHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	data := new(mdl.HttpData)
	data.Request = *mdl.RequestData(req)
	respBytes, err := this.formatter.Format(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writer.Write(respBytes)
}
