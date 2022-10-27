//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Application.
func (a Application) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", a.Etag)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationGetHTTPSEndpoint.
func (a ApplicationGetHTTPSEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessModes", a.AccessModes)
	populate(objectMap, "destinationPort", a.DestinationPort)
	populate(objectMap, "disableGatewayAuth", a.DisableGatewayAuth)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "privateIPAddress", a.PrivateIPAddress)
	populate(objectMap, "publicPort", a.PublicPort)
	populate(objectMap, "subDomainSuffix", a.SubDomainSuffix)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationProperties.
func (a ApplicationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationState", a.ApplicationState)
	populate(objectMap, "applicationType", a.ApplicationType)
	populate(objectMap, "computeProfile", a.ComputeProfile)
	populate(objectMap, "createdDate", a.CreatedDate)
	populate(objectMap, "errors", a.Errors)
	populate(objectMap, "httpsEndpoints", a.HTTPSEndpoints)
	populate(objectMap, "installScriptActions", a.InstallScriptActions)
	populate(objectMap, "marketplaceIdentifier", a.MarketplaceIdentifier)
	populate(objectMap, "privateLinkConfigurations", a.PrivateLinkConfigurations)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "sshEndpoints", a.SSHEndpoints)
	populate(objectMap, "uninstallScriptActions", a.UninstallScriptActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AutoscaleRecurrence.
func (a AutoscaleRecurrence) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "schedule", a.Schedule)
	populate(objectMap, "timeZone", a.TimeZone)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AutoscaleSchedule.
func (a AutoscaleSchedule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "days", a.Days)
	populate(objectMap, "timeAndCapacity", a.TimeAndCapacity)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AzureMonitorSelectedConfigurations.
func (a AzureMonitorSelectedConfigurations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "configurationVersion", a.ConfigurationVersion)
	populate(objectMap, "globalConfigurations", a.GlobalConfigurations)
	populate(objectMap, "tableList", a.TableList)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	populate(objectMap, "zones", c.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterCreateParametersExtended.
func (c ClusterCreateParametersExtended) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "zones", c.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterCreateProperties.
func (c ClusterCreateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clusterDefinition", c.ClusterDefinition)
	populate(objectMap, "clusterVersion", c.ClusterVersion)
	populate(objectMap, "computeIsolationProperties", c.ComputeIsolationProperties)
	populate(objectMap, "computeProfile", c.ComputeProfile)
	populate(objectMap, "diskEncryptionProperties", c.DiskEncryptionProperties)
	populate(objectMap, "encryptionInTransitProperties", c.EncryptionInTransitProperties)
	populate(objectMap, "kafkaRestProperties", c.KafkaRestProperties)
	populate(objectMap, "minSupportedTlsVersion", c.MinSupportedTLSVersion)
	populate(objectMap, "networkProperties", c.NetworkProperties)
	populate(objectMap, "osType", c.OSType)
	populate(objectMap, "privateLinkConfigurations", c.PrivateLinkConfigurations)
	populate(objectMap, "securityProfile", c.SecurityProfile)
	populate(objectMap, "storageProfile", c.StorageProfile)
	populate(objectMap, "tier", c.Tier)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterCreateRequestValidationParameters.
func (c ClusterCreateRequestValidationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fetchAaddsResource", c.FetchAaddsResource)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "tenantId", c.TenantID)
	populate(objectMap, "type", c.Type)
	populate(objectMap, "zones", c.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterDefinition.
func (c ClusterDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "blueprint", c.Blueprint)
	populate(objectMap, "componentVersion", c.ComponentVersion)
	populate(objectMap, "configurations", &c.Configurations)
	populate(objectMap, "kind", c.Kind)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterGetProperties.
func (c ClusterGetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clusterDefinition", c.ClusterDefinition)
	populate(objectMap, "clusterHdpVersion", c.ClusterHdpVersion)
	populate(objectMap, "clusterId", c.ClusterID)
	populate(objectMap, "clusterState", c.ClusterState)
	populate(objectMap, "clusterVersion", c.ClusterVersion)
	populate(objectMap, "computeIsolationProperties", c.ComputeIsolationProperties)
	populate(objectMap, "computeProfile", c.ComputeProfile)
	populate(objectMap, "connectivityEndpoints", c.ConnectivityEndpoints)
	populate(objectMap, "createdDate", c.CreatedDate)
	populate(objectMap, "diskEncryptionProperties", c.DiskEncryptionProperties)
	populate(objectMap, "encryptionInTransitProperties", c.EncryptionInTransitProperties)
	populate(objectMap, "errors", c.Errors)
	populate(objectMap, "excludedServicesConfig", c.ExcludedServicesConfig)
	populate(objectMap, "kafkaRestProperties", c.KafkaRestProperties)
	populate(objectMap, "minSupportedTlsVersion", c.MinSupportedTLSVersion)
	populate(objectMap, "networkProperties", c.NetworkProperties)
	populate(objectMap, "osType", c.OSType)
	populate(objectMap, "privateEndpointConnections", c.PrivateEndpointConnections)
	populate(objectMap, "privateLinkConfigurations", c.PrivateLinkConfigurations)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "quotaInfo", c.QuotaInfo)
	populate(objectMap, "securityProfile", c.SecurityProfile)
	populate(objectMap, "storageProfile", c.StorageProfile)
	populate(objectMap, "tier", c.Tier)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterIdentity.
func (c ClusterIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", c.PrincipalID)
	populate(objectMap, "tenantId", c.TenantID)
	populate(objectMap, "type", c.Type)
	populate(objectMap, "userAssignedIdentities", c.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterPatchParameters.
func (c ClusterPatchParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ComputeProfile.
func (c ComputeProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "roles", c.Roles)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExecuteScriptActionParameters.
func (e ExecuteScriptActionParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "persistOnSuccess", e.PersistOnSuccess)
	populate(objectMap, "scriptActions", e.ScriptActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KafkaRestProperties.
func (k KafkaRestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clientGroupInfo", k.ClientGroupInfo)
	populate(objectMap, "configurationOverride", k.ConfigurationOverride)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkConfigurationProperties.
func (p PrivateLinkConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "ipConfigurations", p.IPConfigurations)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Role.
func (r Role) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoscale", r.AutoscaleConfiguration)
	populate(objectMap, "dataDisksGroups", r.DataDisksGroups)
	populate(objectMap, "encryptDataDisks", r.EncryptDataDisks)
	populate(objectMap, "hardwareProfile", r.HardwareProfile)
	populate(objectMap, "minInstanceCount", r.MinInstanceCount)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "osProfile", r.OSProfile)
	populate(objectMap, "scriptActions", r.ScriptActions)
	populate(objectMap, "targetInstanceCount", r.TargetInstanceCount)
	populate(objectMap, "VMGroupName", r.VMGroupName)
	populate(objectMap, "virtualNetworkProfile", r.VirtualNetworkProfile)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RuntimeScriptAction.
func (r RuntimeScriptAction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationName", r.ApplicationName)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "parameters", r.Parameters)
	populate(objectMap, "roles", r.Roles)
	populate(objectMap, "uri", r.URI)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RuntimeScriptActionDetail.
func (r RuntimeScriptActionDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationName", r.ApplicationName)
	populate(objectMap, "debugInformation", r.DebugInformation)
	populate(objectMap, "endTime", r.EndTime)
	populate(objectMap, "executionSummary", r.ExecutionSummary)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "operation", r.Operation)
	populate(objectMap, "parameters", r.Parameters)
	populate(objectMap, "roles", r.Roles)
	populate(objectMap, "scriptExecutionId", r.ScriptExecutionID)
	populate(objectMap, "startTime", r.StartTime)
	populate(objectMap, "status", r.Status)
	populate(objectMap, "uri", r.URI)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SSHProfile.
func (s SSHProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "publicKeys", s.PublicKeys)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SecurityProfile.
func (s SecurityProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aaddsResourceId", s.AaddsResourceID)
	populate(objectMap, "clusterUsersGroupDNs", s.ClusterUsersGroupDNs)
	populate(objectMap, "directoryType", s.DirectoryType)
	populate(objectMap, "domain", s.Domain)
	populate(objectMap, "domainUserPassword", s.DomainUserPassword)
	populate(objectMap, "domainUsername", s.DomainUsername)
	populate(objectMap, "ldapsUrls", s.LdapsUrls)
	populate(objectMap, "msiResourceId", s.MsiResourceID)
	populate(objectMap, "organizationalUnitDN", s.OrganizationalUnitDN)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type StorageProfile.
func (s StorageProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "storageaccounts", s.Storageaccounts)
	return json.Marshal(objectMap)
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