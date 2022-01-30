package format

import (
	"encoding/json"
	"github.com/zhangyww/httphelper/mdl"
)

type JsonFormatter struct {
	option JsonFormatterOption
}

func NewJsonFormatter(optFuncs ...JsonFormatterOptionFunc) *JsonFormatter {
	option := defaultJsonFormatterOption
	for _, optFunc := range optFuncs {
		optFunc(&option)
	}

	return &JsonFormatter{
		option: option,
	}
}

func (this *JsonFormatter) Format(data *mdl.HttpData) ([]byte, error) {
	if this.option.indent {
		return json.MarshalIndent(data, "", "    ")
	}
	return json.Marshal(data)
}

type JsonFormatterOption struct {
	indent bool
}

var defaultJsonFormatterOption = JsonFormatterOption{
	indent: true,
}

type JsonFormatterOptionFunc func(option *JsonFormatterOption)

func WithIndent(indent bool) JsonFormatterOptionFunc {
	return func(option *JsonFormatterOption) {
		option.indent = indent
	}
}
