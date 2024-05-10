package customtransaction

import (
	"context"
	"errors"
	"fmt"
	"runtime"
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

func (confs *Configs) Log(errs, transactiontype string) {

	pc, _, line, _ := runtime.Caller(1)

	callingFunc := "Unknown"
	if pc != 0 {
		callingFunc = runtime.FuncForPC(pc).Name()
	}

	errorText := fmt.Sprintf("Error occurred in function %s at line %d: %s", callingFunc, line, errs)

	err := pkg.Erros{
		Erros:           errorText,
		TransactionType: transactiontype,
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
