// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// GetRolesOKCode is the HTTP code returned for type GetRolesOK
const GetRolesOKCode int = 200

/*
GetRolesOK Successful response.

swagger:response getRolesOK
*/
type GetRolesOK struct {

	/*
	  In: Body
	*/
	Payload models.RolesListResponse `json:"body,omitempty"`
}

// NewGetRolesOK creates GetRolesOK with default headers values
func NewGetRolesOK() *GetRolesOK {

	return &GetRolesOK{}
}

// WithPayload adds the payload to the get roles o k response
func (o *GetRolesOK) WithPayload(payload models.RolesListResponse) *GetRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles o k response
func (o *GetRolesOK) SetPayload(payload models.RolesListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.RolesListResponse{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRolesBadRequestCode is the HTTP code returned for type GetRolesBadRequest
const GetRolesBadRequestCode int = 400

/*
GetRolesBadRequest Malformed request.

swagger:response getRolesBadRequest
*/
type GetRolesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesBadRequest creates GetRolesBadRequest with default headers values
func NewGetRolesBadRequest() *GetRolesBadRequest {

	return &GetRolesBadRequest{}
}

// WithPayload adds the payload to the get roles bad request response
func (o *GetRolesBadRequest) WithPayload(payload *models.ErrorResponse) *GetRolesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles bad request response
func (o *GetRolesBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesUnauthorizedCode is the HTTP code returned for type GetRolesUnauthorized
const GetRolesUnauthorizedCode int = 401

/*
GetRolesUnauthorized Unauthorized or invalid credentials.

swagger:response getRolesUnauthorized
*/
type GetRolesUnauthorized struct {
}

// NewGetRolesUnauthorized creates GetRolesUnauthorized with default headers values
func NewGetRolesUnauthorized() *GetRolesUnauthorized {

	return &GetRolesUnauthorized{}
}

// WriteResponse to the client
func (o *GetRolesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetRolesForbiddenCode is the HTTP code returned for type GetRolesForbidden
const GetRolesForbiddenCode int = 403

/*
GetRolesForbidden Forbidden

swagger:response getRolesForbidden
*/
type GetRolesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesForbidden creates GetRolesForbidden with default headers values
func NewGetRolesForbidden() *GetRolesForbidden {

	return &GetRolesForbidden{}
}

// WithPayload adds the payload to the get roles forbidden response
func (o *GetRolesForbidden) WithPayload(payload *models.ErrorResponse) *GetRolesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles forbidden response
func (o *GetRolesForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesInternalServerErrorCode is the HTTP code returned for type GetRolesInternalServerError
const GetRolesInternalServerErrorCode int = 500

/*
GetRolesInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response getRolesInternalServerError
*/
type GetRolesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesInternalServerError creates GetRolesInternalServerError with default headers values
func NewGetRolesInternalServerError() *GetRolesInternalServerError {

	return &GetRolesInternalServerError{}
}

// WithPayload adds the payload to the get roles internal server error response
func (o *GetRolesInternalServerError) WithPayload(payload *models.ErrorResponse) *GetRolesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles internal server error response
func (o *GetRolesInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}