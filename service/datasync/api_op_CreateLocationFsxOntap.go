// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon FSx for NetApp ONTAP file system that DataSync
// can access for a transfer. For more information, see Creating a location for FSx
// for ONTAP
// (https://docs.aws.amazon.com/datasync/latest/userguide/create-ontap-location.html).
func (c *Client) CreateLocationFsxOntap(ctx context.Context, params *CreateLocationFsxOntapInput, optFns ...func(*Options)) (*CreateLocationFsxOntapOutput, error) {
	if params == nil {
		params = &CreateLocationFsxOntapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationFsxOntap", params, optFns, c.addOperationCreateLocationFsxOntapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationFsxOntapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationFsxOntapInput struct {

	// Specifies the data transfer protocol that DataSync uses to access your Amazon
	// FSx file system.
	//
	// This member is required.
	Protocol *types.FsxProtocol

	// Specifies the Amazon EC2 security groups that provide access to your file
	// system's preferred subnet. The security groups must allow outbound traffic on
	// the following ports (depending on the protocol you use):
	//
	// * Network File System
	// (NFS): TCP ports 111, 635, and 2049
	//
	// * Server Message Block (SMB): TCP port
	// 445
	//
	// Your file system's security groups must also allow inbound traffic on the
	// same ports.
	//
	// This member is required.
	SecurityGroupArns []string

	// Specifies the ARN of the storage virtual machine (SVM) on your file system where
	// you're copying data to or from.
	//
	// This member is required.
	StorageVirtualMachineArn *string

	// Specifies the junction path (also known as a mount point) in the SVM volume
	// where you're copying data to or from (for example, /vol1). Don't specify a
	// junction path in the SVM's root volume. For more information, see Managing FSx
	// for ONTAP storage virtual machines
	// (https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) in the
	// Amazon FSx for NetApp ONTAP User Guide.
	Subdirectory *string

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationFsxOntapOutput struct {

	// Specifies the ARN of the FSx for ONTAP file system location that you create.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationFsxOntapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationFsxOntap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationFsxOntap{}, middleware.After)
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
	if err = addOpCreateLocationFsxOntapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationFsxOntap(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationFsxOntap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationFsxOntap",
	}
}
