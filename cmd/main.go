package main

import mongodb "github.com/rafaelsouzaribeiro/apm-error-request/pkg/transaction-request"

func main() {

	conf := mongodb.NewConfigs("mongodb")
	conf.Alert("Test")

}
