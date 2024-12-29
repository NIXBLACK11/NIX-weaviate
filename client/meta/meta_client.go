// Code generated by go-swagger; DO NOT EDIT.

package meta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new meta API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for meta API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	MetaGet(params *MetaGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetaGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
MetaGet returns meta information of the current weaviate instance

Returns meta information about the server. Can be used to provide information to another Weaviate instance that wants to interact with the current instance.
*/
func (a *Client) MetaGet(params *MetaGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MetaGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetaGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "meta.get",
		Method:             "GET",
		PathPattern:        "/meta",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MetaGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MetaGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for meta.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
