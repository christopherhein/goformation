package applicationautoscaling

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// ScalableTarget_ScheduledAction AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalableTarget.ScheduledAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html
type ScalableTarget_ScheduledAction struct {

	// EndTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html#cfn-applicationautoscaling-scalabletarget-scheduledaction-endtime
	EndTime string `json:"EndTime,omitempty"`

	// ScalableTargetAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html#cfn-applicationautoscaling-scalabletarget-scheduledaction-scalabletargetaction
	ScalableTargetAction *ScalableTarget_ScalableTargetAction `json:"ScalableTargetAction,omitempty"`

	// Schedule AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html#cfn-applicationautoscaling-scalabletarget-scheduledaction-schedule
	Schedule string `json:"Schedule,omitempty"`

	// ScheduledActionName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html#cfn-applicationautoscaling-scalabletarget-scheduledaction-scheduledactionname
	ScheduledActionName string `json:"ScheduledActionName,omitempty"`

	// StartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html#cfn-applicationautoscaling-scalabletarget-scheduledaction-starttime
	StartTime string `json:"StartTime,omitempty"`

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
func (r *ScalableTarget_ScheduledAction) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *ScalableTarget_ScheduledAction) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ScalableTarget_ScheduledAction) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalableTarget.ScheduledAction"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ScalableTarget_ScheduledAction) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *ScalableTarget_ScheduledAction) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ScalableTarget_ScheduledAction) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *ScalableTarget_ScheduledAction) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ScalableTarget_ScheduledAction) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *ScalableTarget_ScheduledAction) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
