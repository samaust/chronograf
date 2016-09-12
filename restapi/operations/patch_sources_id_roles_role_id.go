package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchSourcesIDRolesRoleIDHandlerFunc turns a function with the right signature into a patch sources ID roles role ID handler
type PatchSourcesIDRolesRoleIDHandlerFunc func(context.Context, PatchSourcesIDRolesRoleIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchSourcesIDRolesRoleIDHandlerFunc) Handle(ctx context.Context, params PatchSourcesIDRolesRoleIDParams) middleware.Responder {
	return fn(ctx, params)
}

// PatchSourcesIDRolesRoleIDHandler interface for that can handle valid patch sources ID roles role ID params
type PatchSourcesIDRolesRoleIDHandler interface {
	Handle(context.Context, PatchSourcesIDRolesRoleIDParams) middleware.Responder
}

// NewPatchSourcesIDRolesRoleID creates a new http.Handler for the patch sources ID roles role ID operation
func NewPatchSourcesIDRolesRoleID(ctx *middleware.Context, handler PatchSourcesIDRolesRoleIDHandler) *PatchSourcesIDRolesRoleID {
	return &PatchSourcesIDRolesRoleID{Context: ctx, Handler: handler}
}

/*PatchSourcesIDRolesRoleID swagger:route PATCH /sources/{id}/roles/{role_id} patchSourcesIdRolesRoleId

Update role configuration

*/
type PatchSourcesIDRolesRoleID struct {
	Context *middleware.Context
	Handler PatchSourcesIDRolesRoleIDHandler
}

func (o *PatchSourcesIDRolesRoleID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchSourcesIDRolesRoleIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
