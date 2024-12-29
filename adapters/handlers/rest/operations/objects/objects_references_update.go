// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsReferencesUpdateHandlerFunc turns a function with the right signature into a objects references update handler
type ObjectsReferencesUpdateHandlerFunc func(ObjectsReferencesUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsReferencesUpdateHandlerFunc) Handle(params ObjectsReferencesUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsReferencesUpdateHandler interface for that can handle valid objects references update params
type ObjectsReferencesUpdateHandler interface {
	Handle(ObjectsReferencesUpdateParams, *models.Principal) middleware.Responder
}

// NewObjectsReferencesUpdate creates a new http.Handler for the objects references update operation
func NewObjectsReferencesUpdate(ctx *middleware.Context, handler ObjectsReferencesUpdateHandler) *ObjectsReferencesUpdate {
	return &ObjectsReferencesUpdate{Context: ctx, Handler: handler}
}

/*
	ObjectsReferencesUpdate swagger:route PUT /objects/{id}/references/{propertyName} objects objectsReferencesUpdate

Replace all references to a class-property.

Replace all references in cross-reference property of an object.
*/
type ObjectsReferencesUpdate struct {
	Context *middleware.Context
	Handler ObjectsReferencesUpdateHandler
}

func (o *ObjectsReferencesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewObjectsReferencesUpdateParams()
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
