// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation. Adds a streaming source to your Amazon Kinesis
// application. For conceptual information, see Configuring Application Input
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
// You can add a streaming source either when you create an application or you can
// use this operation to add a streaming source after you create an application.
// For more information, see CreateApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_CreateApplication.html).
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the
// DescribeApplication
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
// operation to find the current application version. This operation requires
// permissions to perform the kinesisanalytics:AddApplicationInput action.
func (c *Client) AddApplicationInput(ctx context.Context, params *AddApplicationInputInput, optFns ...func(*Options)) (*AddApplicationInputOutput, error) {
	if params == nil {
		params = &AddApplicationInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationInput", params, optFns, addOperationAddApplicationInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type AddApplicationInputInput struct {

	// Name of your existing Amazon Kinesis Analytics application to which you want to
	// add the streaming source.
	//
	// This member is required.
	ApplicationName *string

	// Current version of your Amazon Kinesis Analytics application. You can use the
	// DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to find the current application version.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The Input
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_Input.html) to add.
	//
	// This member is required.
	Input *types.Input
}

//
type AddApplicationInputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAddApplicationInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationInput{}, middleware.After)
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
	if err = addOpAddApplicationInputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationInput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationInput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "AddApplicationInput",
	}
}