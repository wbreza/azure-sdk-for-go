## Go

``` yaml
input-file: https://raw.githubusercontent.com/Azure/azure-rest-api-specs/refs/heads/main/specification/search/data-plane/Azure.Search/stable/2024-07-01/searchindex.json
license-header: MICROSOFT_MIT_NO_VERSION
openapi-type: data-plane
go: true
clear-output-folder: false
output-folder: ../azsearchindex
namespace: azsearchindex
go-module: github.com/wbreza/azure-sdk-for-go/sdk/data/azsearchindex
data-plane: true
credential-type: TokenCredential
module: github.com/wbreza/azure-sdk-for-go/sdk/data/azsearchindex
module-version: 0.1.0
generate-metadata: true
add-credential: true
use: "@autorest/go@latest"
directive:
  - from: swagger-document
    where: $.definitions.SearchIndexerStatus
    transform: >
      $["x-ms-client-name"] = "SearchIndexerStatus2";
