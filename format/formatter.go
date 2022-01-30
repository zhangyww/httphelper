package format

import "github.com/zhangyww/httphelper/mdl"

type Formatter interface {
	Format(data *mdl.HttpData) ([]byte, error)
}
