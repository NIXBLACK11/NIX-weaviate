// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsGetHandlerFunc turns a function with the right signature into a objects get handler
type ObjectsGetHandlerFunc func(ObjectsGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsGetHandlerFunc) Handle(params ObjectsGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsGetHandler interface for that can handle valid objects get params
type ObjectsGetHandler interface {
	Handle(ObjectsGetParams, *models.Principal) middleware.Responder
}

// NewObjectsGet creates a new http.Handler for the objects get operation
func NewObjectsGet(ctx *middleware.Context, handler ObjectsGetHandler) *ObjectsGet {
	return &ObjectsGet{Context: ctx, Handler: handler}
}

/*
	ObjectsGet swagger:route GET /objects/{id} objects objectsGet

Get a specific Object based on its UUID and a Object UUID. Also available as Websocket bus.

Get a specific object based on its UUID. Also available as Websocket bus.
*/
type ObjectsGet struct {
	Context *middleware.Context
	Handler ObjectsGetHandler
}

func (o *ObjectsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewObjectsGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
