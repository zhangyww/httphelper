package main

import (
	"flag"
	"fmt"
	"github.com/zhangyww/httphelper/flags"
	"github.com/zhangyww/httphelper/format"
	"github.com/zhangyww/httphelper/handler"
	"net/http"
)

func main() {
	flag.Parse()

	httpHelperHandler := handler.NewHttpHelperHandler()
	httpHelperHandler.SetFormatter(format.NewHtmlFormatter())

	http.Handle("/", httpHelperHandler)

	fmt.Printf("httphelper listen at %s\n", flags.ListenAddr)
	err := http.ListenAndServe(flags.ListenAddr, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("bye~")
}
