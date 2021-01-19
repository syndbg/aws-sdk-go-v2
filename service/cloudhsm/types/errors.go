// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Indicates that an internal error occurred.
type CloudHsmInternalException struct {
	Message *string

	Retryable bool
}

func (e *CloudHsmInternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudHsmInternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudHsmInternalException) ErrorCode() string             { return "CloudHsmInternalException" }
func (e *CloudHsmInternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Indicates that an exception occurred in the AWS CloudHSM service.
type CloudHsmServiceException struct {
	Message *string

	Retryable bool
}

func (e *CloudHsmServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudHsmServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudHsmServiceException) ErrorCode() string             { return "CloudHsmServiceException" }
func (e *CloudHsmServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the request parameters are not valid.
type InvalidRequestException struct {
	Message *string

	Retryable bool
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }