//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PipelinesClient contains the methods for the Pipelines group.
// Don't use this type directly, use NewPipelinesClient() instead.
type PipelinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPipelinesClient creates a new instance of PipelinesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelinesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PipelinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// pipelineName - The pipeline name.
// pipeline - Pipeline resource definition.
// options - PipelinesClientCreateOrUpdateOptions contains the optional parameters for the PipelinesClient.CreateOrUpdate
// method.
func (client *PipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesClientCreateOrUpdateOptions) (PipelinesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, pipeline, options)
	if err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PipelinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, pipeline)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PipelinesClient) createOrUpdateHandleResponse(resp *http.Response) (PipelinesClientCreateOrUpdateResponse, error) {
	result := PipelinesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateRun - Creates a run of a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// pipelineName - The pipeline name.
// options - PipelinesClientCreateRunOptions contains the optional parameters for the PipelinesClient.CreateRun method.
func (client *PipelinesClient) CreateRun(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientCreateRunOptions) (PipelinesClientCreateRunResponse, error) {
	req, err := client.createRunCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesClientCreateRunResponse{}, runtime.NewResponseError(resp)
	}
	return client.createRunHandleResponse(resp)
}

// createRunCreateRequest creates the CreateRun request.
func (client *PipelinesClient) createRunCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientCreateRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}/createRun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	if options != nil && options.StartFromFailure != nil {
		reqQP.Set("startFromFailure", strconv.FormatBool(*options.StartFromFailure))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, options.Parameters)
	}
	return req, nil
}

// createRunHandleResponse handles the CreateRun response.
func (client *PipelinesClient) createRunHandleResponse(resp *http.Response) (PipelinesClientCreateRunResponse, error) {
	result := PipelinesClientCreateRunResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateRunResponse); err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// pipelineName - The pipeline name.
// options - PipelinesClientDeleteOptions contains the optional parameters for the PipelinesClient.Delete method.
func (client *PipelinesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientDeleteOptions) (PipelinesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PipelinesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PipelinesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// pipelineName - The pipeline name.
// options - PipelinesClientGetOptions contains the optional parameters for the PipelinesClient.Get method.
func (client *PipelinesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientGetOptions) (PipelinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return PipelinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelinesClient) getHandleResponse(resp *http.Response) (PipelinesClientGetResponse, error) {
	result := PipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists pipelines.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - PipelinesClientListByFactoryOptions contains the optional parameters for the PipelinesClient.ListByFactory method.
func (client *PipelinesClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *PipelinesClientListByFactoryOptions) *runtime.Pager[PipelinesClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelinesClientListByFactoryResponse]{
		More: func(page PipelinesClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelinesClientListByFactoryResponse) (PipelinesClientListByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelinesClientListByFactoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PipelinesClientListByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelinesClientListByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFactoryHandleResponse(resp)
		},
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *PipelinesClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *PipelinesClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *PipelinesClient) listByFactoryHandleResponse(resp *http.Response) (PipelinesClientListByFactoryResponse, error) {
	result := PipelinesClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResponse); err != nil {
		return PipelinesClientListByFactoryResponse{}, err
	}
	return result, nil
}