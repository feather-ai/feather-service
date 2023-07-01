// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1DebugExecuteRequestSchemaHandlerFunc turns a function with the right signature into a get v1 debug execute request schema handler
type GetV1DebugExecuteRequestSchemaHandlerFunc func(GetV1DebugExecuteRequestSchemaParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1DebugExecuteRequestSchemaHandlerFunc) Handle(params GetV1DebugExecuteRequestSchemaParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetV1DebugExecuteRequestSchemaHandler interface for that can handle valid get v1 debug execute request schema params
type GetV1DebugExecuteRequestSchemaHandler interface {
	Handle(GetV1DebugExecuteRequestSchemaParams, interface{}) middleware.Responder
}

// NewGetV1DebugExecuteRequestSchema creates a new http.Handler for the get v1 debug execute request schema operation
func NewGetV1DebugExecuteRequestSchema(ctx *middleware.Context, handler GetV1DebugExecuteRequestSchemaHandler) *GetV1DebugExecuteRequestSchema {
	return &GetV1DebugExecuteRequestSchema{Context: ctx, Handler: handler}
}

/* GetV1DebugExecuteRequestSchema swagger:route GET /v1/debug/executeRequestSchema getV1DebugExecuteRequestSchema

Internal

*/
type GetV1DebugExecuteRequestSchema struct {
	Context *middleware.Context
	Handler GetV1DebugExecuteRequestSchemaHandler
}

func (o *GetV1DebugExecuteRequestSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1DebugExecuteRequestSchemaParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
