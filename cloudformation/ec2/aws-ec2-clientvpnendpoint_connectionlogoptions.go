package ec2

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// ClientVpnEndpoint_ConnectionLogOptions AWS CloudFormation Resource (AWS::EC2::ClientVpnEndpoint.ConnectionLogOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html
type ClientVpnEndpoint_ConnectionLogOptions struct {

	// CloudwatchLogGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-cloudwatchloggroup
	CloudwatchLogGroup string `json:"CloudwatchLogGroup,omitempty"`

	// CloudwatchLogStream AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-cloudwatchlogstream
	CloudwatchLogStream string `json:"CloudwatchLogStream,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-connectionlogoptions.html#cfn-ec2-clientvpnendpoint-connectionlogoptions-enabled
	Enabled bool `json:"Enabled"`

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
func (r *ClientVpnEndpoint_ConnectionLogOptions) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *ClientVpnEndpoint_ConnectionLogOptions) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ClientVpnEndpoint_ConnectionLogOptions) AWSCloudFormationType() string {
	return "AWS::EC2::ClientVpnEndpoint.ConnectionLogOptions"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ClientVpnEndpoint_ConnectionLogOptions) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
