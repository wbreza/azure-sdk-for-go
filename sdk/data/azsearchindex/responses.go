// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsearchindex

// DocumentsClientAutocompleteGetResponse contains the response from method DocumentsClient.AutocompleteGet.
type DocumentsClientAutocompleteGetResponse struct {
	// The result of Autocomplete query.
	AutocompleteResult
}

// DocumentsClientAutocompletePostResponse contains the response from method DocumentsClient.AutocompletePost.
type DocumentsClientAutocompletePostResponse struct {
	// The result of Autocomplete query.
	AutocompleteResult
}

// DocumentsClientCountResponse contains the response from method DocumentsClient.Count.
type DocumentsClientCountResponse struct {
	Value *int64
}

// DocumentsClientGetResponse contains the response from method DocumentsClient.Get.
type DocumentsClientGetResponse struct {
	// A document retrieved via a document lookup operation.
	Value map[string]any
}

// DocumentsClientIndexResponse contains the response from method DocumentsClient.Index.
type DocumentsClientIndexResponse struct {
	// Response containing the status of operations for all documents in the indexing request.
	IndexDocumentsResult
}

// DocumentsClientSearchGetResponse contains the response from method DocumentsClient.SearchGet.
type DocumentsClientSearchGetResponse struct {
	// Response containing search results from an index.
	SearchDocumentsResult
}

// DocumentsClientSearchPostResponse contains the response from method DocumentsClient.SearchPost.
type DocumentsClientSearchPostResponse struct {
	// Response containing search results from an index.
	SearchDocumentsResult
}

// DocumentsClientSuggestGetResponse contains the response from method DocumentsClient.SuggestGet.
type DocumentsClientSuggestGetResponse struct {
	// Response containing suggestion query results from an index.
	SuggestDocumentsResult
}

// DocumentsClientSuggestPostResponse contains the response from method DocumentsClient.SuggestPost.
type DocumentsClientSuggestPostResponse struct {
	// Response containing suggestion query results from an index.
	SuggestDocumentsResult
}

