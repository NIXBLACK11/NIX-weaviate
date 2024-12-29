// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// HasPermissionHandlerFunc turns a function with the right signature into a has permission handler
type HasPermissionHandlerFunc func(HasPermissionParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn HasPermissionHandlerFunc) Handle(params HasPermissionParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// HasPermissionHandler interface for that can handle valid has permission params
type HasPermissionHandler interface {
	Handle(HasPermissionParams, *models.Principal) middleware.Responder
}

// NewHasPermission creates a new http.Handler for the has permission operation
func NewHasPermission(ctx *middleware.Context, handler HasPermissionHandler) *HasPermission {
	return &HasPermission{Context: ctx, Handler: handler}
}

/*
	HasPermission swagger:route POST /authz/roles/{id}/has-permission authz hasPermission

Check whether role possesses this permission.
*/
type HasPermission struct {
	Context *middleware.Context
	Handler HasPermissionHandler
}

func (o *HasPermission) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHasPermissionParams()
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
