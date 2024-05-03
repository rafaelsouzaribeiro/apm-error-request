package customtransaction

import (
	"context"
	"errors"
	"sync"

	"github.com/rafaelsouzaribeiro/apm-error-request/pkg"
	"go.elastic.co/apm/v2"
)

func (conf *Configs) Send(er *pkg.Erros) {

	ctx := context.Background()
	transaction := apm.DefaultTracer().StartTransaction("Error "+conf.Name, er.TransactionType)
	ctx = apm.ContextWithTransaction(ctx, transaction)
	span, _ := apm.StartSpan(ctx, "Error "+conf.Name, er.SpanType)
	span.Context.SetLabel("error-group", er.Erros)
	defer span.End()
	defer transaction.End()
	apm.CaptureError(ctx, errors.New(er.Erros)).Send()

}

func (confs *Configs) Alert(errs string) {

	err := pkg.Erros{
		Erros:           errs,
		TransactionType: "custom-errors",
		SpanType:        "Send",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		confs.Send(&err)
		wg.Done()
	}()

	wg.Wait()
}