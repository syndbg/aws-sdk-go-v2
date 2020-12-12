// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the fields that are included in log events in the specified
// log group, along with the percentage of log events that contain each field. The
// search is limited to a time period that you specify. In the results, fields that
// start with @ are fields generated by CloudWatch Logs. For example, @timestamp is
// the timestamp of each log event. For more information about the fields that are
// generated by CloudWatch logs, see Supported Logs and Discovered Fields
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html).
// The response results are sorted by the frequency percentage, starting with the
// highest percentage.
func (c *Client) GetLogGroupFields(ctx context.Context, params *GetLogGroupFieldsInput, optFns ...func(*Options)) (*GetLogGroupFieldsOutput, error) {
	if params == nil {
		params = &GetLogGroupFieldsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLogGroupFields", params, optFns, addOperationGetLogGroupFieldsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLogGroupFieldsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLogGroupFieldsInput struct {

	// The name of the log group to search.
	//
	// This member is required.
	LogGroupName *string

	// The time to set as the center of the query. If you specify time, the 8 minutes
	// before and 8 minutes after this time are searched. If you omit time, the past 15
	// minutes are queried. The time value is specified as epoch time, the number of
	// seconds since January 1, 1970, 00:00:00 UTC.
	Time *int64
}

type GetLogGroupFieldsOutput struct {

	// The array of fields found in the query. Each object in the array contains the
	// name of the field, along with the percentage of time it appeared in the log
	// events that were queried.
	LogGroupFields []types.LogGroupField

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLogGroupFieldsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLogGroupFields{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLogGroupFields{}, middleware.After)
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
	if err = addOpGetLogGroupFieldsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLogGroupFields(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLogGroupFields(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "GetLogGroupFields",
	}
}