// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2ImagesExportGetReader is a Reader for the PcloudV2ImagesExportGet structure.
type PcloudV2ImagesExportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2ImagesExportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudV2ImagesExportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudV2ImagesExportGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2ImagesExportGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2ImagesExportGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudV2ImagesExportGetOK creates a PcloudV2ImagesExportGetOK with default headers values
func NewPcloudV2ImagesExportGetOK() *PcloudV2ImagesExportGetOK {
	return &PcloudV2ImagesExportGetOK{}
}

/* PcloudV2ImagesExportGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudV2ImagesExportGetOK struct {
	Payload *models.Job
}

func (o *PcloudV2ImagesExportGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportGetOK  %+v", 200, o.Payload)
}
func (o *PcloudV2ImagesExportGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *PcloudV2ImagesExportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportGetUnauthorized creates a PcloudV2ImagesExportGetUnauthorized with default headers values
func NewPcloudV2ImagesExportGetUnauthorized() *PcloudV2ImagesExportGetUnauthorized {
	return &PcloudV2ImagesExportGetUnauthorized{}
}

/* PcloudV2ImagesExportGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2ImagesExportGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudV2ImagesExportGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudV2ImagesExportGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportGetNotFound creates a PcloudV2ImagesExportGetNotFound with default headers values
func NewPcloudV2ImagesExportGetNotFound() *PcloudV2ImagesExportGetNotFound {
	return &PcloudV2ImagesExportGetNotFound{}
}

/* PcloudV2ImagesExportGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2ImagesExportGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2ImagesExportGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudV2ImagesExportGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2ImagesExportGetInternalServerError creates a PcloudV2ImagesExportGetInternalServerError with default headers values
func NewPcloudV2ImagesExportGetInternalServerError() *PcloudV2ImagesExportGetInternalServerError {
	return &PcloudV2ImagesExportGetInternalServerError{}
}

/* PcloudV2ImagesExportGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2ImagesExportGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2ImagesExportGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/images/{image_id}/export][%d] pcloudV2ImagesExportGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudV2ImagesExportGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2ImagesExportGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}