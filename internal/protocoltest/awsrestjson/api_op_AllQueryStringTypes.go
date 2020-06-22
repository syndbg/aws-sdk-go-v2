// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This example uses all query string types.
func (c *Client) AllQueryStringTypes(ctx context.Context, params *AllQueryStringTypesInput, optFns ...func(*Options)) (*AllQueryStringTypesOutput, error) {
	stack := middleware.NewStack("AllQueryStringTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Serialize.Add(&awsRestjson1_serializeOpAllQueryStringTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAllQueryStringTypes{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "AllQueryStringTypes",
			Err:           err,
		}
	}
	out := result.(*AllQueryStringTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllQueryStringTypesInput struct {
	QueryBoolean       *bool
	QueryBooleanList   []*bool
	QueryByte          *int8
	QueryDouble        *float64
	QueryDoubleList    []*float64
	QueryEnum          types.FooEnum
	QueryEnumList      []types.FooEnum
	QueryFloat         *float32
	QueryInteger       *int32
	QueryIntegerList   []*int32
	QueryIntegerSet    []*int32
	QueryLong          *int64
	QueryShort         *int16
	QueryString        *string
	QueryStringList    []*string
	QueryStringSet     []*string
	QueryTimestamp     *time.Time
	QueryTimestampList []*time.Time
}

type AllQueryStringTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
