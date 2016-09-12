package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetSourcesIDUsersUserIDExplorationsOK Data Explorations saved sessions for user are returned.

swagger:response getSourcesIdUsersUserIdExplorationsOK
*/
type GetSourcesIDUsersUserIDExplorationsOK struct {

	// In: body
	Payload *models.Explorations `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDExplorationsOK creates GetSourcesIDUsersUserIDExplorationsOK with default headers values
func NewGetSourcesIDUsersUserIDExplorationsOK() *GetSourcesIDUsersUserIDExplorationsOK {
	return &GetSourcesIDUsersUserIDExplorationsOK{}
}

// WithPayload adds the payload to the get sources Id users user Id explorations o k response
func (o *GetSourcesIDUsersUserIDExplorationsOK) WithPayload(payload *models.Explorations) *GetSourcesIDUsersUserIDExplorationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id users user Id explorations o k response
func (o *GetSourcesIDUsersUserIDExplorationsOK) SetPayload(payload *models.Explorations) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDExplorationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDUsersUserIDExplorationsNotFound Data source id or user does not exist.

swagger:response getSourcesIdUsersUserIdExplorationsNotFound
*/
type GetSourcesIDUsersUserIDExplorationsNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDExplorationsNotFound creates GetSourcesIDUsersUserIDExplorationsNotFound with default headers values
func NewGetSourcesIDUsersUserIDExplorationsNotFound() *GetSourcesIDUsersUserIDExplorationsNotFound {
	return &GetSourcesIDUsersUserIDExplorationsNotFound{}
}

// WithPayload adds the payload to the get sources Id users user Id explorations not found response
func (o *GetSourcesIDUsersUserIDExplorationsNotFound) WithPayload(payload *models.Error) *GetSourcesIDUsersUserIDExplorationsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id users user Id explorations not found response
func (o *GetSourcesIDUsersUserIDExplorationsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDExplorationsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDUsersUserIDExplorationsDefault Unexpected internal service error

swagger:response getSourcesIdUsersUserIdExplorationsDefault
*/
type GetSourcesIDUsersUserIDExplorationsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDUsersUserIDExplorationsDefault creates GetSourcesIDUsersUserIDExplorationsDefault with default headers values
func NewGetSourcesIDUsersUserIDExplorationsDefault(code int) *GetSourcesIDUsersUserIDExplorationsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesIDUsersUserIDExplorationsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources ID users user ID explorations default response
func (o *GetSourcesIDUsersUserIDExplorationsDefault) WithStatusCode(code int) *GetSourcesIDUsersUserIDExplorationsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources ID users user ID explorations default response
func (o *GetSourcesIDUsersUserIDExplorationsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources ID users user ID explorations default response
func (o *GetSourcesIDUsersUserIDExplorationsDefault) WithPayload(payload *models.Error) *GetSourcesIDUsersUserIDExplorationsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources ID users user ID explorations default response
func (o *GetSourcesIDUsersUserIDExplorationsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersUserIDExplorationsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
