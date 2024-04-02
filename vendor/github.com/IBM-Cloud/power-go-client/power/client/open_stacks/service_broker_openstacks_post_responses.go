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

// ServiceBrokerOpenstacksPostReader is a Reader for the ServiceBrokerOpenstacksPost structure.
type ServiceBrokerOpenstacksPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerOpenstacksPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerOpenstacksPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewServiceBrokerOpenstacksPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerOpenstacksPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerOpenstacksPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerOpenstacksPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerOpenstacksPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewServiceBrokerOpenstacksPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceBrokerOpenstacksPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerOpenstacksPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /broker/v1/openstacks] serviceBroker.openstacks.post", response, response.Code())
	}
}

// NewServiceBrokerOpenstacksPostOK creates a ServiceBrokerOpenstacksPostOK with default headers values
func NewServiceBrokerOpenstacksPostOK() *ServiceBrokerOpenstacksPostOK {
	return &ServiceBrokerOpenstacksPostOK{}
}

/*
ServiceBrokerOpenstacksPostOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerOpenstacksPostOK struct {
	Payload *models.OpenStack
}

// IsSuccess returns true when this service broker openstacks post o k response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks post o k response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post o k response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post o k response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post o k response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker openstacks post o k response
func (o *ServiceBrokerOpenstacksPostOK) Code() int {
	return 200
}

func (o *ServiceBrokerOpenstacksPostOK) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostOK) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostOK) GetPayload() *models.OpenStack {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostCreated creates a ServiceBrokerOpenstacksPostCreated with default headers values
func NewServiceBrokerOpenstacksPostCreated() *ServiceBrokerOpenstacksPostCreated {
	return &ServiceBrokerOpenstacksPostCreated{}
}

/*
ServiceBrokerOpenstacksPostCreated describes a response with status code 201, with default header values.

Created
*/
type ServiceBrokerOpenstacksPostCreated struct {
	Payload *models.OpenStack
}

// IsSuccess returns true when this service broker openstacks post created response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker openstacks post created response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post created response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post created response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post created response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the service broker openstacks post created response
func (o *ServiceBrokerOpenstacksPostCreated) Code() int {
	return 201
}

func (o *ServiceBrokerOpenstacksPostCreated) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostCreated  %+v", 201, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostCreated) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostCreated  %+v", 201, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostCreated) GetPayload() *models.OpenStack {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostBadRequest creates a ServiceBrokerOpenstacksPostBadRequest with default headers values
func NewServiceBrokerOpenstacksPostBadRequest() *ServiceBrokerOpenstacksPostBadRequest {
	return &ServiceBrokerOpenstacksPostBadRequest{}
}

/*
ServiceBrokerOpenstacksPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerOpenstacksPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post bad request response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post bad request response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post bad request response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post bad request response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post bad request response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker openstacks post bad request response
func (o *ServiceBrokerOpenstacksPostBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerOpenstacksPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostBadRequest) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostUnauthorized creates a ServiceBrokerOpenstacksPostUnauthorized with default headers values
func NewServiceBrokerOpenstacksPostUnauthorized() *ServiceBrokerOpenstacksPostUnauthorized {
	return &ServiceBrokerOpenstacksPostUnauthorized{}
}

/*
ServiceBrokerOpenstacksPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerOpenstacksPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post unauthorized response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post unauthorized response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post unauthorized response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post unauthorized response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post unauthorized response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker openstacks post unauthorized response
func (o *ServiceBrokerOpenstacksPostUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerOpenstacksPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostForbidden creates a ServiceBrokerOpenstacksPostForbidden with default headers values
func NewServiceBrokerOpenstacksPostForbidden() *ServiceBrokerOpenstacksPostForbidden {
	return &ServiceBrokerOpenstacksPostForbidden{}
}

/*
ServiceBrokerOpenstacksPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerOpenstacksPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post forbidden response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post forbidden response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post forbidden response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post forbidden response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post forbidden response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker openstacks post forbidden response
func (o *ServiceBrokerOpenstacksPostForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerOpenstacksPostForbidden) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostForbidden) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostNotFound creates a ServiceBrokerOpenstacksPostNotFound with default headers values
func NewServiceBrokerOpenstacksPostNotFound() *ServiceBrokerOpenstacksPostNotFound {
	return &ServiceBrokerOpenstacksPostNotFound{}
}

/*
ServiceBrokerOpenstacksPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerOpenstacksPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post not found response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post not found response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post not found response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post not found response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post not found response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker openstacks post not found response
func (o *ServiceBrokerOpenstacksPostNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerOpenstacksPostNotFound) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostNotFound) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostConflict creates a ServiceBrokerOpenstacksPostConflict with default headers values
func NewServiceBrokerOpenstacksPostConflict() *ServiceBrokerOpenstacksPostConflict {
	return &ServiceBrokerOpenstacksPostConflict{}
}

/*
ServiceBrokerOpenstacksPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type ServiceBrokerOpenstacksPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post conflict response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post conflict response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post conflict response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post conflict response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post conflict response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the service broker openstacks post conflict response
func (o *ServiceBrokerOpenstacksPostConflict) Code() int {
	return 409
}

func (o *ServiceBrokerOpenstacksPostConflict) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostConflict  %+v", 409, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostConflict) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostConflict  %+v", 409, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostUnprocessableEntity creates a ServiceBrokerOpenstacksPostUnprocessableEntity with default headers values
func NewServiceBrokerOpenstacksPostUnprocessableEntity() *ServiceBrokerOpenstacksPostUnprocessableEntity {
	return &ServiceBrokerOpenstacksPostUnprocessableEntity{}
}

/*
ServiceBrokerOpenstacksPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceBrokerOpenstacksPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post unprocessable entity response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post unprocessable entity response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post unprocessable entity response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker openstacks post unprocessable entity response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker openstacks post unprocessable entity response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the service broker openstacks post unprocessable entity response
func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) Code() int {
	return 422
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerOpenstacksPostInternalServerError creates a ServiceBrokerOpenstacksPostInternalServerError with default headers values
func NewServiceBrokerOpenstacksPostInternalServerError() *ServiceBrokerOpenstacksPostInternalServerError {
	return &ServiceBrokerOpenstacksPostInternalServerError{}
}

/*
ServiceBrokerOpenstacksPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerOpenstacksPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker openstacks post internal server error response has a 2xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker openstacks post internal server error response has a 3xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker openstacks post internal server error response has a 4xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker openstacks post internal server error response has a 5xx status code
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker openstacks post internal server error response a status code equal to that given
func (o *ServiceBrokerOpenstacksPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker openstacks post internal server error response
func (o *ServiceBrokerOpenstacksPostInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /broker/v1/openstacks][%d] serviceBrokerOpenstacksPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerOpenstacksPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}