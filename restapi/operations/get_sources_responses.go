package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetSourcesOK An array of data sources

swagger:response getSourcesOK
*/
type GetSourcesOK struct {

	// In: body
	Payload *models.Sources `json:"body,omitempty"`
}

// NewGetSourcesOK creates GetSourcesOK with default headers values
func NewGetSourcesOK() *GetSourcesOK {
	return &GetSourcesOK{}
}

// WithPayload adds the payload to the get sources o k response
func (o *GetSourcesOK) WithPayload(payload *models.Sources) *GetSourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources o k response
func (o *GetSourcesOK) SetPayload(payload *models.Sources) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesDefault Unexpected internal service error

swagger:response getSourcesDefault
*/
type GetSourcesDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesDefault creates GetSourcesDefault with default headers values
func NewGetSourcesDefault(code int) *GetSourcesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources default response
func (o *GetSourcesDefault) WithStatusCode(code int) *GetSourcesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources default response
func (o *GetSourcesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources default response
func (o *GetSourcesDefault) WithPayload(payload *models.Error) *GetSourcesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources default response
func (o *GetSourcesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
