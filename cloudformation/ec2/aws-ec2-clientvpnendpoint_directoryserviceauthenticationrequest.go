package ec2

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// ClientVpnEndpoint_DirectoryServiceAuthenticationRequest AWS CloudFormation Resource (AWS::EC2::ClientVpnEndpoint.DirectoryServiceAuthenticationRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-directoryserviceauthenticationrequest.html
type ClientVpnEndpoint_DirectoryServiceAuthenticationRequest struct {

	// DirectoryId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-directoryserviceauthenticationrequest.html#cfn-ec2-clientvpnendpoint-directoryserviceauthenticationrequest-directoryid
	DirectoryId string `json:"DirectoryId,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}

	// _resourceCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	_resourceCondition string
}

// Condition returns the logical ID of the condition that must be satisfied for this resource to be created
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) AWSCloudFormationType() string {
	return "AWS::EC2::ClientVpnEndpoint.DirectoryServiceAuthenticationRequest"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ClientVpnEndpoint_DirectoryServiceAuthenticationRequest) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
