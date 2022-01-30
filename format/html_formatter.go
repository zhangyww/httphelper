package format

import (
	"bytes"
	"github.com/zhangyww/httphelper/mdl"
	"html/template"
)

const (
	htmlFormatTemplate = `<html>
    <head>
        <title>httphelper</title>
        <style>
            body {
                margin: 0px;
                padding: 0px;
            }
            .request-info-key {
                background-color: #FBFBFB;
            }
            .request-info-value {
                background-color: #FBFBFB;
            }
            .request-method-key {
                background-color: rgb(252, 233, 221);
            }
            .request-uri-key {
                background-color: rgb(231, 252, 221);
            }
            .request-proto-key {
                background-color: rgb(221, 243, 252);
            }
            .request-header-key {
                background-color: #FBFBFB;
            }
            .request-header-value {
                background-color: #FBFBFB;
            }
            .request-body-key {
                background-color: #fff19f;
            }
        </style>
    </head>
    <body>
        <div id="http-request-data">
            <table>
                <tbody>
                    <tr class="request-info">
                        <td class="request-info-key">remote_ip</td>
                        <td class="request-info-value">{{.Request.RemoteIP}}</td>
                    </tr>
                    <tr class="request-info">
                        <td class="request-info-key">remote_port</td>
                        <td class="request-info-value">{{.Request.RemotePort}}</td>
                    </tr>
                    <tr class="request-info">
                        <td class="request-info-key">request_host</td>
                        <td class="request-info-value">{{.Request.Host}}</td>
                    </tr>
                    <tr class="request-method">
                        <td class="request-method-key">method</td>
                        <td class="request-method-value">{{.Request.Method}}</td>
                    </tr>
                    <tr class="request-uri">
                        <td class="request-uri-key">uri</td>
                        <td class="request-uri-value">{{.Request.Uri}}</td>
                    </tr>
                    <tr class="request-proto">
                        <td class="request-proto-key">proto</td>
                        <td class="request-proto-value">{{.Request.Proto}}</td>
                    </tr>
                    <tr>
                        <td>headers</td>
                        <td>values</td>
                    </tr>
					{{range $hkey, $hvalues := .Request.Headers}}
						{{range $index, $hvalue := $hvalues}}
                    <tr class="request-header">
                        <td class="request-header-key">{{$hkey}}</td>
                        <td class="request-header-value">{{$hvalue}}</td>
                    </tr>
						{{end}}
					{{end}}
                    <tr class="request-body">
                        <td class="request-body-key">body</td>
                        <td class="request-body-value">
                            <prev>{{.Request.Body}}</prev>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </body>
</html>`
)

type HtmlFormatter struct {
	template *template.Template
}

func NewHtmlFormatter() *HtmlFormatter {
	templ, err := template.New("html_template").Parse(htmlFormatTemplate)
	if err != nil {
		panic(err.Error())
	}
	return &HtmlFormatter{
		template: templ,
	}
}

func (this *HtmlFormatter) Format(data *mdl.HttpData) ([]byte, error) {
	buff := bytes.Buffer{}
	err := this.template.Execute(&buff, data)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
