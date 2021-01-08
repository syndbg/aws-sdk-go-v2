// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the detector version’s status. You can perform the following promotions
// or demotions using UpdateDetectorVersionStatus: DRAFT to ACTIVE, ACTIVE to
// INACTIVE, and INACTIVE to ACTIVE.
func (c *Client) UpdateDetectorVersionStatus(ctx context.Context, params *UpdateDetectorVersionStatusInput, optFns ...func(*Options)) (*UpdateDetectorVersionStatusOutput, error) {
	if params == nil {
		params = &UpdateDetectorVersionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDetectorVersionStatus", params, optFns, addOperationUpdateDetectorVersionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDetectorVersionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorVersionStatusInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	// The new status.
	//
	// This member is required.
	Status types.DetectorVersionStatus
}

type UpdateDetectorVersionStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDetectorVersionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDetectorVersionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDetectorVersionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateDetectorVersionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetectorVersionStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDetectorVersionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateDetectorVersionStatus",
	}
}