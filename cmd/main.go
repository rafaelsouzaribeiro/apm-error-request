package main

import (
	"context"

	customtransaction "github.com/rafaelsouzaribeiro/apm-error-request/pkg/custom-transaction"
	mongodb "github.com/rafaelsouzaribeiro/apm-error-request/pkg/transaction-request"
)

func main() {

	// Transaction Request
	conf1 := mongodb.NewConfigs("mongodb1")
	conf1.Log("Test")

	// Custom Transaction
	conf2 := customtransaction.NewConfigs("mongodb2")
	conf2.Log("server error", "custom-errors", "main", context.Background())

}
