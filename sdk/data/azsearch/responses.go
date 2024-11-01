// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsearch

// DataSourcesClientCreateOrUpdateResponse contains the response from method DataSourcesClient.CreateOrUpdate.
type DataSourcesClientCreateOrUpdateResponse struct {
	// Represents a datasource definition, which can be used to configure an indexer.
	IndexerDataSource
}

// DataSourcesClientCreateResponse contains the response from method DataSourcesClient.Create.
type DataSourcesClientCreateResponse struct {
	// Represents a datasource definition, which can be used to configure an indexer.
	IndexerDataSource
}

// DataSourcesClientDeleteResponse contains the response from method DataSourcesClient.Delete.
type DataSourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// DataSourcesClientGetResponse contains the response from method DataSourcesClient.Get.
type DataSourcesClientGetResponse struct {
	// Represents a datasource definition, which can be used to configure an indexer.
	IndexerDataSource
}

// DataSourcesClientListResponse contains the response from method DataSourcesClient.List.
type DataSourcesClientListResponse struct {
	// Response from a List Datasources request. If successful, it includes the full definitions of all datasources.
	ListDataSourcesResult
}

// IndexersClientCreateOrUpdateResponse contains the response from method IndexersClient.CreateOrUpdate.
type IndexersClientCreateOrUpdateResponse struct {
	// Represents an indexer.
	Indexer
}

// IndexersClientCreateResponse contains the response from method IndexersClient.Create.
type IndexersClientCreateResponse struct {
	// Represents an indexer.
	Indexer
}

// IndexersClientDeleteResponse contains the response from method IndexersClient.Delete.
type IndexersClientDeleteResponse struct {
	// placeholder for future response values
}

// IndexersClientGetResponse contains the response from method IndexersClient.Get.
type IndexersClientGetResponse struct {
	// Represents an indexer.
	Indexer
}

// IndexersClientGetStatusResponse contains the response from method IndexersClient.GetStatus.
type IndexersClientGetStatusResponse struct {
	// Represents the current status and execution history of an indexer.
	IndexerStatus2
}

// IndexersClientListResponse contains the response from method IndexersClient.List.
type IndexersClientListResponse struct {
	// Response from a List Indexers request. If successful, it includes the full definitions of all indexers.
	ListIndexersResult
}

// IndexersClientResetResponse contains the response from method IndexersClient.Reset.
type IndexersClientResetResponse struct {
	// placeholder for future response values
}

// IndexersClientRunResponse contains the response from method IndexersClient.Run.
type IndexersClientRunResponse struct {
	// placeholder for future response values
}

// IndexesClientAnalyzeResponse contains the response from method IndexesClient.Analyze.
type IndexesClientAnalyzeResponse struct {
	// The result of testing an analyzer on text.
	AnalyzeResult
}

// IndexesClientCreateOrUpdateResponse contains the response from method IndexesClient.CreateOrUpdate.
type IndexesClientCreateOrUpdateResponse struct {
	// Represents a search index definition, which describes the fields and search behavior of an index.
	Index
}

// IndexesClientCreateResponse contains the response from method IndexesClient.Create.
type IndexesClientCreateResponse struct {
	// Represents a search index definition, which describes the fields and search behavior of an index.
	Index
}

// IndexesClientDeleteResponse contains the response from method IndexesClient.Delete.
type IndexesClientDeleteResponse struct {
	// placeholder for future response values
}

// IndexesClientGetResponse contains the response from method IndexesClient.Get.
type IndexesClientGetResponse struct {
	// Represents a search index definition, which describes the fields and search behavior of an index.
	Index
}

// IndexesClientGetStatisticsResponse contains the response from method IndexesClient.GetStatistics.
type IndexesClientGetStatisticsResponse struct {
	// Statistics for a given index. Statistics are collected periodically and are not guaranteed to always be up-to-date.
	GetIndexStatisticsResult
}

// IndexesClientListResponse contains the response from method IndexesClient.NewListPager.
type IndexesClientListResponse struct {
	// Response from a List Indexes request. If successful, it includes the full definitions of all indexes.
	ListIndexesResult
}

// ServiceClientGetServiceStatisticsResponse contains the response from method ServiceClient.GetServiceStatistics.
type ServiceClientGetServiceStatisticsResponse struct {
	// Response from a get service statistics request. If successful, it includes service level counters and limits.
	ServiceStatistics
}

// SkillsetsClientCreateOrUpdateResponse contains the response from method SkillsetsClient.CreateOrUpdate.
type SkillsetsClientCreateOrUpdateResponse struct {
	// A list of skills.
	IndexerSkillset
}

// SkillsetsClientCreateResponse contains the response from method SkillsetsClient.Create.
type SkillsetsClientCreateResponse struct {
	// A list of skills.
	IndexerSkillset
}

// SkillsetsClientDeleteResponse contains the response from method SkillsetsClient.Delete.
type SkillsetsClientDeleteResponse struct {
	// placeholder for future response values
}

// SkillsetsClientGetResponse contains the response from method SkillsetsClient.Get.
type SkillsetsClientGetResponse struct {
	// A list of skills.
	IndexerSkillset
}

// SkillsetsClientListResponse contains the response from method SkillsetsClient.List.
type SkillsetsClientListResponse struct {
	// Response from a list skillset request. If successful, it includes the full definitions of all skillsets.
	ListSkillsetsResult
}

// SynonymMapsClientCreateOrUpdateResponse contains the response from method SynonymMapsClient.CreateOrUpdate.
type SynonymMapsClientCreateOrUpdateResponse struct {
	// Represents a synonym map definition.
	SynonymMap
}

// SynonymMapsClientCreateResponse contains the response from method SynonymMapsClient.Create.
type SynonymMapsClientCreateResponse struct {
	// Represents a synonym map definition.
	SynonymMap
}

// SynonymMapsClientDeleteResponse contains the response from method SynonymMapsClient.Delete.
type SynonymMapsClientDeleteResponse struct {
	// placeholder for future response values
}

// SynonymMapsClientGetResponse contains the response from method SynonymMapsClient.Get.
type SynonymMapsClientGetResponse struct {
	// Represents a synonym map definition.
	SynonymMap
}

// SynonymMapsClientListResponse contains the response from method SynonymMapsClient.List.
type SynonymMapsClientListResponse struct {
	// Response from a List SynonymMaps request. If successful, it includes the full definitions of all synonym maps.
	ListSynonymMapsResult
}

