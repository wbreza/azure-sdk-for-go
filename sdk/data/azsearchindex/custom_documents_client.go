package azsearchindex

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDocumentsClient(endpoint string, indexName string, credential azcore.TokenCredential, options *azcore.ClientOptions) (*DocumentsClient, error) {
	if options == nil {
		options = &azcore.ClientOptions{}
	}

	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{tokenScope}, nil)
	pipelineOptions := runtime.PipelineOptions{
		PerRetry: []policy.Policy{authPolicy},
	}

	azcoreClient, err := azcore.NewClient(clientName, moduleVersion, pipelineOptions, options)
	if err != nil {
		return nil, err
	}

	return &DocumentsClient{
		endpoint:  endpoint,
		indexName: indexName,
		internal:  azcoreClient,
	}, nil
}
