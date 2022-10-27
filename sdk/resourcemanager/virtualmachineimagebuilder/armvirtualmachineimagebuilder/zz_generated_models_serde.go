//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ImageTemplate.
func (i ImageTemplate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", i.ID)
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "tags", i.Tags)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateDistributor.
func (i ImageTemplateDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = i.Type
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateFileCustomizer.
func (i ImageTemplateFileCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "destination", i.Destination)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "sourceUri", i.SourceURI)
	objectMap["type"] = "File"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateFileCustomizer.
func (i *ImageTemplateFileCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "destination":
			err = unpopulate(val, "Destination", &i.Destination)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, "SHA256Checksum", &i.SHA256Checksum)
			delete(rawMsg, key)
		case "sourceUri":
			err = unpopulate(val, "SourceURI", &i.SourceURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateIdentity.
func (i ImageTemplateIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateLastRunStatus.
func (i ImageTemplateLastRunStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "endTime", i.EndTime)
	populate(objectMap, "message", i.Message)
	populate(objectMap, "runState", i.RunState)
	populate(objectMap, "runSubState", i.RunSubState)
	populateTimeRFC3339(objectMap, "startTime", i.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateLastRunStatus.
func (i *ImageTemplateLastRunStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &i.EndTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &i.Message)
			delete(rawMsg, key)
		case "runState":
			err = unpopulate(val, "RunState", &i.RunState)
			delete(rawMsg, key)
		case "runSubState":
			err = unpopulate(val, "RunSubState", &i.RunSubState)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &i.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateManagedImageDistributor.
func (i ImageTemplateManagedImageDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "imageId", i.ImageID)
	populate(objectMap, "location", i.Location)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = "ManagedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateManagedImageDistributor.
func (i *ImageTemplateManagedImageDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, "ArtifactTags", &i.ArtifactTags)
			delete(rawMsg, key)
		case "imageId":
			err = unpopulate(val, "ImageID", &i.ImageID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &i.Location)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, "RunOutputName", &i.RunOutputName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateManagedImageSource.
func (i ImageTemplateManagedImageSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "imageId", i.ImageID)
	objectMap["type"] = "ManagedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateManagedImageSource.
func (i *ImageTemplateManagedImageSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "imageId":
			err = unpopulate(val, "ImageID", &i.ImageID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePlatformImageSource.
func (i ImageTemplatePlatformImageSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "exactVersion", i.ExactVersion)
	populate(objectMap, "offer", i.Offer)
	populate(objectMap, "planInfo", i.PlanInfo)
	populate(objectMap, "publisher", i.Publisher)
	populate(objectMap, "sku", i.SKU)
	objectMap["type"] = "PlatformImage"
	populate(objectMap, "version", i.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePlatformImageSource.
func (i *ImageTemplatePlatformImageSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "exactVersion":
			err = unpopulate(val, "ExactVersion", &i.ExactVersion)
			delete(rawMsg, key)
		case "offer":
			err = unpopulate(val, "Offer", &i.Offer)
			delete(rawMsg, key)
		case "planInfo":
			err = unpopulate(val, "PlanInfo", &i.PlanInfo)
			delete(rawMsg, key)
		case "publisher":
			err = unpopulate(val, "Publisher", &i.Publisher)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &i.SKU)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &i.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePowerShellCustomizer.
func (i ImageTemplatePowerShellCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "runAsSystem", i.RunAsSystem)
	populate(objectMap, "runElevated", i.RunElevated)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "PowerShell"
	populate(objectMap, "validExitCodes", i.ValidExitCodes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePowerShellCustomizer.
func (i *ImageTemplatePowerShellCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, "Inline", &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "runAsSystem":
			err = unpopulate(val, "RunAsSystem", &i.RunAsSystem)
			delete(rawMsg, key)
		case "runElevated":
			err = unpopulate(val, "RunElevated", &i.RunElevated)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, "SHA256Checksum", &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, "ScriptURI", &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		case "validExitCodes":
			err = unpopulate(val, "ValidExitCodes", &i.ValidExitCodes)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePowerShellValidator.
func (i ImageTemplatePowerShellValidator) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "runAsSystem", i.RunAsSystem)
	populate(objectMap, "runElevated", i.RunElevated)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "PowerShell"
	populate(objectMap, "validExitCodes", i.ValidExitCodes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePowerShellValidator.
func (i *ImageTemplatePowerShellValidator) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, "Inline", &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "runAsSystem":
			err = unpopulate(val, "RunAsSystem", &i.RunAsSystem)
			delete(rawMsg, key)
		case "runElevated":
			err = unpopulate(val, "RunElevated", &i.RunElevated)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, "SHA256Checksum", &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, "ScriptURI", &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		case "validExitCodes":
			err = unpopulate(val, "ValidExitCodes", &i.ValidExitCodes)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateProperties.
func (i ImageTemplateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "buildTimeoutInMinutes", i.BuildTimeoutInMinutes)
	populate(objectMap, "customize", i.Customize)
	populate(objectMap, "distribute", i.Distribute)
	populate(objectMap, "exactStagingResourceGroup", i.ExactStagingResourceGroup)
	populate(objectMap, "lastRunStatus", i.LastRunStatus)
	populate(objectMap, "provisioningError", i.ProvisioningError)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "source", i.Source)
	populate(objectMap, "stagingResourceGroup", i.StagingResourceGroup)
	populate(objectMap, "vmProfile", i.VMProfile)
	populate(objectMap, "validate", i.Validate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateProperties.
func (i *ImageTemplateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "buildTimeoutInMinutes":
			err = unpopulate(val, "BuildTimeoutInMinutes", &i.BuildTimeoutInMinutes)
			delete(rawMsg, key)
		case "customize":
			i.Customize, err = unmarshalImageTemplateCustomizerClassificationArray(val)
			delete(rawMsg, key)
		case "distribute":
			i.Distribute, err = unmarshalImageTemplateDistributorClassificationArray(val)
			delete(rawMsg, key)
		case "exactStagingResourceGroup":
			err = unpopulate(val, "ExactStagingResourceGroup", &i.ExactStagingResourceGroup)
			delete(rawMsg, key)
		case "lastRunStatus":
			err = unpopulate(val, "LastRunStatus", &i.LastRunStatus)
			delete(rawMsg, key)
		case "provisioningError":
			err = unpopulate(val, "ProvisioningError", &i.ProvisioningError)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "source":
			i.Source, err = unmarshalImageTemplateSourceClassification(val)
			delete(rawMsg, key)
		case "stagingResourceGroup":
			err = unpopulate(val, "StagingResourceGroup", &i.StagingResourceGroup)
			delete(rawMsg, key)
		case "vmProfile":
			err = unpopulate(val, "VMProfile", &i.VMProfile)
			delete(rawMsg, key)
		case "validate":
			err = unpopulate(val, "Validate", &i.Validate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplatePropertiesValidate.
func (i ImageTemplatePropertiesValidate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "continueDistributeOnFailure", i.ContinueDistributeOnFailure)
	populate(objectMap, "inVMValidations", i.InVMValidations)
	populate(objectMap, "sourceValidationOnly", i.SourceValidationOnly)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplatePropertiesValidate.
func (i *ImageTemplatePropertiesValidate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "continueDistributeOnFailure":
			err = unpopulate(val, "ContinueDistributeOnFailure", &i.ContinueDistributeOnFailure)
			delete(rawMsg, key)
		case "inVMValidations":
			i.InVMValidations, err = unmarshalImageTemplateInVMValidatorClassificationArray(val)
			delete(rawMsg, key)
		case "sourceValidationOnly":
			err = unpopulate(val, "SourceValidationOnly", &i.SourceValidationOnly)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateRestartCustomizer.
func (i ImageTemplateRestartCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", i.Name)
	populate(objectMap, "restartCheckCommand", i.RestartCheckCommand)
	populate(objectMap, "restartCommand", i.RestartCommand)
	populate(objectMap, "restartTimeout", i.RestartTimeout)
	objectMap["type"] = "WindowsRestart"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateRestartCustomizer.
func (i *ImageTemplateRestartCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "restartCheckCommand":
			err = unpopulate(val, "RestartCheckCommand", &i.RestartCheckCommand)
			delete(rawMsg, key)
		case "restartCommand":
			err = unpopulate(val, "RestartCommand", &i.RestartCommand)
			delete(rawMsg, key)
		case "restartTimeout":
			err = unpopulate(val, "RestartTimeout", &i.RestartTimeout)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateSharedImageDistributor.
func (i ImageTemplateSharedImageDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "excludeFromLatest", i.ExcludeFromLatest)
	populate(objectMap, "galleryImageId", i.GalleryImageID)
	populate(objectMap, "replicationRegions", i.ReplicationRegions)
	populate(objectMap, "runOutputName", i.RunOutputName)
	populate(objectMap, "storageAccountType", i.StorageAccountType)
	objectMap["type"] = "SharedImage"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateSharedImageDistributor.
func (i *ImageTemplateSharedImageDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, "ArtifactTags", &i.ArtifactTags)
			delete(rawMsg, key)
		case "excludeFromLatest":
			err = unpopulate(val, "ExcludeFromLatest", &i.ExcludeFromLatest)
			delete(rawMsg, key)
		case "galleryImageId":
			err = unpopulate(val, "GalleryImageID", &i.GalleryImageID)
			delete(rawMsg, key)
		case "replicationRegions":
			err = unpopulate(val, "ReplicationRegions", &i.ReplicationRegions)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, "RunOutputName", &i.RunOutputName)
			delete(rawMsg, key)
		case "storageAccountType":
			err = unpopulate(val, "StorageAccountType", &i.StorageAccountType)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateSharedImageVersionSource.
func (i ImageTemplateSharedImageVersionSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "imageVersionId", i.ImageVersionID)
	objectMap["type"] = "SharedImageVersion"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateSharedImageVersionSource.
func (i *ImageTemplateSharedImageVersionSource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "imageVersionId":
			err = unpopulate(val, "ImageVersionID", &i.ImageVersionID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateShellCustomizer.
func (i ImageTemplateShellCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "Shell"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateShellCustomizer.
func (i *ImageTemplateShellCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, "Inline", &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, "SHA256Checksum", &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, "ScriptURI", &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateShellValidator.
func (i ImageTemplateShellValidator) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "inline", i.Inline)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "sha256Checksum", i.SHA256Checksum)
	populate(objectMap, "scriptUri", i.ScriptURI)
	objectMap["type"] = "Shell"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateShellValidator.
func (i *ImageTemplateShellValidator) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "inline":
			err = unpopulate(val, "Inline", &i.Inline)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "sha256Checksum":
			err = unpopulate(val, "SHA256Checksum", &i.SHA256Checksum)
			delete(rawMsg, key)
		case "scriptUri":
			err = unpopulate(val, "ScriptURI", &i.ScriptURI)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateUpdateParameters.
func (i ImageTemplateUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", i.Identity)
	populate(objectMap, "tags", i.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateVMProfile.
func (i ImageTemplateVMProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "osDiskSizeGB", i.OSDiskSizeGB)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	populate(objectMap, "vmSize", i.VMSize)
	populate(objectMap, "vnetConfig", i.VnetConfig)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateVhdDistributor.
func (i ImageTemplateVhdDistributor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactTags", i.ArtifactTags)
	populate(objectMap, "runOutputName", i.RunOutputName)
	objectMap["type"] = "VHD"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateVhdDistributor.
func (i *ImageTemplateVhdDistributor) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactTags":
			err = unpopulate(val, "ArtifactTags", &i.ArtifactTags)
			delete(rawMsg, key)
		case "runOutputName":
			err = unpopulate(val, "RunOutputName", &i.RunOutputName)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImageTemplateWindowsUpdateCustomizer.
func (i ImageTemplateWindowsUpdateCustomizer) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "filters", i.Filters)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "searchCriteria", i.SearchCriteria)
	objectMap["type"] = "WindowsUpdate"
	populate(objectMap, "updateLimit", i.UpdateLimit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImageTemplateWindowsUpdateCustomizer.
func (i *ImageTemplateWindowsUpdateCustomizer) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "filters":
			err = unpopulate(val, "Filters", &i.Filters)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "searchCriteria":
			err = unpopulate(val, "SearchCriteria", &i.SearchCriteria)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		case "updateLimit":
			err = unpopulate(val, "UpdateLimit", &i.UpdateLimit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}