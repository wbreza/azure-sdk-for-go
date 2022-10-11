package security

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// GovernanceRulesClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type GovernanceRulesClient struct {
	BaseClient
}

// NewGovernanceRulesClient creates an instance of the GovernanceRulesClient client.
func NewGovernanceRulesClient(subscriptionID string) GovernanceRulesClient {
	return NewGovernanceRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGovernanceRulesClientWithBaseURI creates an instance of the GovernanceRulesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGovernanceRulesClientWithBaseURI(baseURI string, subscriptionID string) GovernanceRulesClient {
	return GovernanceRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or update a security GovernanceRule on the given subscription.
// Parameters:
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
// governanceRule - governanceRule over a subscription scope
func (client GovernanceRulesClient) CreateOrUpdate(ctx context.Context, ruleID string, governanceRule GovernanceRule) (result GovernanceRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GovernanceRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: governanceRule,
			Constraints: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
							{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
					{Target: "governanceRule.GovernanceRuleProperties.SourceResourceType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.ConditionSets", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.OwnerSource", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("security.GovernanceRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, ruleID, governanceRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GovernanceRulesClient) CreateOrUpdatePreparer(ctx context.Context, ruleID string, governanceRule GovernanceRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ruleId":         autorest.Encode("path", ruleID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithJSON(governanceRule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GovernanceRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GovernanceRulesClient) CreateOrUpdateResponder(resp *http.Response) (result GovernanceRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a GovernanceRule over a given scope
// Parameters:
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
func (client GovernanceRulesClient) Delete(ctx context.Context, ruleID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GovernanceRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.GovernanceRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GovernanceRulesClient) DeletePreparer(ctx context.Context, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ruleId":         autorest.Encode("path", ruleID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GovernanceRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GovernanceRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a specific governanceRule for the requested scope by ruleId
// Parameters:
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
func (client GovernanceRulesClient) Get(ctx context.Context, ruleID string) (result GovernanceRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GovernanceRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.GovernanceRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client GovernanceRulesClient) GetPreparer(ctx context.Context, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ruleId":         autorest.Encode("path", ruleID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GovernanceRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GovernanceRulesClient) GetResponder(resp *http.Response) (result GovernanceRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RuleIDExecuteSingleSecurityConnector execute a security GovernanceRule on the given security connector.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// securityConnectorName - the security connector name.
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
// executeGovernanceRuleParams - governanceRule over a subscription scope
func (client GovernanceRulesClient) RuleIDExecuteSingleSecurityConnector(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, executeGovernanceRuleParams *ExecuteGovernanceRuleParams) (result GovernanceRulesRuleIDExecuteSingleSecurityConnectorFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GovernanceRulesClient.RuleIDExecuteSingleSecurityConnector")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.GovernanceRulesClient", "RuleIDExecuteSingleSecurityConnector", err.Error())
	}

	req, err := client.RuleIDExecuteSingleSecurityConnectorPreparer(ctx, resourceGroupName, securityConnectorName, ruleID, executeGovernanceRuleParams)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "RuleIDExecuteSingleSecurityConnector", nil, "Failure preparing request")
		return
	}

	result, err = client.RuleIDExecuteSingleSecurityConnectorSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "RuleIDExecuteSingleSecurityConnector", result.Response(), "Failure sending request")
		return
	}

	return
}

// RuleIDExecuteSingleSecurityConnectorPreparer prepares the RuleIDExecuteSingleSecurityConnector request.
func (client GovernanceRulesClient) RuleIDExecuteSingleSecurityConnectorPreparer(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, executeGovernanceRuleParams *ExecuteGovernanceRuleParams) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"ruleId":                autorest.Encode("path", ruleID),
		"securityConnectorName": autorest.Encode("path", securityConnectorName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}/execute", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if executeGovernanceRuleParams != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(executeGovernanceRuleParams))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RuleIDExecuteSingleSecurityConnectorSender sends the RuleIDExecuteSingleSecurityConnector request. The method will close the
// http.Response Body if it receives an error.
func (client GovernanceRulesClient) RuleIDExecuteSingleSecurityConnectorSender(req *http.Request) (future GovernanceRulesRuleIDExecuteSingleSecurityConnectorFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RuleIDExecuteSingleSecurityConnectorResponder handles the response to the RuleIDExecuteSingleSecurityConnector request. The method always
// closes the http.Response Body.
func (client GovernanceRulesClient) RuleIDExecuteSingleSecurityConnectorResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RuleIDExecuteSingleSubscription execute a security GovernanceRule on the given subscription.
// Parameters:
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
// executeGovernanceRuleParams - governanceRule over a subscription scope
func (client GovernanceRulesClient) RuleIDExecuteSingleSubscription(ctx context.Context, ruleID string, executeGovernanceRuleParams *ExecuteGovernanceRuleParams) (result GovernanceRulesRuleIDExecuteSingleSubscriptionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GovernanceRulesClient.RuleIDExecuteSingleSubscription")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.GovernanceRulesClient", "RuleIDExecuteSingleSubscription", err.Error())
	}

	req, err := client.RuleIDExecuteSingleSubscriptionPreparer(ctx, ruleID, executeGovernanceRuleParams)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "RuleIDExecuteSingleSubscription", nil, "Failure preparing request")
		return
	}

	result, err = client.RuleIDExecuteSingleSubscriptionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.GovernanceRulesClient", "RuleIDExecuteSingleSubscription", result.Response(), "Failure sending request")
		return
	}

	return
}

// RuleIDExecuteSingleSubscriptionPreparer prepares the RuleIDExecuteSingleSubscription request.
func (client GovernanceRulesClient) RuleIDExecuteSingleSubscriptionPreparer(ctx context.Context, ruleID string, executeGovernanceRuleParams *ExecuteGovernanceRuleParams) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ruleId":         autorest.Encode("path", ruleID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules/{ruleId}/execute", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if executeGovernanceRuleParams != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(executeGovernanceRuleParams))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RuleIDExecuteSingleSubscriptionSender sends the RuleIDExecuteSingleSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client GovernanceRulesClient) RuleIDExecuteSingleSubscriptionSender(req *http.Request) (future GovernanceRulesRuleIDExecuteSingleSubscriptionFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RuleIDExecuteSingleSubscriptionResponder handles the response to the RuleIDExecuteSingleSubscription request. The method always
// closes the http.Response Body.
func (client GovernanceRulesClient) RuleIDExecuteSingleSubscriptionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}