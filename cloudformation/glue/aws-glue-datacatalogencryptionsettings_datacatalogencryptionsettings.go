package glue

import (
	"github.com/awslabs/goformation/v3/cloudformation/policies"
)

// DataCatalogEncryptionSettings_DataCatalogEncryptionSettings AWS CloudFormation Resource (AWS::Glue::DataCatalogEncryptionSettings.DataCatalogEncryptionSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html
type DataCatalogEncryptionSettings_DataCatalogEncryptionSettings struct {

	// ConnectionPasswordEncryption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-connectionpasswordencryption
	ConnectionPasswordEncryption *DataCatalogEncryptionSettings_ConnectionPasswordEncryption `json:"ConnectionPasswordEncryption,omitempty"`

	// EncryptionAtRest AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-encryptionatrest
	EncryptionAtRest *DataCatalogEncryptionSettings_EncryptionAtRest `json:"EncryptionAtRest,omitempty"`

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
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) ResourceCondition() string {
	return r._resourceCondition
}

// SetCondition specifies the logical ID of the condition that must be satisfied for this resource to be created
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetResourceCondition(condition string) {
	r._resourceCondition = condition
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) AWSCloudFormationType() string {
	return "AWS::Glue::DataCatalogEncryptionSettings.DataCatalogEncryptionSettings"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// CoreMetadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) CoreMetadata() map[string]interface{} {
	return r._metadata
}

// SetCoreMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetCoreMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *DataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
