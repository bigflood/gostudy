// Code generated by go-swagger; DO NOT EDIT.

package admins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddInventoryHandlerFunc turns a function with the right signature into a add inventory handler
type AddInventoryHandlerFunc func(AddInventoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddInventoryHandlerFunc) Handle(params AddInventoryParams) middleware.Responder {
	return fn(params)
}

// AddInventoryHandler interface for that can handle valid add inventory params
type AddInventoryHandler interface {
	Handle(AddInventoryParams) middleware.Responder
}

// NewAddInventory creates a new http.Handler for the add inventory operation
func NewAddInventory(ctx *middleware.Context, handler AddInventoryHandler) *AddInventory {
	return &AddInventory{Context: ctx, Handler: handler}
}

/*AddInventory swagger:route POST /inventory admins addInventory

adds an inventory item

Adds an item to the system

*/
type AddInventory struct {
	Context *middleware.Context
	Handler AddInventoryHandler
}

func (o *AddInventory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddInventoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
