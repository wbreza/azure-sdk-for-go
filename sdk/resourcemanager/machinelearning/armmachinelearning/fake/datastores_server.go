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

// DatastoresServer is a fake server for instances of the armmachinelearning.DatastoresClient type.
type DatastoresServer struct{
	// CreateOrUpdate is the fake for method DatastoresClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, name string, body armmachinelearning.Datastore, options *armmachinelearning.DatastoresClientCreateOrUpdateOptions) (resp azfake.Responder[armmachinelearning.DatastoresClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DatastoresClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.DatastoresClientDeleteOptions) (resp azfake.Responder[armmachinelearning.DatastoresClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DatastoresClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.DatastoresClientGetOptions) (resp azfake.Responder[armmachinelearning.DatastoresClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DatastoresClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armmachinelearning.DatastoresClientListOptions) (resp azfake.PagerResponder[armmachinelearning.DatastoresClientListResponse])

	// ListSecrets is the fake for method DatastoresClient.ListSecrets
	// HTTP status codes to indicate success: http.StatusOK
	ListSecrets func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.DatastoresClientListSecretsOptions) (resp azfake.Responder[armmachinelearning.DatastoresClientListSecretsResponse], errResp azfake.ErrorResponder)

}

// NewDatastoresServerTransport creates a new instance of DatastoresServerTransport with the provided implementation.
// The returned DatastoresServerTransport instance is connected to an instance of armmachinelearning.DatastoresClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatastoresServerTransport(srv *DatastoresServer) *DatastoresServerTransport {
	return &DatastoresServerTransport{
		srv: srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.DatastoresClientListResponse]](),
	}
}

// DatastoresServerTransport connects instances of armmachinelearning.DatastoresClient to instances of DatastoresServer.
// Don't use this type directly, use NewDatastoresServerTransport instead.
type DatastoresServerTransport struct {
	srv *DatastoresServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.DatastoresClientListResponse]]
}

// Do implements the policy.Transporter interface for DatastoresServerTransport.
func (d *DatastoresServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DatastoresClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DatastoresClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DatastoresClient.Get":
		resp, err = d.dispatchGet(req)
	case "DatastoresClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DatastoresClient.ListSecrets":
		resp, err = d.dispatchListSecrets(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DatastoresServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.Datastore](req)
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	skipValidationUnescaped, err := url.QueryUnescape(qp.Get("skipValidation"))
	if err != nil {
		return nil, err
	}
	skipValidationParam, err := parseOptional(skipValidationUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armmachinelearning.DatastoresClientCreateOrUpdateOptions
	if skipValidationParam != nil {
		options = &armmachinelearning.DatastoresClientCreateOrUpdateOptions{
			SkipValidation: skipValidationParam,
		}
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Datastore, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatastoresServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatastoresServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Datastore, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatastoresServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores`
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
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
	if err != nil {
		return nil, err
	}
	skipParam := getOptional(skipUnescaped)
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
	isDefaultUnescaped, err := url.QueryUnescape(qp.Get("isDefault"))
	if err != nil {
		return nil, err
	}
	isDefaultParam, err := parseOptional(isDefaultUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	namesUnescaped, err := url.QueryUnescape(qp.Get("names"))
	if err != nil {
		return nil, err
	}
	namesParam := splitHelper(namesUnescaped, ",")
	searchTextUnescaped, err := url.QueryUnescape(qp.Get("searchText"))
	if err != nil {
		return nil, err
	}
	searchTextParam := getOptional(searchTextUnescaped)
	orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
	if err != nil {
		return nil, err
	}
	orderByParam := getOptional(orderByUnescaped)
	orderByAscUnescaped, err := url.QueryUnescape(qp.Get("orderByAsc"))
	if err != nil {
		return nil, err
	}
	orderByAscParam, err := parseOptional(orderByAscUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armmachinelearning.DatastoresClientListOptions
	if skipParam != nil || countParam != nil || isDefaultParam != nil || len(namesParam) > 0 || searchTextParam != nil || orderByParam != nil || orderByAscParam != nil {
		options = &armmachinelearning.DatastoresClientListOptions{
			Skip: skipParam,
			Count: countParam,
			IsDefault: isDefaultParam,
			Names: namesParam,
			SearchText: searchTextParam,
			OrderBy: orderByParam,
			OrderByAsc: orderByAscParam,
		}
	}
resp := d.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.DatastoresClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DatastoresServerTransport) dispatchListSecrets(req *http.Request) (*http.Response, error) {
	if d.srv.ListSecrets == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSecrets not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSecrets`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.ListSecrets(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatastoreSecretsClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

