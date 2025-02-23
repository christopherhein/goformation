package applicationautoscaling

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// ScalingPolicy_TargetTrackingScalingPolicyConfiguration AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html
type ScalingPolicy_TargetTrackingScalingPolicyConfiguration struct {

	// CustomizedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-customizedmetricspecification
	CustomizedMetricSpecification *ScalingPolicy_CustomizedMetricSpecification `json:"CustomizedMetricSpecification,omitempty"`

	// DisableScaleIn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-disablescalein
	DisableScaleIn bool `json:"DisableScaleIn,omitempty"`

	// PredefinedMetricSpecification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-predefinedmetricspecification
	PredefinedMetricSpecification *ScalingPolicy_PredefinedMetricSpecification `json:"PredefinedMetricSpecification,omitempty"`

	// ScaleInCooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleincooldown
	ScaleInCooldown int `json:"ScaleInCooldown,omitempty"`

	// ScaleOutCooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleoutcooldown
	ScaleOutCooldown int `json:"ScaleOutCooldown,omitempty"`

	// TargetValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-targetvalue
	TargetValue float64 `json:"TargetValue"`

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
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy.TargetTrackingScalingPolicyConfiguration"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ScalingPolicy_TargetTrackingScalingPolicyConfiguration) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
