package sagemaker

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook AWS CloudFormation Resource (AWS::SageMaker::NotebookInstanceLifecycleConfig.NotebookInstanceLifecycleHook)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.html
type NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook struct {

	// Content AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.html#cfn-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook-content
	Content string `json:"Content,omitempty"`

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
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) AWSCloudFormationType() string {
	return "AWS::SageMaker::NotebookInstanceLifecycleConfig.NotebookInstanceLifecycleHook"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
