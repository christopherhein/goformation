package opsworks

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// Instance_EbsBlockDevice AWS CloudFormation Resource (AWS::OpsWorks::Instance.EbsBlockDevice)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html
type Instance_EbsBlockDevice struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-deleteontermination
	DeleteOnTermination bool `json:"DeleteOnTermination,omitempty"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-iops
	Iops int `json:"Iops,omitempty"`

	// SnapshotId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-snapshotid
	SnapshotId string `json:"SnapshotId,omitempty"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumesize
	VolumeSize int `json:"VolumeSize,omitempty"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-ebsblockdevice.html#cfn-opsworks-instance-ebsblockdevice-volumetype
	VolumeType string `json:"VolumeType,omitempty"`

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
func (r *Instance_EbsBlockDevice) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *Instance_EbsBlockDevice) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Instance_EbsBlockDevice) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance.EbsBlockDevice"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Instance_EbsBlockDevice) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *Instance_EbsBlockDevice) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Instance_EbsBlockDevice) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *Instance_EbsBlockDevice) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Instance_EbsBlockDevice) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *Instance_EbsBlockDevice) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
