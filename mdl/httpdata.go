package mdl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HttpData struct {
	Request  HttpRequest  `json:"request"`
	Response HttpResponse `json:"response"`
}

type HttpRequest struct {
	RemoteIP   string              `json:"remoteIP"`   // 收到http请求时的远程公网IP
	RemotePort int                 `json:"remotePort"` // 收到http请求时的远程公网端口
	Method     string              `json:"method"`
	Host       string              `json:"host"`
	Uri        string              `json:"uri"`
	Proto      string              `json:"proto"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
}

func RequestData(req *http.Request) *HttpRequest {
	data := new(HttpRequest)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}

	ipPort := strings.Split(req.RemoteAddr, ":")
	if len(ipPort) > 0 {
		data.RemoteIP = ipPort[0]
	}
	if len(ipPort) > 1 {
		data.RemotePort, _ = strconv.Atoi(ipPort[1])
	}
	data.Proto = req.Proto
	data.Host = req.Host
	data.Uri = req.RequestURI
	data.Method = req.Method
	data.Headers = req.Header
	data.Body = string(body)
	return data
}

type HttpResponse struct {
}
