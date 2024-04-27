package main

import (
	"github.com/rafaelsouzaribeiro/apm-error-request/pkg/mongodb"
)

func main() {

	conf := mongodb.NewConfigs("mongodb")
	conf.Alert("Test")

}
