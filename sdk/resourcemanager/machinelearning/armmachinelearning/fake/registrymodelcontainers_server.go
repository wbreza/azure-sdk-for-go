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
)

// RegistryModelContainersServer is a fake server for instances of the armmachinelearning.RegistryModelContainersClient type.
type RegistryModelContainersServer struct{
	// BeginCreateOrUpdate is the fake for method RegistryModelContainersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, registryName string, modelName string, body armmachinelearning.ModelContainer, options *armmachinelearning.RegistryModelContainersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RegistryModelContainersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *armmachinelearning.RegistryModelContainersClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegistryModelContainersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *armmachinelearning.RegistryModelContainersClientGetOptions) (resp azfake.Responder[armmachinelearning.RegistryModelContainersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RegistryModelContainersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, options *armmachinelearning.RegistryModelContainersClientListOptions) (resp azfake.PagerResponder[armmachinelearning.RegistryModelContainersClientListResponse])

}

// NewRegistryModelContainersServerTransport creates a new instance of RegistryModelContainersServerTransport with the provided implementation.
// The returned RegistryModelContainersServerTransport instance is connected to an instance of armmachinelearning.RegistryModelContainersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistryModelContainersServerTransport(srv *RegistryModelContainersServer) *RegistryModelContainersServerTransport {
	return &RegistryModelContainersServerTransport{
		srv: srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientCreateOrUpdateResponse]](),
		beginDelete: newTracker[azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.RegistryModelContainersClientListResponse]](),
	}
}

// RegistryModelContainersServerTransport connects instances of armmachinelearning.RegistryModelContainersClient to instances of RegistryModelContainersServer.
// Don't use this type directly, use NewRegistryModelContainersServerTransport instead.
type RegistryModelContainersServerTransport struct {
	srv *RegistryModelContainersServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientCreateOrUpdateResponse]]
	beginDelete *tracker[azfake.PollerResponder[armmachinelearning.RegistryModelContainersClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.RegistryModelContainersClientListResponse]]
}

// Do implements the policy.Transporter interface for RegistryModelContainersServerTransport.
func (r *RegistryModelContainersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegistryModelContainersClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RegistryModelContainersClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RegistryModelContainersClient.Get":
		resp, err = r.dispatchGet(req)
	case "RegistryModelContainersClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistryModelContainersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.ModelContainer](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RegistryModelContainersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RegistryModelContainersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ModelContainer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistryModelContainersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
	if err != nil {
		return nil, err
	}
	skipParam := getOptional(skipUnescaped)
	listViewTypeUnescaped, err := url.QueryUnescape(qp.Get("listViewType"))
	if err != nil {
		return nil, err
	}
	listViewTypeParam := getOptional(armmachinelearning.ListViewType(listViewTypeUnescaped))
	var options *armmachinelearning.RegistryModelContainersClientListOptions
	if skipParam != nil || listViewTypeParam != nil {
		options = &armmachinelearning.RegistryModelContainersClientListOptions{
			Skip: skipParam,
			ListViewType: listViewTypeParam,
		}
	}
resp := r.srv.NewListPager(resourceGroupNameParam, registryNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.RegistryModelContainersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

