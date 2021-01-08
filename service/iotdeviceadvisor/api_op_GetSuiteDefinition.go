// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Device Advisor test suite.
func (c *Client) GetSuiteDefinition(ctx context.Context, params *GetSuiteDefinitionInput, optFns ...func(*Options)) (*GetSuiteDefinitionOutput, error) {
	if params == nil {
		params = &GetSuiteDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSuiteDefinition", params, optFns, addOperationGetSuiteDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSuiteDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSuiteDefinitionInput struct {

	// Requests suite definition Id with GetSuiteDefinition API call.
	//
	// This member is required.
	SuiteDefinitionId *string

	// Requests the suite definition version of a test suite.
	SuiteDefinitionVersion *string
}

type GetSuiteDefinitionOutput struct {

	// Gets the timestamp of the time suite was created with GetSuiteDefinition API
	// call.
	CreatedAt *time.Time

	// Gets the timestamp of the time suite was modified with GetSuiteDefinition API
	// call.
	LastModifiedAt *time.Time

	// Gets latest suite definition version with GetSuiteDefinition API call.
	LatestVersion *string

	// The ARN of the suite definition.
	SuiteDefinitionArn *string

	// Gets the suite configuration with GetSuiteDefinition API call.
	SuiteDefinitionConfiguration *types.SuiteDefinitionConfiguration

	// Gets suite definition Id with GetSuiteDefinition API call.
	SuiteDefinitionId *string

	// Gets suite definition version with GetSuiteDefinition API call.
	SuiteDefinitionVersion *string

	// Tags attached to the suite definition.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSuiteDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSuiteDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSuiteDefinition{}, middleware.After)
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
	if err = addOpGetSuiteDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSuiteDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSuiteDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdeviceadvisor",
		OperationName: "GetSuiteDefinition",
	}
}