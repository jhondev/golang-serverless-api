package main

import "serverlessapi/pkg/testpkg1"
import "serverlessapi/pkg/server/http"

func main() {
	testpkg1.PrintMessage()
	http.PrintHttp()
	http.RunServer()
}
