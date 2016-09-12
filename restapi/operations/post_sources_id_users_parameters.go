package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/influxdata/mrfusion/models"
)

// NewPostSourcesIDUsersParams creates a new PostSourcesIDUsersParams object
// with the default values initialized.
func NewPostSourcesIDUsersParams() PostSourcesIDUsersParams {
	var ()
	return PostSourcesIDUsersParams{}
}

// PostSourcesIDUsersParams contains all the bound params for the post sources ID users operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostSourcesIDUsers
type PostSourcesIDUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of the data source
	  Required: true
	  In: path
	*/
	ID string
	/*Configuration options for new user
	  In: body
	*/
	User *models.User
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostSourcesIDUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.User
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("user", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.User = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSourcesIDUsersParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
