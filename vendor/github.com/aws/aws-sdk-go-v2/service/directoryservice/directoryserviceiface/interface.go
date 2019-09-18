// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directoryserviceiface provides an interface to enable mocking the AWS Directory Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directoryserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

// ClientAPI provides an interface to enable mocking the
// directoryservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Directory Service.
//    func myFunc(svc directoryserviceiface.ClientAPI) bool {
//        // Make svc.AcceptSharedDirectory request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := directoryservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        directoryserviceiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptSharedDirectory(input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AcceptSharedDirectoryRequest(*directoryservice.AcceptSharedDirectoryInput) directoryservice.AcceptSharedDirectoryRequest

	AddIpRoutesRequest(*directoryservice.AddIpRoutesInput) directoryservice.AddIpRoutesRequest

	AddTagsToResourceRequest(*directoryservice.AddTagsToResourceInput) directoryservice.AddTagsToResourceRequest

	CancelSchemaExtensionRequest(*directoryservice.CancelSchemaExtensionInput) directoryservice.CancelSchemaExtensionRequest

	ConnectDirectoryRequest(*directoryservice.ConnectDirectoryInput) directoryservice.ConnectDirectoryRequest

	CreateAliasRequest(*directoryservice.CreateAliasInput) directoryservice.CreateAliasRequest

	CreateComputerRequest(*directoryservice.CreateComputerInput) directoryservice.CreateComputerRequest

	CreateConditionalForwarderRequest(*directoryservice.CreateConditionalForwarderInput) directoryservice.CreateConditionalForwarderRequest

	CreateDirectoryRequest(*directoryservice.CreateDirectoryInput) directoryservice.CreateDirectoryRequest

	CreateLogSubscriptionRequest(*directoryservice.CreateLogSubscriptionInput) directoryservice.CreateLogSubscriptionRequest

	CreateMicrosoftADRequest(*directoryservice.CreateMicrosoftADInput) directoryservice.CreateMicrosoftADRequest

	CreateSnapshotRequest(*directoryservice.CreateSnapshotInput) directoryservice.CreateSnapshotRequest

	CreateTrustRequest(*directoryservice.CreateTrustInput) directoryservice.CreateTrustRequest

	DeleteConditionalForwarderRequest(*directoryservice.DeleteConditionalForwarderInput) directoryservice.DeleteConditionalForwarderRequest

	DeleteDirectoryRequest(*directoryservice.DeleteDirectoryInput) directoryservice.DeleteDirectoryRequest

	DeleteLogSubscriptionRequest(*directoryservice.DeleteLogSubscriptionInput) directoryservice.DeleteLogSubscriptionRequest

	DeleteSnapshotRequest(*directoryservice.DeleteSnapshotInput) directoryservice.DeleteSnapshotRequest

	DeleteTrustRequest(*directoryservice.DeleteTrustInput) directoryservice.DeleteTrustRequest

	DeregisterEventTopicRequest(*directoryservice.DeregisterEventTopicInput) directoryservice.DeregisterEventTopicRequest

	DescribeConditionalForwardersRequest(*directoryservice.DescribeConditionalForwardersInput) directoryservice.DescribeConditionalForwardersRequest

	DescribeDirectoriesRequest(*directoryservice.DescribeDirectoriesInput) directoryservice.DescribeDirectoriesRequest

	DescribeDomainControllersRequest(*directoryservice.DescribeDomainControllersInput) directoryservice.DescribeDomainControllersRequest

	DescribeEventTopicsRequest(*directoryservice.DescribeEventTopicsInput) directoryservice.DescribeEventTopicsRequest

	DescribeSharedDirectoriesRequest(*directoryservice.DescribeSharedDirectoriesInput) directoryservice.DescribeSharedDirectoriesRequest

	DescribeSnapshotsRequest(*directoryservice.DescribeSnapshotsInput) directoryservice.DescribeSnapshotsRequest

	DescribeTrustsRequest(*directoryservice.DescribeTrustsInput) directoryservice.DescribeTrustsRequest

	DisableRadiusRequest(*directoryservice.DisableRadiusInput) directoryservice.DisableRadiusRequest

	DisableSsoRequest(*directoryservice.DisableSsoInput) directoryservice.DisableSsoRequest

	EnableRadiusRequest(*directoryservice.EnableRadiusInput) directoryservice.EnableRadiusRequest

	EnableSsoRequest(*directoryservice.EnableSsoInput) directoryservice.EnableSsoRequest

	GetDirectoryLimitsRequest(*directoryservice.GetDirectoryLimitsInput) directoryservice.GetDirectoryLimitsRequest

	GetSnapshotLimitsRequest(*directoryservice.GetSnapshotLimitsInput) directoryservice.GetSnapshotLimitsRequest

	ListIpRoutesRequest(*directoryservice.ListIpRoutesInput) directoryservice.ListIpRoutesRequest

	ListLogSubscriptionsRequest(*directoryservice.ListLogSubscriptionsInput) directoryservice.ListLogSubscriptionsRequest

	ListSchemaExtensionsRequest(*directoryservice.ListSchemaExtensionsInput) directoryservice.ListSchemaExtensionsRequest

	ListTagsForResourceRequest(*directoryservice.ListTagsForResourceInput) directoryservice.ListTagsForResourceRequest

	RegisterEventTopicRequest(*directoryservice.RegisterEventTopicInput) directoryservice.RegisterEventTopicRequest

	RejectSharedDirectoryRequest(*directoryservice.RejectSharedDirectoryInput) directoryservice.RejectSharedDirectoryRequest

	RemoveIpRoutesRequest(*directoryservice.RemoveIpRoutesInput) directoryservice.RemoveIpRoutesRequest

	RemoveTagsFromResourceRequest(*directoryservice.RemoveTagsFromResourceInput) directoryservice.RemoveTagsFromResourceRequest

	ResetUserPasswordRequest(*directoryservice.ResetUserPasswordInput) directoryservice.ResetUserPasswordRequest

	RestoreFromSnapshotRequest(*directoryservice.RestoreFromSnapshotInput) directoryservice.RestoreFromSnapshotRequest

	ShareDirectoryRequest(*directoryservice.ShareDirectoryInput) directoryservice.ShareDirectoryRequest

	StartSchemaExtensionRequest(*directoryservice.StartSchemaExtensionInput) directoryservice.StartSchemaExtensionRequest

	UnshareDirectoryRequest(*directoryservice.UnshareDirectoryInput) directoryservice.UnshareDirectoryRequest

	UpdateConditionalForwarderRequest(*directoryservice.UpdateConditionalForwarderInput) directoryservice.UpdateConditionalForwarderRequest

	UpdateNumberOfDomainControllersRequest(*directoryservice.UpdateNumberOfDomainControllersInput) directoryservice.UpdateNumberOfDomainControllersRequest

	UpdateRadiusRequest(*directoryservice.UpdateRadiusInput) directoryservice.UpdateRadiusRequest

	UpdateTrustRequest(*directoryservice.UpdateTrustInput) directoryservice.UpdateTrustRequest

	VerifyTrustRequest(*directoryservice.VerifyTrustInput) directoryservice.VerifyTrustRequest
}

var _ ClientAPI = (*directoryservice.Client)(nil)