// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/bigflood/gostudy/swagger/models"
)

// SearchInventoryOKCode is the HTTP code returned for type SearchInventoryOK
const SearchInventoryOKCode int = 200

/*SearchInventoryOK search results matching criteria

swagger:response searchInventoryOK
*/
type SearchInventoryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.InventoryItem `json:"body,omitempty"`
}

// NewSearchInventoryOK creates SearchInventoryOK with default headers values
func NewSearchInventoryOK() *SearchInventoryOK {

	return &SearchInventoryOK{}
}

// WithPayload adds the payload to the search inventory o k response
func (o *SearchInventoryOK) WithPayload(payload []*models.InventoryItem) *SearchInventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search inventory o k response
func (o *SearchInventoryOK) SetPayload(payload []*models.InventoryItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchInventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.InventoryItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// SearchInventoryBadRequestCode is the HTTP code returned for type SearchInventoryBadRequest
const SearchInventoryBadRequestCode int = 400

/*SearchInventoryBadRequest bad input parameter

swagger:response searchInventoryBadRequest
*/
type SearchInventoryBadRequest struct {
}

// NewSearchInventoryBadRequest creates SearchInventoryBadRequest with default headers values
func NewSearchInventoryBadRequest() *SearchInventoryBadRequest {

	return &SearchInventoryBadRequest{}
}

// WriteResponse to the client
func (o *SearchInventoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
