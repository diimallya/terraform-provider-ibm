// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstancesParams creates a new GetInstancesParams object
// with the default values initialized.
func NewGetInstancesParams() *GetInstancesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetInstancesParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesParamsWithTimeout creates a new GetInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesParamsWithTimeout(timeout time.Duration) *GetInstancesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetInstancesParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetInstancesParamsWithContext creates a new GetInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesParamsWithContext(ctx context.Context) *GetInstancesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetInstancesParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetInstancesParamsWithHTTPClient creates a new GetInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesParamsWithHTTPClient(client *http.Client) *GetInstancesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetInstancesParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetInstancesParams contains all the parameters to send to the API endpoint
for the get instances operation typically these are written to a http.Request
*/
type GetInstancesParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*NetworkInterfacesSubnetCrn
	  Filters the collection to instances on the subnet of the specified CRN

	*/
	NetworkInterfacesSubnetCrn *string
	/*NetworkInterfacesSubnetID
	  Filters the collection to instances on the subnet of the specified identifier

	*/
	NetworkInterfacesSubnetID *string
	/*NetworkInterfacesSubnetName
	  Filters the collection to instances on the subnet of the specified name

	*/
	NetworkInterfacesSubnetName *string
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcCrn
	  Filters the collection to resources within the VPC of the specified CRN

	*/
	VpcCrn *string
	/*VpcID
	  Filters the collection to resources within the VPC of the specified identifier

	*/
	VpcID *string
	/*VpcName
	  Filters the collection to resources within the VPC of the specified name

	*/
	VpcName *string
	/*ZoneName
	  Filters the collection to resources within the specified zone

	*/
	ZoneName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances params
func (o *GetInstancesParams) WithTimeout(timeout time.Duration) *GetInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances params
func (o *GetInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances params
func (o *GetInstancesParams) WithContext(ctx context.Context) *GetInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances params
func (o *GetInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances params
func (o *GetInstancesParams) WithHTTPClient(client *http.Client) *GetInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances params
func (o *GetInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get instances params
func (o *GetInstancesParams) WithGeneration(generation int64) *GetInstancesParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get instances params
func (o *GetInstancesParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get instances params
func (o *GetInstancesParams) WithLimit(limit *int32) *GetInstancesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get instances params
func (o *GetInstancesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNetworkInterfacesSubnetCrn adds the networkInterfacesSubnetCrn to the get instances params
func (o *GetInstancesParams) WithNetworkInterfacesSubnetCrn(networkInterfacesSubnetCrn *string) *GetInstancesParams {
	o.SetNetworkInterfacesSubnetCrn(networkInterfacesSubnetCrn)
	return o
}

// SetNetworkInterfacesSubnetCrn adds the networkInterfacesSubnetCrn to the get instances params
func (o *GetInstancesParams) SetNetworkInterfacesSubnetCrn(networkInterfacesSubnetCrn *string) {
	o.NetworkInterfacesSubnetCrn = networkInterfacesSubnetCrn
}

// WithNetworkInterfacesSubnetID adds the networkInterfacesSubnetID to the get instances params
func (o *GetInstancesParams) WithNetworkInterfacesSubnetID(networkInterfacesSubnetID *string) *GetInstancesParams {
	o.SetNetworkInterfacesSubnetID(networkInterfacesSubnetID)
	return o
}

// SetNetworkInterfacesSubnetID adds the networkInterfacesSubnetId to the get instances params
func (o *GetInstancesParams) SetNetworkInterfacesSubnetID(networkInterfacesSubnetID *string) {
	o.NetworkInterfacesSubnetID = networkInterfacesSubnetID
}

// WithNetworkInterfacesSubnetName adds the networkInterfacesSubnetName to the get instances params
func (o *GetInstancesParams) WithNetworkInterfacesSubnetName(networkInterfacesSubnetName *string) *GetInstancesParams {
	o.SetNetworkInterfacesSubnetName(networkInterfacesSubnetName)
	return o
}

// SetNetworkInterfacesSubnetName adds the networkInterfacesSubnetName to the get instances params
func (o *GetInstancesParams) SetNetworkInterfacesSubnetName(networkInterfacesSubnetName *string) {
	o.NetworkInterfacesSubnetName = networkInterfacesSubnetName
}

// WithResourceGroupID adds the resourceGroupID to the get instances params
func (o *GetInstancesParams) WithResourceGroupID(resourceGroupID *string) *GetInstancesParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get instances params
func (o *GetInstancesParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get instances params
func (o *GetInstancesParams) WithStart(start *string) *GetInstancesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get instances params
func (o *GetInstancesParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get instances params
func (o *GetInstancesParams) WithVersion(version string) *GetInstancesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances params
func (o *GetInstancesParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcCrn adds the vpcCrn to the get instances params
func (o *GetInstancesParams) WithVpcCrn(vpcCrn *string) *GetInstancesParams {
	o.SetVpcCrn(vpcCrn)
	return o
}

// SetVpcCrn adds the vpcCrn to the get instances params
func (o *GetInstancesParams) SetVpcCrn(vpcCrn *string) {
	o.VpcCrn = vpcCrn
}

// WithVpcID adds the vpcID to the get instances params
func (o *GetInstancesParams) WithVpcID(vpcID *string) *GetInstancesParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get instances params
func (o *GetInstancesParams) SetVpcID(vpcID *string) {
	o.VpcID = vpcID
}

// WithVpcName adds the vpcName to the get instances params
func (o *GetInstancesParams) WithVpcName(vpcName *string) *GetInstancesParams {
	o.SetVpcName(vpcName)
	return o
}

// SetVpcName adds the vpcName to the get instances params
func (o *GetInstancesParams) SetVpcName(vpcName *string) {
	o.VpcName = vpcName
}

// WithZoneName adds the zoneName to the get instances params
func (o *GetInstancesParams) WithZoneName(zoneName *string) *GetInstancesParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the get instances params
func (o *GetInstancesParams) SetZoneName(zoneName *string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.NetworkInterfacesSubnetCrn != nil {

		// query param network_interfaces.subnet.crn
		var qrNetworkInterfacesSubnetCrn string
		if o.NetworkInterfacesSubnetCrn != nil {
			qrNetworkInterfacesSubnetCrn = *o.NetworkInterfacesSubnetCrn
		}
		qNetworkInterfacesSubnetCrn := qrNetworkInterfacesSubnetCrn
		if qNetworkInterfacesSubnetCrn != "" {
			if err := r.SetQueryParam("network_interfaces.subnet.crn", qNetworkInterfacesSubnetCrn); err != nil {
				return err
			}
		}

	}

	if o.NetworkInterfacesSubnetID != nil {

		// query param network_interfaces.subnet.id
		var qrNetworkInterfacesSubnetID string
		if o.NetworkInterfacesSubnetID != nil {
			qrNetworkInterfacesSubnetID = *o.NetworkInterfacesSubnetID
		}
		qNetworkInterfacesSubnetID := qrNetworkInterfacesSubnetID
		if qNetworkInterfacesSubnetID != "" {
			if err := r.SetQueryParam("network_interfaces.subnet.id", qNetworkInterfacesSubnetID); err != nil {
				return err
			}
		}

	}

	if o.NetworkInterfacesSubnetName != nil {

		// query param network_interfaces.subnet.name
		var qrNetworkInterfacesSubnetName string
		if o.NetworkInterfacesSubnetName != nil {
			qrNetworkInterfacesSubnetName = *o.NetworkInterfacesSubnetName
		}
		qNetworkInterfacesSubnetName := qrNetworkInterfacesSubnetName
		if qNetworkInterfacesSubnetName != "" {
			if err := r.SetQueryParam("network_interfaces.subnet.name", qNetworkInterfacesSubnetName); err != nil {
				return err
			}
		}

	}

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if o.VpcCrn != nil {

		// query param vpc.crn
		var qrVpcCrn string
		if o.VpcCrn != nil {
			qrVpcCrn = *o.VpcCrn
		}
		qVpcCrn := qrVpcCrn
		if qVpcCrn != "" {
			if err := r.SetQueryParam("vpc.crn", qVpcCrn); err != nil {
				return err
			}
		}

	}

	if o.VpcID != nil {

		// query param vpc.id
		var qrVpcID string
		if o.VpcID != nil {
			qrVpcID = *o.VpcID
		}
		qVpcID := qrVpcID
		if qVpcID != "" {
			if err := r.SetQueryParam("vpc.id", qVpcID); err != nil {
				return err
			}
		}

	}

	if o.VpcName != nil {

		// query param vpc.name
		var qrVpcName string
		if o.VpcName != nil {
			qrVpcName = *o.VpcName
		}
		qVpcName := qrVpcName
		if qVpcName != "" {
			if err := r.SetQueryParam("vpc.name", qVpcName); err != nil {
				return err
			}
		}

	}

	if o.ZoneName != nil {

		// query param zone.name
		var qrZoneName string
		if o.ZoneName != nil {
			qrZoneName = *o.ZoneName
		}
		qZoneName := qrZoneName
		if qZoneName != "" {
			if err := r.SetQueryParam("zone.name", qZoneName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
