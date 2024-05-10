package customtransaction

import (
	"context"
	"errors"
	"runtime"
	"strconv"
	"sync"

	"github.com/rafaelsouzaribeiro/apm-error-request/pkg"
	"go.elastic.co/apm/v2"
)

func (conf *Configs) Send(er *pkg.Erros, callingFunc string, line int) {

	ctx := context.Background()
	transaction := apm.DefaultTracer().StartTransaction("Error "+conf.Name, er.TransactionType)
	ctx = apm.ContextWithTransaction(ctx, transaction)
	span, _ := apm.StartSpan(ctx, "Error "+conf.Name, er.SpanType)
	span.Context.SetLabel("error-group", er.Erros)
	span.Context.SetLabel("calling-function", callingFunc)
	span.Context.SetLabel("error-line", strconv.Itoa(line))
	defer span.End()
	defer transaction.End()
	apm.CaptureError(ctx, errors.New(er.Erros)).Send()

}

func (confs *Configs) Log(errs, transactiontype string) {
	_, _, line, _ := runtime.Caller(1)

	callingFunc := "Unknown"
	pc, _, _, _ := runtime.Caller(1)
	if pc != 0 {
		callingFunc = runtime.FuncForPC(pc).Name()
	}

	println(callingFunc, line)

	err := pkg.Erros{
		Erros:           errs,
		TransactionType: transactiontype,
		SpanType:        callingFunc,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		confs.Send(&err, callingFunc, line)
		wg.Done()
	}()

	wg.Wait()
}
