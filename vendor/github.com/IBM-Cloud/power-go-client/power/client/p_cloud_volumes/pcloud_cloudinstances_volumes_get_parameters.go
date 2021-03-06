// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudCloudinstancesVolumesGetParams creates a new PcloudCloudinstancesVolumesGetParams object
// with the default values initialized.
func NewPcloudCloudinstancesVolumesGetParams() *PcloudCloudinstancesVolumesGetParams {
	var ()
	return &PcloudCloudinstancesVolumesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesVolumesGetParamsWithTimeout creates a new PcloudCloudinstancesVolumesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesVolumesGetParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesGetParams {
	var ()
	return &PcloudCloudinstancesVolumesGetParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesVolumesGetParamsWithContext creates a new PcloudCloudinstancesVolumesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesVolumesGetParamsWithContext(ctx context.Context) *PcloudCloudinstancesVolumesGetParams {
	var ()
	return &PcloudCloudinstancesVolumesGetParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesVolumesGetParamsWithHTTPClient creates a new PcloudCloudinstancesVolumesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesVolumesGetParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesGetParams {
	var ()
	return &PcloudCloudinstancesVolumesGetParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesVolumesGetParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances volumes get operation typically these are written to a http.Request
*/
type PcloudCloudinstancesVolumesGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*VolumeID
	  Volume ID

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) WithContext(ctx context.Context) *PcloudCloudinstancesVolumesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesVolumesGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVolumeID adds the volumeID to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) WithVolumeID(volumeID string) *PcloudCloudinstancesVolumesGetParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the pcloud cloudinstances volumes get params
func (o *PcloudCloudinstancesVolumesGetParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesVolumesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
