package azsearch

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewSkillsetsClient(endpoint string, credential azcore.TokenCredential, options *azcore.ClientOptions) (*SkillsetsClient, error) {
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

	return &SkillsetsClient{
		endpoint: endpoint,
		internal: azcoreClient,
	}, nil
}
