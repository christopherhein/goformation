package codepipeline

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// CustomActionType_ConfigurationProperties AWS CloudFormation Resource (AWS::CodePipeline::CustomActionType.ConfigurationProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html
type CustomActionType_ConfigurationProperties struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-description
	Description string `json:"Description,omitempty"`

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	Key bool `json:"Key"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-name
	Name string `json:"Name,omitempty"`

	// Queryable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	Queryable bool `json:"Queryable,omitempty"`

	// Required AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-required
	Required bool `json:"Required"`

	// Secret AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-secret
	Secret bool `json:"Secret"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-type
	Type string `json:"Type,omitempty"`

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
func (r *CustomActionType_ConfigurationProperties) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *CustomActionType_ConfigurationProperties) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *CustomActionType_ConfigurationProperties) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.ConfigurationProperties"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *CustomActionType_ConfigurationProperties) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *CustomActionType_ConfigurationProperties) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *CustomActionType_ConfigurationProperties) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *CustomActionType_ConfigurationProperties) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *CustomActionType_ConfigurationProperties) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *CustomActionType_ConfigurationProperties) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
