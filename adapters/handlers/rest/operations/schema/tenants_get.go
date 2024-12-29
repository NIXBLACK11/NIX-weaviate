// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// TenantsGetHandlerFunc turns a function with the right signature into a tenants get handler
type TenantsGetHandlerFunc func(TenantsGetParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn TenantsGetHandlerFunc) Handle(params TenantsGetParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// TenantsGetHandler interface for that can handle valid tenants get params
type TenantsGetHandler interface {
	Handle(TenantsGetParams, *models.Principal) middleware.Responder
}

// NewTenantsGet creates a new http.Handler for the tenants get operation
func NewTenantsGet(ctx *middleware.Context, handler TenantsGetHandler) *TenantsGet {
	return &TenantsGet{Context: ctx, Handler: handler}
}

/*
	TenantsGet swagger:route GET /schema/{className}/tenants schema tenantsGet

Get the list of tenants.

get all tenants from a specific class
*/
type TenantsGet struct {
	Context *middleware.Context
	Handler TenantsGetHandler
}

func (o *TenantsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTenantsGetParams()
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
