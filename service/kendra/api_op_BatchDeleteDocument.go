// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes one or more documents from an index. The documents must have been added
// with the BatchPutDocument operation. The documents are deleted asynchronously.
// You can see the progress of the deletion by using AWS CloudWatch. Any error
// messages releated to the processing of the batch are sent to you CloudWatch log.
func (c *Client) BatchDeleteDocument(ctx context.Context, params *BatchDeleteDocumentInput, optFns ...func(*Options)) (*BatchDeleteDocumentOutput, error) {
	if params == nil {
		params = &BatchDeleteDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteDocument", params, optFns, addOperationBatchDeleteDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteDocumentInput struct {

	// One or more identifiers for documents to delete from the index.
	//
	// This member is required.
	DocumentIdList []string

	// The identifier of the index that contains the documents to delete.
	//
	// This member is required.
	IndexId *string

	// Maps a particular data source sync job to a particular data source.
	DataSourceSyncJobMetricTarget *types.DataSourceSyncJobMetricTarget
}

type BatchDeleteDocumentOutput struct {

	// A list of documents that could not be removed from the index. Each entry
	// contains an error message that indicates why the document couldn't be removed
	// from the index.
	FailedDocuments []types.BatchDeleteDocumentResponseFailedDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteDocument{}, middleware.After)
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
	if err = addOpBatchDeleteDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "BatchDeleteDocument",
	}
}