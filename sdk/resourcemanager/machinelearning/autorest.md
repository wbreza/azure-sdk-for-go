### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
go: true
azure-arm: true
output-folder: armmachinelearning
tag: package-preview-2024-01
module: github.com/wbreza/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning
require:
- https://github.com/Azure/azure-rest-api-specs/blob/2d973fccf9f28681a481e9760fa12b2334216e21/specification/machinelearningservices/resource-manager/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/2d973fccf9f28681a481e9760fa12b2334216e21/specification/machinelearningservices/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module-version: 3.2.0
```