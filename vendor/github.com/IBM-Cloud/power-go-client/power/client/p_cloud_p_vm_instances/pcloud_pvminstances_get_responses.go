// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesGetReader is a Reader for the PcloudPvminstancesGet structure.
type PcloudPvminstancesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}] pcloud.pvminstances.get", response, response.Code())
	}
}

// NewPcloudPvminstancesGetOK creates a PcloudPvminstancesGetOK with default headers values
func NewPcloudPvminstancesGetOK() *PcloudPvminstancesGetOK {
	return &PcloudPvminstancesGetOK{}
}

/*
PcloudPvminstancesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesGetOK struct {
	Payload *models.PVMInstance
}

// IsSuccess returns true when this pcloud pvminstances get o k response has a 2xx status code
func (o *PcloudPvminstancesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances get o k response has a 3xx status code
func (o *PcloudPvminstancesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get o k response has a 4xx status code
func (o *PcloudPvminstancesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances get o k response has a 5xx status code
func (o *PcloudPvminstancesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances get o k response a status code equal to that given
func (o *PcloudPvminstancesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances get o k response
func (o *PcloudPvminstancesGetOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesGetOK) GetPayload() *models.PVMInstance {
	return o.Payload
}

func (o *PcloudPvminstancesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetBadRequest creates a PcloudPvminstancesGetBadRequest with default headers values
func NewPcloudPvminstancesGetBadRequest() *PcloudPvminstancesGetBadRequest {
	return &PcloudPvminstancesGetBadRequest{}
}

/*
PcloudPvminstancesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances get bad request response has a 2xx status code
func (o *PcloudPvminstancesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances get bad request response has a 3xx status code
func (o *PcloudPvminstancesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get bad request response has a 4xx status code
func (o *PcloudPvminstancesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances get bad request response has a 5xx status code
func (o *PcloudPvminstancesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances get bad request response a status code equal to that given
func (o *PcloudPvminstancesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances get bad request response
func (o *PcloudPvminstancesGetBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetUnauthorized creates a PcloudPvminstancesGetUnauthorized with default headers values
func NewPcloudPvminstancesGetUnauthorized() *PcloudPvminstancesGetUnauthorized {
	return &PcloudPvminstancesGetUnauthorized{}
}

/*
PcloudPvminstancesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances get unauthorized response has a 2xx status code
func (o *PcloudPvminstancesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances get unauthorized response has a 3xx status code
func (o *PcloudPvminstancesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get unauthorized response has a 4xx status code
func (o *PcloudPvminstancesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances get unauthorized response has a 5xx status code
func (o *PcloudPvminstancesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances get unauthorized response a status code equal to that given
func (o *PcloudPvminstancesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances get unauthorized response
func (o *PcloudPvminstancesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetForbidden creates a PcloudPvminstancesGetForbidden with default headers values
func NewPcloudPvminstancesGetForbidden() *PcloudPvminstancesGetForbidden {
	return &PcloudPvminstancesGetForbidden{}
}

/*
PcloudPvminstancesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances get forbidden response has a 2xx status code
func (o *PcloudPvminstancesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances get forbidden response has a 3xx status code
func (o *PcloudPvminstancesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get forbidden response has a 4xx status code
func (o *PcloudPvminstancesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances get forbidden response has a 5xx status code
func (o *PcloudPvminstancesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances get forbidden response a status code equal to that given
func (o *PcloudPvminstancesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances get forbidden response
func (o *PcloudPvminstancesGetForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetNotFound creates a PcloudPvminstancesGetNotFound with default headers values
func NewPcloudPvminstancesGetNotFound() *PcloudPvminstancesGetNotFound {
	return &PcloudPvminstancesGetNotFound{}
}

/*
PcloudPvminstancesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances get not found response has a 2xx status code
func (o *PcloudPvminstancesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances get not found response has a 3xx status code
func (o *PcloudPvminstancesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get not found response has a 4xx status code
func (o *PcloudPvminstancesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances get not found response has a 5xx status code
func (o *PcloudPvminstancesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances get not found response a status code equal to that given
func (o *PcloudPvminstancesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances get not found response
func (o *PcloudPvminstancesGetNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesGetInternalServerError creates a PcloudPvminstancesGetInternalServerError with default headers values
func NewPcloudPvminstancesGetInternalServerError() *PcloudPvminstancesGetInternalServerError {
	return &PcloudPvminstancesGetInternalServerError{}
}

/*
PcloudPvminstancesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances get internal server error response has a 2xx status code
func (o *PcloudPvminstancesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances get internal server error response has a 3xx status code
func (o *PcloudPvminstancesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances get internal server error response has a 4xx status code
func (o *PcloudPvminstancesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances get internal server error response has a 5xx status code
func (o *PcloudPvminstancesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances get internal server error response a status code equal to that given
func (o *PcloudPvminstancesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances get internal server error response
func (o *PcloudPvminstancesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}][%d] pcloudPvminstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}