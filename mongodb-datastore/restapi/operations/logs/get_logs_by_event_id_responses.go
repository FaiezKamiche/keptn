// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetLogsByEventIDOKCode is the HTTP code returned for type GetLogsByEventIDOK
const GetLogsByEventIDOKCode int = 200

/*GetLogsByEventIDOK ok

swagger:response getLogsByEventIdOK
*/
type GetLogsByEventIDOK struct {

	/*
	  In: Body
	*/
	Payload []*GetLogsByEventIDOKBodyItems0 `json:"body,omitempty"`
}

// NewGetLogsByEventIDOK creates GetLogsByEventIDOK with default headers values
func NewGetLogsByEventIDOK() *GetLogsByEventIDOK {

	return &GetLogsByEventIDOK{}
}

// WithPayload adds the payload to the get logs by event Id o k response
func (o *GetLogsByEventIDOK) WithPayload(payload []*GetLogsByEventIDOKBodyItems0) *GetLogsByEventIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs by event Id o k response
func (o *GetLogsByEventIDOK) SetPayload(payload []*GetLogsByEventIDOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsByEventIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetLogsByEventIDOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetLogsByEventIDDefault error

swagger:response getLogsByEventIdDefault
*/
type GetLogsByEventIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *GetLogsByEventIDDefaultBody `json:"body,omitempty"`
}

// NewGetLogsByEventIDDefault creates GetLogsByEventIDDefault with default headers values
func NewGetLogsByEventIDDefault(code int) *GetLogsByEventIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogsByEventIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get logs by event Id default response
func (o *GetLogsByEventIDDefault) WithStatusCode(code int) *GetLogsByEventIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get logs by event Id default response
func (o *GetLogsByEventIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get logs by event Id default response
func (o *GetLogsByEventIDDefault) WithPayload(payload *GetLogsByEventIDDefaultBody) *GetLogsByEventIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logs by event Id default response
func (o *GetLogsByEventIDDefault) SetPayload(payload *GetLogsByEventIDDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogsByEventIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}