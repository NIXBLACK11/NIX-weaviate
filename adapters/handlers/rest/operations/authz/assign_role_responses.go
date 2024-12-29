// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// AssignRoleOKCode is the HTTP code returned for type AssignRoleOK
const AssignRoleOKCode int = 200

/*
AssignRoleOK Role assigned successfully

swagger:response assignRoleOK
*/
type AssignRoleOK struct {
}

// NewAssignRoleOK creates AssignRoleOK with default headers values
func NewAssignRoleOK() *AssignRoleOK {

	return &AssignRoleOK{}
}

// WriteResponse to the client
func (o *AssignRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AssignRoleBadRequestCode is the HTTP code returned for type AssignRoleBadRequest
const AssignRoleBadRequestCode int = 400

/*
AssignRoleBadRequest Bad request

swagger:response assignRoleBadRequest
*/
type AssignRoleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAssignRoleBadRequest creates AssignRoleBadRequest with default headers values
func NewAssignRoleBadRequest() *AssignRoleBadRequest {

	return &AssignRoleBadRequest{}
}

// WithPayload adds the payload to the assign role bad request response
func (o *AssignRoleBadRequest) WithPayload(payload *models.ErrorResponse) *AssignRoleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the assign role bad request response
func (o *AssignRoleBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssignRoleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssignRoleUnauthorizedCode is the HTTP code returned for type AssignRoleUnauthorized
const AssignRoleUnauthorizedCode int = 401

/*
AssignRoleUnauthorized Unauthorized or invalid credentials.

swagger:response assignRoleUnauthorized
*/
type AssignRoleUnauthorized struct {
}

// NewAssignRoleUnauthorized creates AssignRoleUnauthorized with default headers values
func NewAssignRoleUnauthorized() *AssignRoleUnauthorized {

	return &AssignRoleUnauthorized{}
}

// WriteResponse to the client
func (o *AssignRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AssignRoleForbiddenCode is the HTTP code returned for type AssignRoleForbidden
const AssignRoleForbiddenCode int = 403

/*
AssignRoleForbidden Forbidden

swagger:response assignRoleForbidden
*/
type AssignRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAssignRoleForbidden creates AssignRoleForbidden with default headers values
func NewAssignRoleForbidden() *AssignRoleForbidden {

	return &AssignRoleForbidden{}
}

// WithPayload adds the payload to the assign role forbidden response
func (o *AssignRoleForbidden) WithPayload(payload *models.ErrorResponse) *AssignRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the assign role forbidden response
func (o *AssignRoleForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssignRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AssignRoleNotFoundCode is the HTTP code returned for type AssignRoleNotFound
const AssignRoleNotFoundCode int = 404

/*
AssignRoleNotFound role or user is not found.

swagger:response assignRoleNotFound
*/
type AssignRoleNotFound struct {
}

// NewAssignRoleNotFound creates AssignRoleNotFound with default headers values
func NewAssignRoleNotFound() *AssignRoleNotFound {

	return &AssignRoleNotFound{}
}

// WriteResponse to the client
func (o *AssignRoleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// AssignRoleInternalServerErrorCode is the HTTP code returned for type AssignRoleInternalServerError
const AssignRoleInternalServerErrorCode int = 500

/*
AssignRoleInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response assignRoleInternalServerError
*/
type AssignRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAssignRoleInternalServerError creates AssignRoleInternalServerError with default headers values
func NewAssignRoleInternalServerError() *AssignRoleInternalServerError {

	return &AssignRoleInternalServerError{}
}

// WithPayload adds the payload to the assign role internal server error response
func (o *AssignRoleInternalServerError) WithPayload(payload *models.ErrorResponse) *AssignRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the assign role internal server error response
func (o *AssignRoleInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssignRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
