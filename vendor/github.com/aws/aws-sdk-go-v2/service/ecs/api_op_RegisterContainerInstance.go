// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RegisterContainerInstanceRequest
type RegisterContainerInstanceInput struct {
	_ struct{} `type:"structure"`

	// The container instance attributes that this container instance supports.
	Attributes []Attribute `locationName:"attributes" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the cluster with which
	// to register your container instance. If you do not specify a cluster, the
	// default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The ARN of the container instance (if it was previously registered).
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The instance identity document for the EC2 instance to register. This document
	// can be found by running the following command from the instance: curl http://169.254.169.254/latest/dynamic/instance-identity/document/
	InstanceIdentityDocument *string `locationName:"instanceIdentityDocument" type:"string"`

	// The instance identity document signature for the EC2 instance to register.
	// This signature can be found by running the following command from the instance:
	// curl http://169.254.169.254/latest/dynamic/instance-identity/signature/
	InstanceIdentityDocumentSignature *string `locationName:"instanceIdentityDocumentSignature" type:"string"`

	// The devices that are available on the container instance. The only supported
	// device type is a GPU.
	PlatformDevices []PlatformDevice `locationName:"platformDevices" type:"list"`

	// The metadata that you apply to the container instance to help you categorize
	// and organize them. Each tag consists of a key and an optional value, both
	// of which you define.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case-sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for either keys or values as it is reserved for AWS use. You
	//    cannot edit or delete tag keys or values with this prefix. Tags with this
	//    prefix do not count against your tags per resource limit.
	Tags []Tag `locationName:"tags" type:"list"`

	// The resources available on the instance.
	TotalResources []Resource `locationName:"totalResources" type:"list"`

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterContainerInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterContainerInstanceInput"}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PlatformDevices != nil {
		for i, v := range s.PlatformDevices {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PlatformDevices", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RegisterContainerInstanceResponse
type RegisterContainerInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The container instance that was registered.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterContainerInstance = "RegisterContainerInstance"

// RegisterContainerInstanceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Registers an EC2 instance into the specified cluster. This instance becomes
// available to place containers on.
//
//    // Example sending a request using RegisterContainerInstanceRequest.
//    req := client.RegisterContainerInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RegisterContainerInstance
func (c *Client) RegisterContainerInstanceRequest(input *RegisterContainerInstanceInput) RegisterContainerInstanceRequest {
	op := &aws.Operation{
		Name:       opRegisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterContainerInstanceInput{}
	}

	req := c.newRequest(op, input, &RegisterContainerInstanceOutput{})
	return RegisterContainerInstanceRequest{Request: req, Input: input, Copy: c.RegisterContainerInstanceRequest}
}

// RegisterContainerInstanceRequest is the request type for the
// RegisterContainerInstance API operation.
type RegisterContainerInstanceRequest struct {
	*aws.Request
	Input *RegisterContainerInstanceInput
	Copy  func(*RegisterContainerInstanceInput) RegisterContainerInstanceRequest
}

// Send marshals and sends the RegisterContainerInstance API request.
func (r RegisterContainerInstanceRequest) Send(ctx context.Context) (*RegisterContainerInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterContainerInstanceResponse{
		RegisterContainerInstanceOutput: r.Request.Data.(*RegisterContainerInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterContainerInstanceResponse is the response type for the
// RegisterContainerInstance API operation.
type RegisterContainerInstanceResponse struct {
	*RegisterContainerInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterContainerInstance request.
func (r *RegisterContainerInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
