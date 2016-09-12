package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteSourcesIDHandlerFunc turns a function with the right signature into a delete sources ID handler
type DeleteSourcesIDHandlerFunc func(context.Context, DeleteSourcesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSourcesIDHandlerFunc) Handle(ctx context.Context, params DeleteSourcesIDParams) middleware.Responder {
	return fn(ctx, params)
}

// DeleteSourcesIDHandler interface for that can handle valid delete sources ID params
type DeleteSourcesIDHandler interface {
	Handle(context.Context, DeleteSourcesIDParams) middleware.Responder
}

// NewDeleteSourcesID creates a new http.Handler for the delete sources ID operation
func NewDeleteSourcesID(ctx *middleware.Context, handler DeleteSourcesIDHandler) *DeleteSourcesID {
	return &DeleteSourcesID{Context: ctx, Handler: handler}
}

/*DeleteSourcesID swagger:route DELETE /sources/{id} deleteSourcesId

This specific data source will be removed from the data store

*/
type DeleteSourcesID struct {
	Context *middleware.Context
	Handler DeleteSourcesIDHandler
}

func (o *DeleteSourcesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteSourcesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
