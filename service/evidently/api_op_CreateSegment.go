// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to define a segment of your audience. A segment is a portion
// of your audience that share one or more characteristics. Examples could be
// Chrome browser users, users in Europe, or Firefox browser users in Europe who
// also fit other criteria that your application collects, such as age. Using a
// segment in an experiment limits that experiment to evaluate only the users who
// match the segment criteria. Using one or more segments in a launch allow you to
// define different traffic splits for the different audience segments. For more
// information about segment pattern syntax, see  Segment rule pattern syntax
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments-syntax.html).
// The pattern that you define for a segment is matched against the value of
// evaluationContext, which is passed into Evidently in the EvaluateFeature
// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_EvaluateFeature.html)
// operation, when Evidently assigns a feature variation to a user.
func (c *Client) CreateSegment(ctx context.Context, params *CreateSegmentInput, optFns ...func(*Options)) (*CreateSegmentOutput, error) {
	if params == nil {
		params = &CreateSegmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSegment", params, optFns, c.addOperationCreateSegmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSegmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSegmentInput struct {

	// A name for the segment.
	//
	// This member is required.
	Name *string

	// The pattern to use for the segment. For more information about pattern syntax,
	// see  Segment rule pattern syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments-syntax.html).
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Pattern *string

	// An optional description for this segment.
	Description *string

	// Assigns one or more tags (key-value pairs) to the segment. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. Tags don't have any semantic meaning to Amazon Web
	// Services and are interpreted strictly as strings of characters. You can
	// associate as many as 50 tags with a segment. For more information, see Tagging
	// Amazon Web Services resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSegmentOutput struct {

	// A structure that contains the complete information about the segment that was
	// just created.
	//
	// This member is required.
	Segment *types.Segment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSegmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSegment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSegment{}, middleware.After)
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
	if err = addOpCreateSegmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSegment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSegment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "CreateSegment",
	}
}
