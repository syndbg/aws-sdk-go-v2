// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Updates agent status.
func (c *Client) UpdateAgentStatus(ctx context.Context, params *UpdateAgentStatusInput, optFns ...func(*Options)) (*UpdateAgentStatusOutput, error) {
	if params == nil {
		params = &UpdateAgentStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAgentStatus", params, optFns, c.addOperationUpdateAgentStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAgentStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAgentStatusInput struct {

	// The identifier of the agent status.
	//
	// This member is required.
	AgentStatusId *string

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The description of the agent status.
	Description *string

	// The display order of the agent status.
	DisplayOrder *int32

	// The name of the agent status.
	Name *string

	// A number indicating the reset order of the agent status.
	ResetOrderNumber bool

	// The state of the agent status.
	State types.AgentStatusState

	noSmithyDocumentSerde
}

type UpdateAgentStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAgentStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAgentStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAgentStatus{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpUpdateAgentStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAgentStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAgentStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateAgentStatus",
	}
}
