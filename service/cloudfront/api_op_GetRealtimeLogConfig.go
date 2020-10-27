// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a real-time log configuration. To get a real-time log configuration, you
// can provide the configuration’s name or its Amazon Resource Name (ARN). You must
// provide at least one. If you provide both, CloudFront uses the name to identify
// the real-time log configuration to get.
func (c *Client) GetRealtimeLogConfig(ctx context.Context, params *GetRealtimeLogConfigInput, optFns ...func(*Options)) (*GetRealtimeLogConfigOutput, error) {
	if params == nil {
		params = &GetRealtimeLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRealtimeLogConfig", params, optFns, addOperationGetRealtimeLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRealtimeLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRealtimeLogConfigInput struct {

	// The Amazon Resource Name (ARN) of the real-time log configuration to get.
	ARN *string

	// The name of the real-time log configuration to get.
	Name *string
}

type GetRealtimeLogConfigOutput struct {

	// A real-time log configuration.
	RealtimeLogConfig *types.RealtimeLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRealtimeLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetRealtimeLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetRealtimeLogConfig{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRealtimeLogConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRealtimeLogConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetRealtimeLogConfig",
	}
}