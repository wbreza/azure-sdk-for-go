package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ContainerLogsClient is the client for the ContainerLogs methods of the Containerinstance service.
type ContainerLogsClient struct {
	BaseClient
}

// NewContainerLogsClient creates an instance of the ContainerLogsClient client.
func NewContainerLogsClient(subscriptionID string) ContainerLogsClient {
	return NewContainerLogsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerLogsClientWithBaseURI creates an instance of the ContainerLogsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewContainerLogsClientWithBaseURI(baseURI string, subscriptionID string) ContainerLogsClient {
	return ContainerLogsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get the logs for a specified container instance in a specified resource group and container group.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerGroupName - the name of the container group.
// containerName - the name of the container instance.
// tail - the number of lines to show from the tail of the container instance log. If not provided, all
// available logs are shown up to 4mb.
func (client ContainerLogsClient) List(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (result Logs, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerLogsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, containerGroupName, containerName, tail)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerLogsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerLogsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.ContainerLogsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ContainerLogsClient) ListPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if tail != nil {
		queryParameters["tail"] = autorest.Encode("query", *tail)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerLogsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContainerLogsClient) ListResponder(resp *http.Response) (result Logs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}