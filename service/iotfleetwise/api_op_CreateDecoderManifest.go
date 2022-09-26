// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the decoder manifest associated with a model manifest. To create a
// decoder manifest, the following must be true:
//
// * Every signal decoder has a
// unique name.
//
// * Each signal decoder is associated with a network interface.
//
// *
// Each network interface has a unique ID.
//
// * The signal decoders are specified in
// the model manifest.
func (c *Client) CreateDecoderManifest(ctx context.Context, params *CreateDecoderManifestInput, optFns ...func(*Options)) (*CreateDecoderManifestOutput, error) {
	if params == nil {
		params = &CreateDecoderManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDecoderManifest", params, optFns, c.addOperationCreateDecoderManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDecoderManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDecoderManifestInput struct {

	// The Amazon Resource Name (ARN) of the vehicle model (model manifest).
	//
	// This member is required.
	ModelManifestArn *string

	// The unique name of the decoder manifest to create.
	//
	// This member is required.
	Name *string

	// A brief description of the decoder manifest.
	Description *string

	// A list of information about available network interfaces.
	NetworkInterfaces []types.NetworkInterface

	// A list of information about signal decoders.
	SignalDecoders []types.SignalDecoder

	// Metadata that can be used to manage the decoder manifest.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDecoderManifestOutput struct {

	// The ARN of the created decoder manifest.
	//
	// This member is required.
	Arn *string

	// The name of the created decoder manifest.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDecoderManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateDecoderManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateDecoderManifest{}, middleware.After)
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
	if err = addOpCreateDecoderManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDecoderManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDecoderManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "CreateDecoderManifest",
	}
}
