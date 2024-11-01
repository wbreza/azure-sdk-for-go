package azsearch

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewDataSourcesClient(endpoint string, credential azcore.TokenCredential, options *azcore.ClientOptions) (*DataSourcesClient, error) {
	if options == nil {
		options = &azcore.ClientOptions{}
	}

	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{tokenScope}, &policy.BearerTokenOptions{
		InsecureAllowCredentialWithHTTP: false,
	})

	azcoreClient, err := azcore.NewClient(clientName, moduleVersion, runtime.PipelineOptions{
		PerRetry: []policy.Policy{authPolicy},
	}, options)

	if err != nil {
		return nil, err
	}

	return &DataSourcesClient{
		endpoint: endpoint,
		internal: azcoreClient,
	}, nil
}
