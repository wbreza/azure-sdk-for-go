//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armelastic

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
	"strings"
)

// TagRulesClient contains the methods for the TagRules group.
// Don't use this type directly, use NewTagRulesClient() instead.
type TagRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTagRulesClient creates a new instance of TagRulesClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTagRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagRulesClient, error) {
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
	client := &TagRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// ruleSetName - Tag Rule Set resource name
// options - TagRulesClientCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.CreateOrUpdate method.
func (client *TagRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientCreateOrUpdateOptions) (TagRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagRulesClient) createOrUpdateHandleResponse(resp *http.Response) (TagRulesClientCreateOrUpdateResponse, error) {
	result := TagRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// ruleSetName - Tag Rule Set resource name
// options - TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
func (client *TagRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*runtime.Poller[TagRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, ruleSetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[TagRulesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[TagRulesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-07-01-preview
func (client *TagRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// ruleSetName - Tag Rule Set resource name
// options - TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
func (client *TagRulesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (TagRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagRulesClient) getHandleResponse(resp *http.Response) (TagRulesClientGetResponse, error) {
	result := TagRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the tag rules for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// options - TagRulesClientListOptions contains the optional parameters for the TagRulesClient.List method.
func (client *TagRulesClient) NewListPager(resourceGroupName string, monitorName string, options *TagRulesClientListOptions) *runtime.Pager[TagRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagRulesClientListResponse]{
		More: func(page TagRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagRulesClientListResponse) (TagRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TagRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TagRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TagRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TagRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *TagRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/tagRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagRulesClient) listHandleResponse(resp *http.Response) (TagRulesClientListResponse, error) {
	result := TagRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRulesListResponse); err != nil {
		return TagRulesClientListResponse{}, err
	}
	return result, nil
}