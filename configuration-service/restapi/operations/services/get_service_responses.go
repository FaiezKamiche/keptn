// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/configuration-service/models"
)

// GetServiceOKCode is the HTTP code returned for type GetServiceOK
const GetServiceOKCode int = 200

/*GetServiceOK Success

swagger:response getServiceOK
*/
type GetServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ExpandedServiceWithStageInfo `json:"body,omitempty"`
}

// NewGetServiceOK creates GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {

	return &GetServiceOK{}
}

// WithPayload adds the payload to the get service o k response
func (o *GetServiceOK) WithPayload(payload *models.ExpandedServiceWithStageInfo) *GetServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service o k response
func (o *GetServiceOK) SetPayload(payload *models.ExpandedServiceWithStageInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceNotFoundCode is the HTTP code returned for type GetServiceNotFound
const GetServiceNotFoundCode int = 404

/*GetServiceNotFound Failed. Service could not be found.

swagger:response getServiceNotFound
*/
type GetServiceNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServiceNotFound creates GetServiceNotFound with default headers values
func NewGetServiceNotFound() *GetServiceNotFound {

	return &GetServiceNotFound{}
}

// WithPayload adds the payload to the get service not found response
func (o *GetServiceNotFound) WithPayload(payload *models.Error) *GetServiceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service not found response
func (o *GetServiceNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServiceDefault Error

swagger:response getServiceDefault
*/
type GetServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServiceDefault creates GetServiceDefault with default headers values
func NewGetServiceDefault(code int) *GetServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get service default response
func (o *GetServiceDefault) WithStatusCode(code int) *GetServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get service default response
func (o *GetServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get service default response
func (o *GetServiceDefault) WithPayload(payload *models.Error) *GetServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service default response
func (o *GetServiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
