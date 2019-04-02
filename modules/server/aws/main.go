package aws

import (
	"log"
	"serverlessapi/modules/server/handlers"

	"github.com/apex/gateway"
)

func StartLambda() {
	router := handlers.Routes()
	log.Fatal(gateway.ListenAndServe("", router), nil)
}
