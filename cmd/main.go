package main

import (
	"flag"
	"serverlessapi/modules/server/aws"
	"serverlessapi/modules/server/http"
)

var isDev *bool

func init() {
	isDev = flag.Bool("dev", false, "value for dev mode")
	flag.Parse()
}

func main() {
	if *isDev {
		http.RunServer()
	} else {
		aws.StartLambda()
	}
}
