// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchInventoryHandlerFunc turns a function with the right signature into a search inventory handler
type SearchInventoryHandlerFunc func(SearchInventoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchInventoryHandlerFunc) Handle(params SearchInventoryParams) middleware.Responder {
	return fn(params)
}

// SearchInventoryHandler interface for that can handle valid search inventory params
type SearchInventoryHandler interface {
	Handle(SearchInventoryParams) middleware.Responder
}

// NewSearchInventory creates a new http.Handler for the search inventory operation
func NewSearchInventory(ctx *middleware.Context, handler SearchInventoryHandler) *SearchInventory {
	return &SearchInventory{Context: ctx, Handler: handler}
}

/*SearchInventory swagger:route GET /inventory developers searchInventory

searches inventory

By passing in the appropriate options, you can search for
available inventory in the system


*/
type SearchInventory struct {
	Context *middleware.Context
	Handler SearchInventoryHandler
}

func (o *SearchInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchInventoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
