package main

import (
	"sync"

	"github.com/rafaelsouzaribeiro/apm-error-request/pkg"
	"github.com/rafaelsouzaribeiro/apm-error-request/pkg/mongodb"
)

func main() {

	conf := mongodb.NewConfigs("mongodb")
	err := pkg.Erros{
		Erros:           "Test",
		TransactionType: "request",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		conf.Send(&err)
		wg.Done()
	}()

	wg.Wait()

}
