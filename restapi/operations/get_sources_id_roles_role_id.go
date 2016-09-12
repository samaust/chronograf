package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSourcesIDRolesRoleIDHandlerFunc turns a function with the right signature into a get sources ID roles role ID handler
type GetSourcesIDRolesRoleIDHandlerFunc func(context.Context, GetSourcesIDRolesRoleIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSourcesIDRolesRoleIDHandlerFunc) Handle(ctx context.Context, params GetSourcesIDRolesRoleIDParams) middleware.Responder {
	return fn(ctx, params)
}

// GetSourcesIDRolesRoleIDHandler interface for that can handle valid get sources ID roles role ID params
type GetSourcesIDRolesRoleIDHandler interface {
	Handle(context.Context, GetSourcesIDRolesRoleIDParams) middleware.Responder
}

// NewGetSourcesIDRolesRoleID creates a new http.Handler for the get sources ID roles role ID operation
func NewGetSourcesIDRolesRoleID(ctx *middleware.Context, handler GetSourcesIDRolesRoleIDHandler) *GetSourcesIDRolesRoleID {
	return &GetSourcesIDRolesRoleID{Context: ctx, Handler: handler}
}

/*GetSourcesIDRolesRoleID swagger:route GET /sources/{id}/roles/{role_id} getSourcesIdRolesRoleId

Returns information about a specific role

Specific Role and its associated permissions.


*/
type GetSourcesIDRolesRoleID struct {
	Context *middleware.Context
	Handler GetSourcesIDRolesRoleIDHandler
}

func (o *GetSourcesIDRolesRoleID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetSourcesIDRolesRoleIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
