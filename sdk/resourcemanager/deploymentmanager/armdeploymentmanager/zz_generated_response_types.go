//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

// ArtifactSourcesClientCreateOrUpdateResponse contains the response from method ArtifactSourcesClient.CreateOrUpdate.
type ArtifactSourcesClientCreateOrUpdateResponse struct {
	ArtifactSource
}

// ArtifactSourcesClientDeleteResponse contains the response from method ArtifactSourcesClient.Delete.
type ArtifactSourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// ArtifactSourcesClientGetResponse contains the response from method ArtifactSourcesClient.Get.
type ArtifactSourcesClientGetResponse struct {
	ArtifactSource
}

// ArtifactSourcesClientListResponse contains the response from method ArtifactSourcesClient.List.
type ArtifactSourcesClientListResponse struct {
	// The list of artifact sources.
	ArtifactSourceArray []*ArtifactSource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsList
}

// RolloutsClientCancelResponse contains the response from method RolloutsClient.Cancel.
type RolloutsClientCancelResponse struct {
	Rollout
}

// RolloutsClientCreateOrUpdateResponse contains the response from method RolloutsClient.CreateOrUpdate.
type RolloutsClientCreateOrUpdateResponse struct {
	RolloutRequest
}

// RolloutsClientDeleteResponse contains the response from method RolloutsClient.Delete.
type RolloutsClientDeleteResponse struct {
	// placeholder for future response values
}

// RolloutsClientGetResponse contains the response from method RolloutsClient.Get.
type RolloutsClientGetResponse struct {
	Rollout
}

// RolloutsClientListResponse contains the response from method RolloutsClient.List.
type RolloutsClientListResponse struct {
	// The list of rollouts.
	RolloutArray []*Rollout
}

// RolloutsClientRestartResponse contains the response from method RolloutsClient.Restart.
type RolloutsClientRestartResponse struct {
	Rollout
}

// ServiceTopologiesClientCreateOrUpdateResponse contains the response from method ServiceTopologiesClient.CreateOrUpdate.
type ServiceTopologiesClientCreateOrUpdateResponse struct {
	ServiceTopologyResource
}

// ServiceTopologiesClientDeleteResponse contains the response from method ServiceTopologiesClient.Delete.
type ServiceTopologiesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceTopologiesClientGetResponse contains the response from method ServiceTopologiesClient.Get.
type ServiceTopologiesClientGetResponse struct {
	ServiceTopologyResource
}

// ServiceTopologiesClientListResponse contains the response from method ServiceTopologiesClient.List.
type ServiceTopologiesClientListResponse struct {
	// The list of service topologies.
	ServiceTopologyResourceArray []*ServiceTopologyResource
}

// ServiceUnitsClientCreateOrUpdateResponse contains the response from method ServiceUnitsClient.CreateOrUpdate.
type ServiceUnitsClientCreateOrUpdateResponse struct {
	ServiceUnitResource
}

// ServiceUnitsClientDeleteResponse contains the response from method ServiceUnitsClient.Delete.
type ServiceUnitsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceUnitsClientGetResponse contains the response from method ServiceUnitsClient.Get.
type ServiceUnitsClientGetResponse struct {
	ServiceUnitResource
}

// ServiceUnitsClientListResponse contains the response from method ServiceUnitsClient.List.
type ServiceUnitsClientListResponse struct {
	// The list of service units.
	ServiceUnitResourceArray []*ServiceUnitResource
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServiceResource
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	// The list of services.
	ServiceResourceArray []*ServiceResource
}

// StepsClientCreateOrUpdateResponse contains the response from method StepsClient.CreateOrUpdate.
type StepsClientCreateOrUpdateResponse struct {
	StepResource
}

// StepsClientDeleteResponse contains the response from method StepsClient.Delete.
type StepsClientDeleteResponse struct {
	// placeholder for future response values
}

// StepsClientGetResponse contains the response from method StepsClient.Get.
type StepsClientGetResponse struct {
	StepResource
}

// StepsClientListResponse contains the response from method StepsClient.List.
type StepsClientListResponse struct {
	// The list of steps.
	StepResourceArray []*StepResource
}