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

// ManagedNetworkProvisionsClient contains the methods for the ManagedNetworkProvisions group.
// Don't use this type directly, use NewManagedNetworkProvisionsClient() instead.
type ManagedNetworkProvisionsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewManagedNetworkProvisionsClient creates a new instance of ManagedNetworkProvisionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedNetworkProvisionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedNetworkProvisionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedNetworkProvisionsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// BeginProvisionManagedNetwork - Provisions the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - options - ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions contains the optional parameters for the ManagedNetworkProvisionsClient.BeginProvisionManagedNetwork
//     method.
func (client *ManagedNetworkProvisionsClient) BeginProvisionManagedNetwork(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions) (*runtime.Poller[ManagedNetworkProvisionsClientProvisionManagedNetworkResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.provisionManagedNetwork(ctx, resourceGroupName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedNetworkProvisionsClientProvisionManagedNetworkResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedNetworkProvisionsClientProvisionManagedNetworkResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ProvisionManagedNetwork - Provisions the managed network of a machine learning workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *ManagedNetworkProvisionsClient) provisionManagedNetwork(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedNetworkProvisionsClient.BeginProvisionManagedNetwork"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.provisionManagedNetworkCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// provisionManagedNetworkCreateRequest creates the ProvisionManagedNetwork request.
func (client *ManagedNetworkProvisionsClient) provisionManagedNetworkCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/provisionManagedNetwork"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
	if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
	return nil, err
}
		return req, nil
	}
	return req, nil
}

