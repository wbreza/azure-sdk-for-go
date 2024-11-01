// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsearch

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ServiceClient contains the methods for the SearchServiceClient group.
// Don't use this type directly, use a constructor function instead.
type ServiceClient struct {
	internal *azcore.Client
	endpoint string
}

// GetServiceStatistics - Gets service level statistics for a search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - RequestOptions - RequestOptions contains a group of parameters for the DataSourcesClient.CreateOrUpdate method.
//   - options - ServiceClientGetServiceStatisticsOptions contains the optional parameters for the ServiceClient.GetServiceStatistics
//     method.
func (client *ServiceClient) GetServiceStatistics(ctx context.Context, requestOptions *RequestOptions, options *ServiceClientGetServiceStatisticsOptions) (ServiceClientGetServiceStatisticsResponse, error) {
	var err error
	req, err := client.getServiceStatisticsCreateRequest(ctx, requestOptions, options)
	if err != nil {
		return ServiceClientGetServiceStatisticsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetServiceStatisticsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceClientGetServiceStatisticsResponse{}, err
	}
	resp, err := client.getServiceStatisticsHandleResponse(httpResp)
	return resp, err
}

// getServiceStatisticsCreateRequest creates the GetServiceStatistics request.
func (client *ServiceClient) getServiceStatisticsCreateRequest(ctx context.Context, requestOptions *RequestOptions, _ *ServiceClientGetServiceStatisticsOptions) (*policy.Request, error) {
	urlPath := "/servicestats"
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

// getServiceStatisticsHandleResponse handles the GetServiceStatistics response.
func (client *ServiceClient) getServiceStatisticsHandleResponse(resp *http.Response) (ServiceClientGetServiceStatisticsResponse, error) {
	result := ServiceClientGetServiceStatisticsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceStatistics); err != nil {
		return ServiceClientGetServiceStatisticsResponse{}, err
	}
	return result, nil
}

