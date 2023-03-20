//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azsecrets

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client contains the methods for the Client group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	endpoint string
	pl       runtime.Pipeline
}

// BackupSecret - Requests that a backup of the specified secret be downloaded to the client. All versions of the secret will
// be downloaded. This operation requires the secrets/backup permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// options - BackupSecretOptions contains the optional parameters for the Client.BackupSecret method.
func (client *Client) BackupSecret(ctx context.Context, name string, options *BackupSecretOptions) (BackupSecretResponse, error) {
	req, err := client.backupSecretCreateRequest(ctx, name, options)
	if err != nil {
		return BackupSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.backupSecretHandleResponse(resp)
}

// backupSecretCreateRequest creates the BackupSecret request.
func (client *Client) backupSecretCreateRequest(ctx context.Context, name string, options *BackupSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}/backup"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// backupSecretHandleResponse handles the BackupSecret response.
func (client *Client) backupSecretHandleResponse(resp *http.Response) (BackupSecretResponse, error) {
	result := BackupSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupSecretResult); err != nil {
		return BackupSecretResponse{}, err
	}
	return result, nil
}

// DeleteSecret - The DELETE operation applies to any secret stored in Azure Key Vault. DELETE cannot be applied to an individual
// version of a secret. This operation requires the secrets/delete permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// options - DeleteSecretOptions contains the optional parameters for the Client.DeleteSecret method.
func (client *Client) DeleteSecret(ctx context.Context, name string, options *DeleteSecretOptions) (DeleteSecretResponse, error) {
	req, err := client.deleteSecretCreateRequest(ctx, name, options)
	if err != nil {
		return DeleteSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeleteSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeleteSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteSecretHandleResponse(resp)
}

// deleteSecretCreateRequest creates the DeleteSecret request.
func (client *Client) deleteSecretCreateRequest(ctx context.Context, name string, options *DeleteSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteSecretHandleResponse handles the DeleteSecret response.
func (client *Client) deleteSecretHandleResponse(resp *http.Response) (DeleteSecretResponse, error) {
	result := DeleteSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSecretBundle); err != nil {
		return DeleteSecretResponse{}, err
	}
	return result, nil
}

// GetDeletedSecret - The Get Deleted Secret operation returns the specified deleted secret along with its attributes. This
// operation requires the secrets/get permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// options - GetDeletedSecretOptions contains the optional parameters for the Client.GetDeletedSecret method.
func (client *Client) GetDeletedSecret(ctx context.Context, name string, options *GetDeletedSecretOptions) (GetDeletedSecretResponse, error) {
	req, err := client.getDeletedSecretCreateRequest(ctx, name, options)
	if err != nil {
		return GetDeletedSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetDeletedSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetDeletedSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedSecretHandleResponse(resp)
}

// getDeletedSecretCreateRequest creates the GetDeletedSecret request.
func (client *Client) getDeletedSecretCreateRequest(ctx context.Context, name string, options *GetDeletedSecretOptions) (*policy.Request, error) {
	urlPath := "/deletedsecrets/{secret-name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeletedSecretHandleResponse handles the GetDeletedSecret response.
func (client *Client) getDeletedSecretHandleResponse(resp *http.Response) (GetDeletedSecretResponse, error) {
	result := GetDeletedSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSecretBundle); err != nil {
		return GetDeletedSecretResponse{}, err
	}
	return result, nil
}

// GetSecret - The GET operation is applicable to any secret stored in Azure Key Vault. This operation requires the secrets/get
// permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// version - The version of the secret. This URI fragment is optional. If not specified, the latest version of the secret
// is returned.
// options - GetSecretOptions contains the optional parameters for the Client.GetSecret method.
func (client *Client) GetSecret(ctx context.Context, name string, version string, options *GetSecretOptions) (GetSecretResponse, error) {
	req, err := client.getSecretCreateRequest(ctx, name, version, options)
	if err != nil {
		return GetSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSecretHandleResponse(resp)
}

// getSecretCreateRequest creates the GetSecret request.
func (client *Client) getSecretCreateRequest(ctx context.Context, name string, version string, options *GetSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}/{secret-version}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	urlPath = strings.ReplaceAll(urlPath, "{secret-version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSecretHandleResponse handles the GetSecret response.
func (client *Client) getSecretHandleResponse(resp *http.Response) (GetSecretResponse, error) {
	result := GetSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretBundle); err != nil {
		return GetSecretResponse{}, err
	}
	return result, nil
}

// NewListDeletedSecretsPager - The Get Deleted Secrets operation returns the secrets that have been deleted for a vault enabled
// for soft-delete. This operation requires the secrets/list permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// options - ListDeletedSecretsOptions contains the optional parameters for the Client.ListDeletedSecrets method.
func (client *Client) NewListDeletedSecretsPager(options *ListDeletedSecretsOptions) *runtime.Pager[ListDeletedSecretsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListDeletedSecretsResponse]{
		More: func(page ListDeletedSecretsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListDeletedSecretsResponse) (ListDeletedSecretsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listDeletedSecretsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ListDeletedSecretsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ListDeletedSecretsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ListDeletedSecretsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listDeletedSecretsHandleResponse(resp)
		},
	})
}

// listDeletedSecretsCreateRequest creates the ListDeletedSecrets request.
func (client *Client) listDeletedSecretsCreateRequest(ctx context.Context, options *ListDeletedSecretsOptions) (*policy.Request, error) {
	urlPath := "/deletedsecrets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDeletedSecretsHandleResponse handles the ListDeletedSecrets response.
func (client *Client) listDeletedSecretsHandleResponse(resp *http.Response) (ListDeletedSecretsResponse, error) {
	result := ListDeletedSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSecretListResult); err != nil {
		return ListDeletedSecretsResponse{}, err
	}
	return result, nil
}

// NewListSecretVersionsPager - The full secret identifier and attributes are provided in the response. No values are returned
// for the secrets. This operations requires the secrets/list permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// options - ListSecretVersionsOptions contains the optional parameters for the Client.ListSecretVersions method.
func (client *Client) NewListSecretVersionsPager(name string, options *ListSecretVersionsOptions) *runtime.Pager[ListSecretVersionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListSecretVersionsResponse]{
		More: func(page ListSecretVersionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListSecretVersionsResponse) (ListSecretVersionsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSecretVersionsCreateRequest(ctx, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ListSecretVersionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ListSecretVersionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ListSecretVersionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSecretVersionsHandleResponse(resp)
		},
	})
}

// listSecretVersionsCreateRequest creates the ListSecretVersions request.
func (client *Client) listSecretVersionsCreateRequest(ctx context.Context, name string, options *ListSecretVersionsOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}/versions"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretVersionsHandleResponse handles the ListSecretVersions response.
func (client *Client) listSecretVersionsHandleResponse(resp *http.Response) (ListSecretVersionsResponse, error) {
	result := ListSecretVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretListResult); err != nil {
		return ListSecretVersionsResponse{}, err
	}
	return result, nil
}

// NewListSecretsPager - The Get Secrets operation is applicable to the entire vault. However, only the base secret identifier
// and its attributes are provided in the response. Individual secret versions are not listed in the
// response. This operation requires the secrets/list permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// options - ListSecretsOptions contains the optional parameters for the Client.ListSecrets method.
func (client *Client) NewListSecretsPager(options *ListSecretsOptions) *runtime.Pager[ListSecretsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListSecretsResponse]{
		More: func(page ListSecretsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListSecretsResponse) (ListSecretsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSecretsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ListSecretsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ListSecretsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ListSecretsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSecretsHandleResponse(resp)
		},
	})
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *Client) listSecretsCreateRequest(ctx context.Context, options *ListSecretsOptions) (*policy.Request, error) {
	urlPath := "/secrets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxResults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.MaxResults), 10))
	}
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *Client) listSecretsHandleResponse(resp *http.Response) (ListSecretsResponse, error) {
	result := ListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretListResult); err != nil {
		return ListSecretsResponse{}, err
	}
	return result, nil
}

// PurgeDeletedSecret - The purge deleted secret operation removes the secret permanently, without the possibility of recovery.
// This operation can only be enabled on a soft-delete enabled vault. This operation requires the
// secrets/purge permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// options - PurgeDeletedSecretOptions contains the optional parameters for the Client.PurgeDeletedSecret method.
func (client *Client) PurgeDeletedSecret(ctx context.Context, name string, options *PurgeDeletedSecretOptions) (PurgeDeletedSecretResponse, error) {
	req, err := client.purgeDeletedSecretCreateRequest(ctx, name, options)
	if err != nil {
		return PurgeDeletedSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PurgeDeletedSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return PurgeDeletedSecretResponse{}, runtime.NewResponseError(resp)
	}
	return PurgeDeletedSecretResponse{}, nil
}

// purgeDeletedSecretCreateRequest creates the PurgeDeletedSecret request.
func (client *Client) purgeDeletedSecretCreateRequest(ctx context.Context, name string, options *PurgeDeletedSecretOptions) (*policy.Request, error) {
	urlPath := "/deletedsecrets/{secret-name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// RecoverDeletedSecret - Recovers the deleted secret in the specified vault. This operation can only be performed on a soft-delete
// enabled vault. This operation requires the secrets/recover permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the deleted secret.
// options - RecoverDeletedSecretOptions contains the optional parameters for the Client.RecoverDeletedSecret method.
func (client *Client) RecoverDeletedSecret(ctx context.Context, name string, options *RecoverDeletedSecretOptions) (RecoverDeletedSecretResponse, error) {
	req, err := client.recoverDeletedSecretCreateRequest(ctx, name, options)
	if err != nil {
		return RecoverDeletedSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecoverDeletedSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoverDeletedSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.recoverDeletedSecretHandleResponse(resp)
}

// recoverDeletedSecretCreateRequest creates the RecoverDeletedSecret request.
func (client *Client) recoverDeletedSecretCreateRequest(ctx context.Context, name string, options *RecoverDeletedSecretOptions) (*policy.Request, error) {
	urlPath := "/deletedsecrets/{secret-name}/recover"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// recoverDeletedSecretHandleResponse handles the RecoverDeletedSecret response.
func (client *Client) recoverDeletedSecretHandleResponse(resp *http.Response) (RecoverDeletedSecretResponse, error) {
	result := RecoverDeletedSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretBundle); err != nil {
		return RecoverDeletedSecretResponse{}, err
	}
	return result, nil
}

// RestoreSecret - Restores a backed up secret, and all its versions, to a vault. This operation requires the secrets/restore
// permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// parameters - The parameters to restore the secret.
// options - RestoreSecretOptions contains the optional parameters for the Client.RestoreSecret method.
func (client *Client) RestoreSecret(ctx context.Context, parameters RestoreSecretParameters, options *RestoreSecretOptions) (RestoreSecretResponse, error) {
	req, err := client.restoreSecretCreateRequest(ctx, parameters, options)
	if err != nil {
		return RestoreSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestoreSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestoreSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.restoreSecretHandleResponse(resp)
}

// restoreSecretCreateRequest creates the RestoreSecret request.
func (client *Client) restoreSecretCreateRequest(ctx context.Context, parameters RestoreSecretParameters, options *RestoreSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/restore"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// restoreSecretHandleResponse handles the RestoreSecret response.
func (client *Client) restoreSecretHandleResponse(resp *http.Response) (RestoreSecretResponse, error) {
	result := RestoreSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretBundle); err != nil {
		return RestoreSecretResponse{}, err
	}
	return result, nil
}

// SetSecret - The SET operation adds a secret to the Azure Key Vault. If the named secret already exists, Azure Key Vault
// creates a new version of that secret. This operation requires the secrets/set permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// parameters - The parameters for setting the secret.
// options - SetSecretOptions contains the optional parameters for the Client.SetSecret method.
func (client *Client) SetSecret(ctx context.Context, name string, parameters SetSecretParameters, options *SetSecretOptions) (SetSecretResponse, error) {
	req, err := client.setSecretCreateRequest(ctx, name, parameters, options)
	if err != nil {
		return SetSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SetSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SetSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.setSecretHandleResponse(resp)
}

// setSecretCreateRequest creates the SetSecret request.
func (client *Client) setSecretCreateRequest(ctx context.Context, name string, parameters SetSecretParameters, options *SetSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setSecretHandleResponse handles the SetSecret response.
func (client *Client) setSecretHandleResponse(resp *http.Response) (SetSecretResponse, error) {
	result := SetSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretBundle); err != nil {
		return SetSecretResponse{}, err
	}
	return result, nil
}

// UpdateSecret - The UPDATE operation changes specified attributes of an existing stored secret. Attributes that are not
// specified in the request are left unchanged. The value of a secret itself cannot be changed.
// This operation requires the secrets/set permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.3
// name - The name of the secret.
// version - The version of the secret.
// parameters - The parameters for update secret operation.
// options - UpdateSecretOptions contains the optional parameters for the Client.UpdateSecret method.
func (client *Client) UpdateSecret(ctx context.Context, name string, version string, parameters UpdateSecretParameters, options *UpdateSecretOptions) (UpdateSecretResponse, error) {
	req, err := client.updateSecretCreateRequest(ctx, name, version, parameters, options)
	if err != nil {
		return UpdateSecretResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdateSecretResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdateSecretResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateSecretHandleResponse(resp)
}

// updateSecretCreateRequest creates the UpdateSecret request.
func (client *Client) updateSecretCreateRequest(ctx context.Context, name string, version string, parameters UpdateSecretParameters, options *UpdateSecretOptions) (*policy.Request, error) {
	urlPath := "/secrets/{secret-name}/{secret-version}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secret-name}", url.PathEscape(name))
	urlPath = strings.ReplaceAll(urlPath, "{secret-version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.3")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateSecretHandleResponse handles the UpdateSecret response.
func (client *Client) updateSecretHandleResponse(resp *http.Response) (UpdateSecretResponse, error) {
	result := UpdateSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretBundle); err != nil {
		return UpdateSecretResponse{}, err
	}
	return result, nil
}
