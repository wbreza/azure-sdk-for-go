//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_List.json
func ExampleOpenShiftClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_ListByResourceGroup.json
func ExampleOpenShiftClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("resourceGroup",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_Get.json
func ExampleOpenShiftClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"resourceGroup",
		"resourceName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_CreateOrUpdate.json
func ExampleOpenShiftClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroup",
		"resourceName",
		armredhatopenshift.OpenShiftCluster{
			Location: to.Ptr("location"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armredhatopenshift.OpenShiftClusterProperties{
				ApiserverProfile: &armredhatopenshift.APIServerProfile{
					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
				},
				ClusterProfile: &armredhatopenshift.ClusterProfile{
					Domain:               to.Ptr("cluster.location.aroapp.io"),
					FipsValidatedModules: to.Ptr(armredhatopenshift.FipsValidatedModulesEnabled),
					PullSecret:           to.Ptr("{\"auths\":{\"registry.connect.redhat.com\":{\"auth\":\"\"},\"registry.redhat.io\":{\"auth\":\"\"}}}"),
					ResourceGroupID:      to.Ptr("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
				},
				ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
				IngressProfiles: []*armredhatopenshift.IngressProfile{
					{
						Name:       to.Ptr("default"),
						Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
					}},
				MasterProfile: &armredhatopenshift.MasterProfile{
					EncryptionAtHost: to.Ptr(armredhatopenshift.EncryptionAtHostEnabled),
					SubnetID:         to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
					VMSize:           to.Ptr("Standard_D8s_v3"),
				},
				NetworkProfile: &armredhatopenshift.NetworkProfile{
					PodCidr:     to.Ptr("10.128.0.0/14"),
					ServiceCidr: to.Ptr("172.30.0.0/16"),
				},
				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
					ClientID:     to.Ptr("clientId"),
					ClientSecret: to.Ptr("clientSecret"),
				},
				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
					{
						Name:       to.Ptr("worker"),
						Count:      to.Ptr[int32](3),
						DiskSizeGB: to.Ptr[int32](128),
						SubnetID:   to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
						VMSize:     to.Ptr("Standard_D2s_v3"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_Delete.json
func ExampleOpenShiftClustersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"resourceGroup",
		"resourceName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_Update.json
func ExampleOpenShiftClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"resourceGroup",
		"resourceName",
		armredhatopenshift.OpenShiftClusterUpdate{
			Properties: &armredhatopenshift.OpenShiftClusterProperties{
				ApiserverProfile: &armredhatopenshift.APIServerProfile{
					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
				},
				ClusterProfile: &armredhatopenshift.ClusterProfile{
					Domain:               to.Ptr("cluster.location.aroapp.io"),
					FipsValidatedModules: to.Ptr(armredhatopenshift.FipsValidatedModulesEnabled),
					PullSecret:           to.Ptr("{\"auths\":{\"registry.connect.redhat.com\":{\"auth\":\"\"},\"registry.redhat.io\":{\"auth\":\"\"}}}"),
					ResourceGroupID:      to.Ptr("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
				},
				ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
				IngressProfiles: []*armredhatopenshift.IngressProfile{
					{
						Name:       to.Ptr("default"),
						Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
					}},
				MasterProfile: &armredhatopenshift.MasterProfile{
					EncryptionAtHost: to.Ptr(armredhatopenshift.EncryptionAtHostEnabled),
					SubnetID:         to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
					VMSize:           to.Ptr("Standard_D8s_v3"),
				},
				NetworkProfile: &armredhatopenshift.NetworkProfile{
					PodCidr:     to.Ptr("10.128.0.0/14"),
					ServiceCidr: to.Ptr("172.30.0.0/16"),
				},
				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
					ClientID:     to.Ptr("clientId"),
					ClientSecret: to.Ptr("clientSecret"),
				},
				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
					{
						Name:       to.Ptr("worker"),
						Count:      to.Ptr[int32](3),
						DiskSizeGB: to.Ptr[int32](128),
						SubnetID:   to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
						VMSize:     to.Ptr("Standard_D2s_v3"),
					}},
			},
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_ListAdminCredentials.json
func ExampleOpenShiftClustersClient_ListAdminCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListAdminCredentials(ctx,
		"resourceGroup",
		"resourceName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-04-01/examples/OpenShiftClusters_ListCredentials.json
func ExampleOpenShiftClustersClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewOpenShiftClustersClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListCredentials(ctx,
		"resourceGroup",
		"resourceName",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}