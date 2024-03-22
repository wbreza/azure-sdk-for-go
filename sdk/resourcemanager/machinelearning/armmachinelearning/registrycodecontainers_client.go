//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegistryCodeContainersClient contains the methods for the RegistryCodeContainers group.
// Don't use this type directly, use NewRegistryCodeContainersClient() instead.
type RegistryCodeContainersClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewRegistryCodeContainersClient creates a new instance of RegistryCodeContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryCodeContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryCodeContainersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryCodeContainersClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Code container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - codeName - Container name.
//   - body - Container entity to create or update.
//   - options - RegistryCodeContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryCodeContainersClient.BeginCreateOrUpdate
//     method.
func (client *RegistryCodeContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, codeName string, body CodeContainer, options *RegistryCodeContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryCodeContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, codeName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryCodeContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryCodeContainersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update Code container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *RegistryCodeContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, codeName string, body CodeContainer, options *RegistryCodeContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryCodeContainersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, codeName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistryCodeContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, body CodeContainer, options *RegistryCodeContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Delete Code container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - codeName - Container name.
//   - options - RegistryCodeContainersClientBeginDeleteOptions contains the optional parameters for the RegistryCodeContainersClient.BeginDelete
//     method.
func (client *RegistryCodeContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeContainersClientBeginDeleteOptions) (*runtime.Poller[RegistryCodeContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, codeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryCodeContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryCodeContainersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete Code container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *RegistryCodeContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeContainersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryCodeContainersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, codeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistryCodeContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Code container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - codeName - Container name.
//   - options - RegistryCodeContainersClientGetOptions contains the optional parameters for the RegistryCodeContainersClient.Get
//     method.
func (client *RegistryCodeContainersClient) Get(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeContainersClientGetOptions) (RegistryCodeContainersClientGetResponse, error) {
	var err error
	const operationName = "RegistryCodeContainersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, codeName, options)
	if err != nil {
		return RegistryCodeContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryCodeContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegistryCodeContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegistryCodeContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, codeName string, options *RegistryCodeContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes/{codeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if codeName == "" {
		return nil, errors.New("parameter codeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{codeName}", url.PathEscape(codeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistryCodeContainersClient) getHandleResponse(resp *http.Response) (RegistryCodeContainersClientGetResponse, error) {
	result := RegistryCodeContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeContainer); err != nil {
		return RegistryCodeContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List containers.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - options - RegistryCodeContainersClientListOptions contains the optional parameters for the RegistryCodeContainersClient.NewListPager
//     method.
func (client *RegistryCodeContainersClient) NewListPager(resourceGroupName string, registryName string, options *RegistryCodeContainersClientListOptions) (*runtime.Pager[RegistryCodeContainersClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RegistryCodeContainersClientListResponse]{
		More: func(page RegistryCodeContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryCodeContainersClientListResponse) (RegistryCodeContainersClientListResponse, error) {
		ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegistryCodeContainersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			}, nil)
			if err != nil {
				return RegistryCodeContainersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
			},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RegistryCodeContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistryCodeContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/codes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryCodeContainersClient) listHandleResponse(resp *http.Response) (RegistryCodeContainersClientListResponse, error) {
	result := RegistryCodeContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeContainerResourceArmPaginatedResult); err != nil {
		return RegistryCodeContainersClientListResponse{}, err
	}
	return result, nil
}

