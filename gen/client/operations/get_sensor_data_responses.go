// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/home-IoT/home-weather/gen/models"
)

// GetSensorDataReader is a Reader for the GetSensorData structure.
type GetSensorDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSensorDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSensorDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetSensorDataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSensorDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSensorDataOK creates a GetSensorDataOK with default headers values
func NewGetSensorDataOK() *GetSensorDataOK {
	return &GetSensorDataOK{}
}

/*GetSensorDataOK handles this case with default header values.

Success
*/
type GetSensorDataOK struct {
	Payload *models.SensorResponse
}

func (o *GetSensorDataOK) Error() string {
	return fmt.Sprintf("[GET /sensors/{sensorId}][%d] getSensorDataOK  %+v", 200, o.Payload)
}

func (o *GetSensorDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SensorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorDataNotFound creates a GetSensorDataNotFound with default headers values
func NewGetSensorDataNotFound() *GetSensorDataNotFound {
	return &GetSensorDataNotFound{}
}

/*GetSensorDataNotFound handles this case with default header values.

Sensor not found.
*/
type GetSensorDataNotFound struct {
}

func (o *GetSensorDataNotFound) Error() string {
	return fmt.Sprintf("[GET /sensors/{sensorId}][%d] getSensorDataNotFound ", 404)
}

func (o *GetSensorDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSensorDataGatewayTimeout creates a GetSensorDataGatewayTimeout with default headers values
func NewGetSensorDataGatewayTimeout() *GetSensorDataGatewayTimeout {
	return &GetSensorDataGatewayTimeout{}
}

/*GetSensorDataGatewayTimeout handles this case with default header values.

Sensor is not available.
*/
type GetSensorDataGatewayTimeout struct {
}

func (o *GetSensorDataGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /sensors/{sensorId}][%d] getSensorDataGatewayTimeout ", 504)
}

func (o *GetSensorDataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSensorDataDefault creates a GetSensorDataDefault with default headers values
func NewGetSensorDataDefault(code int) *GetSensorDataDefault {
	return &GetSensorDataDefault{
		_statusCode: code,
	}
}

/*GetSensorDataDefault handles this case with default header values.

Error
*/
type GetSensorDataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get sensor data default response
func (o *GetSensorDataDefault) Code() int {
	return o._statusCode
}

func (o *GetSensorDataDefault) Error() string {
	return fmt.Sprintf("[GET /sensors/{sensorId}][%d] getSensorData default  %+v", o._statusCode, o.Payload)
}

func (o *GetSensorDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
