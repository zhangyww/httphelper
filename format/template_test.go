package format

import (
	"bytes"
	"fmt"
	"github.com/zhangyww/httphelper/mdl"
	"html/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	data := &mdl.HttpData{
		Request: mdl.HttpRequest{
			RemoteIP:   "123.123.123.123",
			RemotePort: 0,
			Method:     "",
			Host:       "",
			Uri:        "",
			Proto:      "",
			Headers:    nil,
			Body:       "",
		},
	}

	templ, err := template.New("test_templ").Parse("{{.Request.RemoteIP}}")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	buff := bytes.Buffer{}
	err = templ.Execute(&buff, data)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	fmt.Println(buff.String())
}
