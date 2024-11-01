// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsearch

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SkillsetsClient contains the methods for the Skillsets group.
// Don't use this type directly, use a constructor function instead.
type SkillsetsClient struct {
	internal *azcore.Client
	endpoint string
}

// Create - Creates a new skillset in a search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - skillset - The skillset containing one or more skills to create in a search service.
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
//   - options - SkillsetsClientCreateOptions contains the optional parameters for the SkillsetsClient.Create method.
func (client *SkillsetsClient) Create(ctx context.Context, skillset IndexerSkillset, requestOptions *RequestOptions, options *SkillsetsClientCreateOptions) (SkillsetsClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, skillset, requestOptions, options)
	if err != nil {
		return SkillsetsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SkillsetsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SkillsetsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SkillsetsClient) createCreateRequest(ctx context.Context, skillset IndexerSkillset, requestOptions *RequestOptions, _ *SkillsetsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/skillsets"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if requestOptions != nil && requestOptions.XMSClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*requestOptions.XMSClientRequestID}
	}
	if err := runtime.MarshalAsJSON(req, skillset); err != nil {
	return nil, err
}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SkillsetsClient) createHandleResponse(resp *http.Response) (SkillsetsClientCreateResponse, error) {
	result := SkillsetsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IndexerSkillset); err != nil {
		return SkillsetsClientCreateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Creates a new skillset in a search service or updates the skillset if it already exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - skillsetName - The name of the skillset to create or update.
//   - prefer - For HTTP PUT requests, instructs the service to return the created/updated resource on success.
//   - skillset - The skillset containing one or more skills to create or update in a search service.
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
//   - options - SkillsetsClientCreateOrUpdateOptions contains the optional parameters for the SkillsetsClient.CreateOrUpdate
//     method.
func (client *SkillsetsClient) CreateOrUpdate(ctx context.Context, skillsetName string, prefer Enum0, skillset IndexerSkillset, requestOptions *RequestOptions, options *SkillsetsClientCreateOrUpdateOptions) (SkillsetsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, skillsetName, prefer, skillset, requestOptions, options)
	if err != nil {
		return SkillsetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SkillsetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SkillsetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SkillsetsClient) createOrUpdateCreateRequest(ctx context.Context, skillsetName string, prefer Enum0, skillset IndexerSkillset, requestOptions *RequestOptions, options *SkillsetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/skillsets('{skillsetName}')"
	if skillsetName == "" {
		return nil, errors.New("parameter skillsetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skillsetName}", url.PathEscape(skillsetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Prefer"] = []string{string(prefer)}
	if requestOptions != nil && requestOptions.XMSClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*requestOptions.XMSClientRequestID}
	}
	if err := runtime.MarshalAsJSON(req, skillset); err != nil {
	return nil, err
}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SkillsetsClient) createOrUpdateHandleResponse(resp *http.Response) (SkillsetsClientCreateOrUpdateResponse, error) {
	result := SkillsetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IndexerSkillset); err != nil {
		return SkillsetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a skillset in a search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - skillsetName - The name of the skillset to delete.
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
//   - options - SkillsetsClientDeleteOptions contains the optional parameters for the SkillsetsClient.Delete method.
func (client *SkillsetsClient) Delete(ctx context.Context, skillsetName string, requestOptions *RequestOptions, options *SkillsetsClientDeleteOptions) (SkillsetsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, skillsetName, requestOptions, options)
	if err != nil {
		return SkillsetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SkillsetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return SkillsetsClientDeleteResponse{}, err
	}
	return SkillsetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SkillsetsClient) deleteCreateRequest(ctx context.Context, skillsetName string, requestOptions *RequestOptions, options *SkillsetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/skillsets('{skillsetName}')"
	if skillsetName == "" {
		return nil, errors.New("parameter skillsetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skillsetName}", url.PathEscape(skillsetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if requestOptions != nil && requestOptions.XMSClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*requestOptions.XMSClientRequestID}
	}
	return req, nil
}

// Get - Retrieves a skillset in a search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - skillsetName - The name of the skillset to retrieve.
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
//   - options - SkillsetsClientGetOptions contains the optional parameters for the SkillsetsClient.Get method.
func (client *SkillsetsClient) Get(ctx context.Context, skillsetName string, requestOptions *RequestOptions, options *SkillsetsClientGetOptions) (SkillsetsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, skillsetName, requestOptions, options)
	if err != nil {
		return SkillsetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SkillsetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SkillsetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SkillsetsClient) getCreateRequest(ctx context.Context, skillsetName string, requestOptions *RequestOptions, _ *SkillsetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/skillsets('{skillsetName}')"
	if skillsetName == "" {
		return nil, errors.New("parameter skillsetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skillsetName}", url.PathEscape(skillsetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if requestOptions != nil && requestOptions.XMSClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*requestOptions.XMSClientRequestID}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SkillsetsClient) getHandleResponse(resp *http.Response) (SkillsetsClientGetResponse, error) {
	result := SkillsetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IndexerSkillset); err != nil {
		return SkillsetsClientGetResponse{}, err
	}
	return result, nil
}

// List - List all skillsets in a search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - options - SkillsetsClientListOptions contains the optional parameters for the SkillsetsClient.List method.
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
func (client *SkillsetsClient) List(ctx context.Context, options *SkillsetsClientListOptions, requestOptions *RequestOptions) (SkillsetsClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, options, requestOptions)
	if err != nil {
		return SkillsetsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SkillsetsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SkillsetsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *SkillsetsClient) listCreateRequest(ctx context.Context, options *SkillsetsClientListOptions, requestOptions *RequestOptions) (*policy.Request, error) {
	urlPath := "/skillsets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if requestOptions != nil && requestOptions.XMSClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*requestOptions.XMSClientRequestID}
	}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SkillsetsClient) listHandleResponse(resp *http.Response) (SkillsetsClientListResponse, error) {
	result := SkillsetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListSkillsetsResult); err != nil {
		return SkillsetsClientListResponse{}, err
	}
	return result, nil
}

