// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// RevokeRoleOKCode is the HTTP code returned for type RevokeRoleOK
const RevokeRoleOKCode int = 200

/*
RevokeRoleOK Role revoked successfully

swagger:response revokeRoleOK
*/
type RevokeRoleOK struct {
}

// NewRevokeRoleOK creates RevokeRoleOK with default headers values
func NewRevokeRoleOK() *RevokeRoleOK {

	return &RevokeRoleOK{}
}

// WriteResponse to the client
func (o *RevokeRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RevokeRoleBadRequestCode is the HTTP code returned for type RevokeRoleBadRequest
const RevokeRoleBadRequestCode int = 400

/*
RevokeRoleBadRequest Bad request

swagger:response revokeRoleBadRequest
*/
type RevokeRoleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRevokeRoleBadRequest creates RevokeRoleBadRequest with default headers values
func NewRevokeRoleBadRequest() *RevokeRoleBadRequest {

	return &RevokeRoleBadRequest{}
}

// WithPayload adds the payload to the revoke role bad request response
func (o *RevokeRoleBadRequest) WithPayload(payload *models.ErrorResponse) *RevokeRoleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke role bad request response
func (o *RevokeRoleBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeRoleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RevokeRoleUnauthorizedCode is the HTTP code returned for type RevokeRoleUnauthorized
const RevokeRoleUnauthorizedCode int = 401

/*
RevokeRoleUnauthorized Unauthorized or invalid credentials.

swagger:response revokeRoleUnauthorized
*/
type RevokeRoleUnauthorized struct {
}

// NewRevokeRoleUnauthorized creates RevokeRoleUnauthorized with default headers values
func NewRevokeRoleUnauthorized() *RevokeRoleUnauthorized {

	return &RevokeRoleUnauthorized{}
}

// WriteResponse to the client
func (o *RevokeRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RevokeRoleForbiddenCode is the HTTP code returned for type RevokeRoleForbidden
const RevokeRoleForbiddenCode int = 403

/*
RevokeRoleForbidden Forbidden

swagger:response revokeRoleForbidden
*/
type RevokeRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRevokeRoleForbidden creates RevokeRoleForbidden with default headers values
func NewRevokeRoleForbidden() *RevokeRoleForbidden {

	return &RevokeRoleForbidden{}
}

// WithPayload adds the payload to the revoke role forbidden response
func (o *RevokeRoleForbidden) WithPayload(payload *models.ErrorResponse) *RevokeRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke role forbidden response
func (o *RevokeRoleForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RevokeRoleNotFoundCode is the HTTP code returned for type RevokeRoleNotFound
const RevokeRoleNotFoundCode int = 404

/*
RevokeRoleNotFound role or user is not found.

swagger:response revokeRoleNotFound
*/
type RevokeRoleNotFound struct {
}

// NewRevokeRoleNotFound creates RevokeRoleNotFound with default headers values
func NewRevokeRoleNotFound() *RevokeRoleNotFound {

	return &RevokeRoleNotFound{}
}

// WriteResponse to the client
func (o *RevokeRoleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RevokeRoleInternalServerErrorCode is the HTTP code returned for type RevokeRoleInternalServerError
const RevokeRoleInternalServerErrorCode int = 500

/*
RevokeRoleInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response revokeRoleInternalServerError
*/
type RevokeRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewRevokeRoleInternalServerError creates RevokeRoleInternalServerError with default headers values
func NewRevokeRoleInternalServerError() *RevokeRoleInternalServerError {

	return &RevokeRoleInternalServerError{}
}

// WithPayload adds the payload to the revoke role internal server error response
func (o *RevokeRoleInternalServerError) WithPayload(payload *models.ErrorResponse) *RevokeRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke role internal server error response
func (o *RevokeRoleInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}