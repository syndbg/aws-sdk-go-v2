// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example uses a constant query string parameters and a label. This simply
// tests that labels and query string parameters are compatible. The fixed query
// string parameter named "hello" should in no way conflict with the label,
// {hello}.
func (c *Client) ConstantQueryString(ctx context.Context, params *ConstantQueryStringInput, optFns ...func(*Options)) (*ConstantQueryStringOutput, error) {
	if params == nil {
		params = &ConstantQueryStringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConstantQueryString", params, optFns, addOperationConstantQueryStringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConstantQueryStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConstantQueryStringInput struct {
	Hello *string
}

type ConstantQueryStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConstantQueryStringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConstantQueryString{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConstantQueryString{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpConstantQueryStringValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConstantQueryString(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opConstantQueryString(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConstantQueryString",
	}
}