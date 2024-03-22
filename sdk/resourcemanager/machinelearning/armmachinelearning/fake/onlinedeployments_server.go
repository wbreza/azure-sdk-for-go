//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/wbreza/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// OnlineDeploymentsServer is a fake server for instances of the armmachinelearning.OnlineDeploymentsClient type.
type OnlineDeploymentsServer struct{
	// BeginCreateOrUpdate is the fake for method OnlineDeploymentsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body armmachinelearning.OnlineDeployment, options *armmachinelearning.OnlineDeploymentsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OnlineDeploymentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *armmachinelearning.OnlineDeploymentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OnlineDeploymentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *armmachinelearning.OnlineDeploymentsClientGetOptions) (resp azfake.Responder[armmachinelearning.OnlineDeploymentsClientGetResponse], errResp azfake.ErrorResponder)

	// GetLogs is the fake for method OnlineDeploymentsClient.GetLogs
	// HTTP status codes to indicate success: http.StatusOK
	GetLogs func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body armmachinelearning.DeploymentLogsRequest, options *armmachinelearning.OnlineDeploymentsClientGetLogsOptions) (resp azfake.Responder[armmachinelearning.OnlineDeploymentsClientGetLogsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method OnlineDeploymentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, endpointName string, options *armmachinelearning.OnlineDeploymentsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListResponse])

	// NewListSKUsPager is the fake for method OnlineDeploymentsClient.NewListSKUsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSKUsPager func(resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *armmachinelearning.OnlineDeploymentsClientListSKUsOptions) (resp azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListSKUsResponse])

	// BeginUpdate is the fake for method OnlineDeploymentsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body armmachinelearning.PartialMinimalTrackedResourceWithSKU, options *armmachinelearning.OnlineDeploymentsClientBeginUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientUpdateResponse], errResp azfake.ErrorResponder)

}

// NewOnlineDeploymentsServerTransport creates a new instance of OnlineDeploymentsServerTransport with the provided implementation.
// The returned OnlineDeploymentsServerTransport instance is connected to an instance of armmachinelearning.OnlineDeploymentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOnlineDeploymentsServerTransport(srv *OnlineDeploymentsServer) *OnlineDeploymentsServerTransport {
	return &OnlineDeploymentsServerTransport{
		srv: srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientCreateOrUpdateResponse]](),
		beginDelete: newTracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListResponse]](),
		newListSKUsPager: newTracker[azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListSKUsResponse]](),
		beginUpdate: newTracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientUpdateResponse]](),
	}
}

// OnlineDeploymentsServerTransport connects instances of armmachinelearning.OnlineDeploymentsClient to instances of OnlineDeploymentsServer.
// Don't use this type directly, use NewOnlineDeploymentsServerTransport instead.
type OnlineDeploymentsServerTransport struct {
	srv *OnlineDeploymentsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientCreateOrUpdateResponse]]
	beginDelete *tracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListResponse]]
	newListSKUsPager *tracker[azfake.PagerResponder[armmachinelearning.OnlineDeploymentsClientListSKUsResponse]]
	beginUpdate *tracker[azfake.PollerResponder[armmachinelearning.OnlineDeploymentsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for OnlineDeploymentsServerTransport.
func (o *OnlineDeploymentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OnlineDeploymentsClient.BeginCreateOrUpdate":
		resp, err = o.dispatchBeginCreateOrUpdate(req)
	case "OnlineDeploymentsClient.BeginDelete":
		resp, err = o.dispatchBeginDelete(req)
	case "OnlineDeploymentsClient.Get":
		resp, err = o.dispatchGet(req)
	case "OnlineDeploymentsClient.GetLogs":
		resp, err = o.dispatchGetLogs(req)
	case "OnlineDeploymentsClient.NewListPager":
		resp, err = o.dispatchNewListPager(req)
	case "OnlineDeploymentsClient.NewListSKUsPager":
		resp, err = o.dispatchNewListSKUsPager(req)
	case "OnlineDeploymentsClient.BeginUpdate":
		resp, err = o.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := o.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.OnlineDeployment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginCreateOrUpdate = &respr
		o.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		o.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		o.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OnlineDeployment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchGetLogs(req *http.Request) (*http.Response, error) {
	if o.srv.GetLogs == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetLogs not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getLogs`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.DeploymentLogsRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.GetLogs(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeploymentLogs, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderBy"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
	if err != nil {
		return nil, err
	}
	skipParam := getOptional(skipUnescaped)
	var options *armmachinelearning.OnlineDeploymentsClientListOptions
	if orderByParam != nil || topParam != nil || skipParam != nil {
		options = &armmachinelearning.OnlineDeploymentsClientListOptions{
			OrderBy: orderByParam,
			Top: topParam,
			Skip: skipParam,
		}
	}
resp := o.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, endpointNameParam, options)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.OnlineDeploymentsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchNewListSKUsPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListSKUsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSKUsPager not implemented")}
	}
	newListSKUsPager := o.newListSKUsPager.get(req)
	if newListSKUsPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	countUnescaped, err := url.QueryUnescape(qp.Get("count"))
	if err != nil {
		return nil, err
	}
	countParam, err := parseOptional(countUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
	if err != nil {
		return nil, err
	}
	skipParam := getOptional(skipUnescaped)
	var options *armmachinelearning.OnlineDeploymentsClientListSKUsOptions
	if countParam != nil || skipParam != nil {
		options = &armmachinelearning.OnlineDeploymentsClientListSKUsOptions{
			Count: countParam,
			Skip: skipParam,
		}
	}
resp := o.srv.NewListSKUsPager(resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, options)
		newListSKUsPager = &resp
		o.newListSKUsPager.add(req, newListSKUsPager)
		server.PagerResponderInjectNextLinks(newListSKUsPager, req, func(page *armmachinelearning.OnlineDeploymentsClientListSKUsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSKUsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListSKUsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSKUsPager) {
		o.newListSKUsPager.remove(req)
	}
	return resp, nil
}

func (o *OnlineDeploymentsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := o.beginUpdate.get(req)
	if beginUpdate == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/onlineEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deployments/(?P<deploymentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.PartialMinimalTrackedResourceWithSKU](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	deploymentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deploymentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.BeginUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, deploymentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginUpdate = &respr
		o.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		o.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		o.beginUpdate.remove(req)
	}

	return resp, nil
}

