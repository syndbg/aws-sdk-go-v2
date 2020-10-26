// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a set of tag keys to include in scheduled event notifications for your
// resources. To remove tags, use .
func (c *Client) RegisterInstanceEventNotificationAttributes(ctx context.Context, params *RegisterInstanceEventNotificationAttributesInput, optFns ...func(*Options)) (*RegisterInstanceEventNotificationAttributesOutput, error) {
	if params == nil {
		params = &RegisterInstanceEventNotificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterInstanceEventNotificationAttributes", params, optFns, addOperationRegisterInstanceEventNotificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterInstanceEventNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterInstanceEventNotificationAttributesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Information about the tag keys to register.
	InstanceTagAttribute *types.RegisterInstanceTagAttributeRequest
}

type RegisterInstanceEventNotificationAttributesOutput struct {

	// The resulting set of tag keys.
	InstanceTagAttribute *types.InstanceTagNotificationAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterInstanceEventNotificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRegisterInstanceEventNotificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterInstanceEventNotificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterInstanceEventNotificationAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterInstanceEventNotificationAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterInstanceEventNotificationAttributes",
	}
}