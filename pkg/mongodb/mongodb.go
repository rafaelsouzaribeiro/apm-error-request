package mongodb

import (
	"errors"

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
