package medialive

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// Input_MediaConnectFlowRequest AWS CloudFormation Resource (AWS::MediaLive::Input.MediaConnectFlowRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflowrequest.html
type Input_MediaConnectFlowRequest struct {

	// FlowArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-mediaconnectflowrequest.html#cfn-medialive-input-mediaconnectflowrequest-flowarn
	FlowArn string `json:"FlowArn,omitempty"`

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
func (r *Input_MediaConnectFlowRequest) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *Input_MediaConnectFlowRequest) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Input_MediaConnectFlowRequest) AWSCloudFormationType() string {
	return "AWS::MediaLive::Input.MediaConnectFlowRequest"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Input_MediaConnectFlowRequest) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Input_MediaConnectFlowRequest) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Input_MediaConnectFlowRequest) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Input_MediaConnectFlowRequest) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Input_MediaConnectFlowRequest) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Input_MediaConnectFlowRequest) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
