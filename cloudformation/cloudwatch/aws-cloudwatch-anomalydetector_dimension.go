package cloudwatch

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// AnomalyDetector_Dimension AWS CloudFormation Resource (AWS::CloudWatch::AnomalyDetector.Dimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-dimension.html
type AnomalyDetector_Dimension struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-dimension.html#cfn-cloudwatch-anomalydetector-dimension-name
	Name string `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-dimension.html#cfn-cloudwatch-anomalydetector-dimension-value
	Value string `json:"Value,omitempty"`

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
func (r *AnomalyDetector_Dimension) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *AnomalyDetector_Dimension) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AnomalyDetector_Dimension) AWSCloudFormationType() string {
	return "AWS::CloudWatch::AnomalyDetector.Dimension"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AnomalyDetector_Dimension) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AnomalyDetector_Dimension) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AnomalyDetector_Dimension) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AnomalyDetector_Dimension) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AnomalyDetector_Dimension) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AnomalyDetector_Dimension) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
