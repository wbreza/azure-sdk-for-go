//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedRuleSetsClient contains the methods for the ManagedRuleSets group.
// Don't use this type directly, use NewManagedRuleSetsClient() instead.
type ManagedRuleSetsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedRuleSetsClient creates a new instance of ManagedRuleSetsClient with the specified values.
func NewManagedRuleSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedRuleSetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagedRuleSetsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Lists all available managed rule sets.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ManagedRuleSetsClient) List(options *ManagedRuleSetsListOptions) *ManagedRuleSetsListPager {
	return &ManagedRuleSetsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ManagedRuleSetsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedRuleSetDefinitionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ManagedRuleSetsClient) listCreateRequest(ctx context.Context, options *ManagedRuleSetsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedRuleSetsClient) listHandleResponse(resp *http.Response) (ManagedRuleSetsListResponse, error) {
	result := ManagedRuleSetsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedRuleSetDefinitionList); err != nil {
		return ManagedRuleSetsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ManagedRuleSetsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
