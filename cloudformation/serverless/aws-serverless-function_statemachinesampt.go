package serverless

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// Function_StateMachineSAMPT AWS CloudFormation Resource (AWS::Serverless::Function.StateMachineSAMPT)
// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
type Function_StateMachineSAMPT struct {

	// StateMachineName AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/docs/policy_templates.rst
	StateMachineName string `json:"StateMachineName,omitempty"`

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
func (r *Function_StateMachineSAMPT) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *Function_StateMachineSAMPT) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_StateMachineSAMPT) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.StateMachineSAMPT"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Function_StateMachineSAMPT) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Function_StateMachineSAMPT) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Function_StateMachineSAMPT) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Function_StateMachineSAMPT) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Function_StateMachineSAMPT) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Function_StateMachineSAMPT) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
