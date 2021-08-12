// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a hosted zone. If the hosted zone was created by another service, such
// as Cloud Map, see Deleting Public Hosted Zones That Were Created by Another
// Service
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DeleteHostedZone.html#delete-public-hosted-zone-created-by-another-service)
// in the Amazon Route 53 Developer Guide for information about how to delete it.
// (The process is the same for public and private hosted zones that were created
// by another service.) If you want to keep your domain registration but you want
// to stop routing internet traffic to your website or web application, we
// recommend that you delete resource record sets in the hosted zone instead of
// deleting the hosted zone. If you delete a hosted zone, you can't undelete it.
// You must create a new hosted zone and update the name servers for your domain
// registration, which can require up to 48 hours to take effect. (If you delegated
// responsibility for a subdomain to a hosted zone and you delete the child hosted
// zone, you must update the name servers in the parent hosted zone.) In addition,
// if you delete a hosted zone, someone could hijack the domain and route traffic
// to their own resources using your domain name. If you want to avoid the monthly
// charge for the hosted zone, you can transfer DNS service for the domain to a
// free DNS service. When you transfer DNS service, you have to update the name
// servers for the domain registration. If the domain is registered with Route 53,
// see UpdateDomainNameservers
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainNameservers.html)
// for information about how to replace Route 53 name servers with name servers for
// the new DNS service. If the domain is registered with another registrar, use the
// method provided by the registrar to update name servers for the domain
// registration. For more information, perform an internet search on "free DNS
// service." You can delete a hosted zone only if it contains only the default SOA
// record and NS resource record sets. If the hosted zone contains other resource
// record sets, you must delete them before you can delete the hosted zone. If you
// try to delete a hosted zone that contains other resource record sets, the
// request fails, and Route 53 returns a HostedZoneNotEmpty error. For information
// about deleting records from your hosted zone, see ChangeResourceRecordSets
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html).
// To verify that the hosted zone has been deleted, do one of the following:
//
// * Use
// the GetHostedZone action to request information about the hosted zone.
//
// * Use
// the ListHostedZones action to get a list of the hosted zones associated with the
// current Amazon Web Services account.
func (c *Client) DeleteHostedZone(ctx context.Context, params *DeleteHostedZoneInput, optFns ...func(*Options)) (*DeleteHostedZoneOutput, error) {
	if params == nil {
		params = &DeleteHostedZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHostedZone", params, optFns, c.addOperationDeleteHostedZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a hosted zone.
type DeleteHostedZoneInput struct {

	// The ID of the hosted zone you want to delete.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// A complex type that contains the response to a DeleteHostedZone request.
type DeleteHostedZoneOutput struct {

	// A complex type that contains the ID, the status, and the date and time of a
	// request to delete a hosted zone.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteHostedZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteHostedZone{}, middleware.After)
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
	if err = addOpDeleteHostedZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHostedZone(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteHostedZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteHostedZone",
	}
}
