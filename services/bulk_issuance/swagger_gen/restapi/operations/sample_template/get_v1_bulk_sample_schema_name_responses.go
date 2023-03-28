// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// GetV1BulkSampleSchemaNameOKCode is the HTTP code returned for type GetV1BulkSampleSchemaNameOK
const GetV1BulkSampleSchemaNameOKCode int = 200

/*GetV1BulkSampleSchemaNameOK OK

swagger:response getV1BulkSampleSchemaNameOK
*/
type GetV1BulkSampleSchemaNameOK struct {

	/*
	  In: Body
	*/
	Payload models.SampleTemplateResponse `json:"body,omitempty"`
}

// NewGetV1BulkSampleSchemaNameOK creates GetV1BulkSampleSchemaNameOK with default headers values
func NewGetV1BulkSampleSchemaNameOK() *GetV1BulkSampleSchemaNameOK {

	return &GetV1BulkSampleSchemaNameOK{}
}

// WithPayload adds the payload to the get v1 bulk sample schema name o k response
func (o *GetV1BulkSampleSchemaNameOK) WithPayload(payload models.SampleTemplateResponse) *GetV1BulkSampleSchemaNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 bulk sample schema name o k response
func (o *GetV1BulkSampleSchemaNameOK) SetPayload(payload models.SampleTemplateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1BulkSampleSchemaNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}