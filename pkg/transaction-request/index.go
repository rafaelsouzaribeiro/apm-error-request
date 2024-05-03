package transactionrequest

import (
	"errors"
	"sync"

	"github.com/rafaelsouzaribeiro/apm-error-request/pkg"
	"go.elastic.co/apm/v2"
)

func (conf *Configs) Send(er *pkg.Erros) {

	tx := apm.DefaultTracer().StartTransaction(conf.Name, er.TransactionType)
	defer tx.End()
	e := apm.DefaultTracer().NewError(errors.New(er.Erros))
	e.SetTransaction(tx)
	e.Send()
}

func (confs *Configs) Log(errs string) {

	err := pkg.Erros{
		Erros:           errs,
		TransactionType: "request",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		confs.Send(&err)
		wg.Done()
	}()

	wg.Wait()
}
