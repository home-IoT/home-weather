// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSensorData Returns the data of a particular sensor
*/
func (a *Client) GetSensorData(params *GetSensorDataParams) (*GetSensorDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSensorData",
		Method:             "GET",
		PathPattern:        "/sensors/{sensorId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSensorDataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSensorDataOK), nil

}

/*
GetSensorDataRaw Returns the data of a particular sensor in simple JSON
*/
func (a *Client) GetSensorDataRaw(params *GetSensorDataRawParams) (*GetSensorDataRawOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorDataRawParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSensorDataRaw",
		Method:             "GET",
		PathPattern:        "/sensors/{sensorId}/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSensorDataRawReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSensorDataRawOK), nil

}

/*
GetSensorsList Returns the list of sensors
*/
func (a *Client) GetSensorsList(params *GetSensorsListParams) (*GetSensorsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSensorsList",
		Method:             "GET",
		PathPattern:        "/sensors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSensorsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSensorsListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
