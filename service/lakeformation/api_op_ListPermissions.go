// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the principal permissions on the resource, filtered by the
// permissions of the caller. For example, if you are granted an ALTER permission,
// you are able to see only the principal permissions for ALTER. This operation
// returns only those permissions that have been explicitly granted. For
// information about permissions, see Security and Access Control to Metadata and
// Data
// (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
func (c *Client) ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissions", params, optFns, addOperationListPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionsInput struct {

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// Specifies a principal to filter the permissions returned.
	Principal *types.DataLakePrincipal

	// A resource where you will get a list of the principal permissions. This
	// operation does not support getting privileges on a table with columns. Instead,
	// call this operation on the table, and the operation returns the table and the
	// table w columns.
	Resource *types.Resource

	// Specifies a resource type to filter the permissions returned.
	ResourceType types.DataLakeResourceType
}

type ListPermissionsOutput struct {

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// A list of principals and their permissions on the resource for the specified
	// principal and resource types.
	PrincipalResourcePermissions []types.PrincipalResourcePermissions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissions{}, middleware.After)
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
	if err = addOpListPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissions(options.Region), middleware.Before); err != nil {
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

// ListPermissionsAPIClient is a client that implements the ListPermissions
// operation.
type ListPermissionsAPIClient interface {
	ListPermissions(context.Context, *ListPermissionsInput, ...func(*Options)) (*ListPermissionsOutput, error)
}

var _ ListPermissionsAPIClient = (*Client)(nil)

// ListPermissionsPaginatorOptions is the paginator options for ListPermissions
type ListPermissionsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionsPaginator is a paginator for ListPermissions
type ListPermissionsPaginator struct {
	options   ListPermissionsPaginatorOptions
	client    ListPermissionsAPIClient
	params    *ListPermissionsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionsPaginator returns a new ListPermissionsPaginator
func NewListPermissionsPaginator(client ListPermissionsAPIClient, params *ListPermissionsInput, optFns ...func(*ListPermissionsPaginatorOptions)) *ListPermissionsPaginator {
	options := ListPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListPermissionsInput{}
	}

	return &ListPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPermissions page.
func (p *ListPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPermissions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "ListPermissions",
	}
}