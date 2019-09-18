// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetTypeRequest
type GetTypeInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The type format: SDL or JSON.
	//
	// Format is a required field
	Format TypeDefinitionFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`

	// The type name.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTypeInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTypeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TypeName != nil {
		v := *s.TypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "typeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetTypeResponse
type GetTypeOutput struct {
	_ struct{} `type:"structure"`

	// The Type object.
	Type *Type `locationName:"type" type:"structure"`
}

// String returns the string representation
func (s GetTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Type != nil {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "type", v, metadata)
	}
	return nil
}

const opGetType = "GetType"

// GetTypeRequest returns a request value for making API operation for
// AWS AppSync.
//
// Retrieves a Type object.
//
//    // Example sending a request using GetTypeRequest.
//    req := client.GetTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetType
func (c *Client) GetTypeRequest(input *GetTypeInput) GetTypeRequest {
	op := &aws.Operation{
		Name:       opGetType,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}",
	}

	if input == nil {
		input = &GetTypeInput{}
	}

	req := c.newRequest(op, input, &GetTypeOutput{})
	return GetTypeRequest{Request: req, Input: input, Copy: c.GetTypeRequest}
}

// GetTypeRequest is the request type for the
// GetType API operation.
type GetTypeRequest struct {
	*aws.Request
	Input *GetTypeInput
	Copy  func(*GetTypeInput) GetTypeRequest
}

// Send marshals and sends the GetType API request.
func (r GetTypeRequest) Send(ctx context.Context) (*GetTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTypeResponse{
		GetTypeOutput: r.Request.Data.(*GetTypeOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTypeResponse is the response type for the
// GetType API operation.
type GetTypeResponse struct {
	*GetTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetType request.
func (r *GetTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}