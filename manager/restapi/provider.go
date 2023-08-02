package restapi

import (
	"github.com/go-openapi/runtime"
	"github.com/sithell/perun/manager/restapi/operations"
	"net/http"
)

type registerProviderResponder struct {
	params operations.RegisterProviderParams
	app    *App
}

func (rs registerProviderResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	operations.NewRegisterProviderOK().WriteResponse(rw, producer)
}
