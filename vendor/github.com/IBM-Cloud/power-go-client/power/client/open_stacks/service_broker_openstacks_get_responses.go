// Code generated by go-swagger; DO NOT EDIT.

package open_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerOpenstacksGetReader is a Reader for the ServiceBrokerOpenstacksGet structure.
type ServiceBrokerOpenstacksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerOpenstacksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerOpenstacksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerOpenstacksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/openstacks] serviceBroker.openstacks.get", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksGetOK creates a ServiceBrokerOpenstacksGetOK with default headers values
func NewServiceBrokerOpenstacksGetOK() *ServiceBrokerOpenstacksGetOK {
	return &ServiceBrokerOpenstacksGetOK{}
}

/*
ServiceBrokerOpenstacksGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksGetOK struct {
	Payload *models.OpenStacks
}

// IsSuccess returns true when this service broker openstacks get o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks get o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks get o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks get o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks get o k response
func (o *ServiceBrokerOpenstacksGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksGetOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetOK) GetPayload() *models.OpenStacks {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStacks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksGetBadRequest creates a ServiceBrokerOpenstacksGetBadRequest with default headers values
func NewServiceBrokerOpenstacksGetBadRequest() *ServiceBrokerOpenstacksGetBadRequest {
	return &ServiceBrokerOpenstacksGetBadRequest{}
}

/*
ServiceBrokerOpenstacksGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks get bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks get bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks get bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks get bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks get bad request response
func (o *ServiceBrokerOpenstacksGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksGetUnauthorized creates a ServiceBrokerOpenstacksGetUnauthorized with default headers values
func NewServiceBrokerOpenstacksGetUnauthorized() *ServiceBrokerOpenstacksGetUnauthorized {
	return &ServiceBrokerOpenstacksGetUnauthorized{}
}

/*
ServiceBrokerOpenstacksGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerOpenstacksGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks get unauthorized response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks get unauthorized response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get unauthorized response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks get unauthorized response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks get unauthorized response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker openstacks get unauthorized response
func (o *ServiceBrokerOpenstacksGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerOpenstacksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksGetForbidden creates a ServiceBrokerOpenstacksGetForbidden with default headers values
func NewServiceBrokerOpenstacksGetForbidden() *ServiceBrokerOpenstacksGetForbidden {
	return &ServiceBrokerOpenstacksGetForbidden{}
}

/*
ServiceBrokerOpenstacksGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerOpenstacksGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks get forbidden response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks get forbidden response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get forbidden response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks get forbidden response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks get forbidden response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker openstacks get forbidden response
func (o *ServiceBrokerOpenstacksGetForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerOpenstacksGetForbidden) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetForbidden) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksGetNotFound creates a ServiceBrokerOpenstacksGetNotFound with default headers values
func NewServiceBrokerOpenstacksGetNotFound() *ServiceBrokerOpenstacksGetNotFound {
	return &ServiceBrokerOpenstacksGetNotFound{}
}

/*
ServiceBrokerOpenstacksGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerOpenstacksGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks get not found response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks get not found response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get not found response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks get not found response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks get not found response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker openstacks get not found response
func (o *ServiceBrokerOpenstacksGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerOpenstacksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetNotFound) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksGetInternalServerError creates a ServiceBrokerOpenstacksGetInternalServerError with default headers values
func NewServiceBrokerOpenstacksGetInternalServerError() *ServiceBrokerOpenstacksGetInternalServerError {
	return &ServiceBrokerOpenstacksGetInternalServerError{}
}

/*
ServiceBrokerOpenstacksGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks get internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks get internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks get internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks get internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks get internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks get internal server error response
func (o *ServiceBrokerOpenstacksGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /broker/v1/openstacks][%d] serviceBrokerOpenstacksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}