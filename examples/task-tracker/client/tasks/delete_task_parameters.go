package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewDeleteTaskParams creates a new DeleteTaskParams object
// with the default values initialized.
func NewDeleteTaskParams() *DeleteTaskParams {
	var ()
	return &DeleteTaskParams{}
}

/*DeleteTaskParams contains all the parameters to send to the API endpoint
for the delete task operation typically these are written to a http.Request
*/
type DeleteTaskParams struct {

	/*ID
	  The id of the item

	*/
	ID int64
}

// WithID adds the id to the delete task params
func (o *DeleteTaskParams) WithID(id int64) *DeleteTaskParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTaskParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}