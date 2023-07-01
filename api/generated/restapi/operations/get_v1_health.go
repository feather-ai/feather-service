// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1HealthHandlerFunc turns a function with the right signature into a get v1 health handler
type GetV1HealthHandlerFunc func(GetV1HealthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1HealthHandlerFunc) Handle(params GetV1HealthParams) middleware.Responder {
	return fn(params)
}

// GetV1HealthHandler interface for that can handle valid get v1 health params
type GetV1HealthHandler interface {
	Handle(GetV1HealthParams) middleware.Responder
}

// NewGetV1Health creates a new http.Handler for the get v1 health operation
func NewGetV1Health(ctx *middleware.Context, handler GetV1HealthHandler) *GetV1Health {
	return &GetV1Health{Context: ctx, Handler: handler}
}

/* GetV1Health swagger:route GET /v1/health getV1Health

Health check endpoint

*/
type GetV1Health struct {
	Context *middleware.Context
	Handler GetV1HealthHandler
}

func (o *GetV1Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1HealthParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}