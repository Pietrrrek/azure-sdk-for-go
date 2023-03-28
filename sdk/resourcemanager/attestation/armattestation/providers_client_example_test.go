//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armattestation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvider.json
func ExampleProvidersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().Get(ctx, "MyResourceGroup", "myattestationprovider", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armattestation.Provider{
	// 	Name: to.Ptr("myattestationprovider"),
	// 	Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Property1": to.Ptr("Value1"),
	// 		"Property2": to.Ptr("Value2"),
	// 		"Property3": to.Ptr("Value3"),
	// 	},
	// 	Properties: &armattestation.StatusResult{
	// 		AttestURI: to.Ptr("https://superservice.attestation.azure.net"),
	// 		Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 		TrustModel: to.Ptr("Isolated"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Create_AttestationProvider.json
func ExampleProvidersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().Create(ctx, "MyResourceGroup", "myattestationprovider", armattestation.ServiceCreationParams{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armattestation.Provider{
	// 	Name: to.Ptr("myattestationprovider"),
	// 	Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Property1": to.Ptr("Value1"),
	// 		"Property2": to.Ptr("Value2"),
	// 		"Property3": to.Ptr("Value3"),
	// 	},
	// 	Properties: &armattestation.StatusResult{
	// 		AttestURI: to.Ptr("https://superservice.attestation.azure.net"),
	// 		Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 		TrustModel: to.Ptr("Isolated"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Update_AttestationProvider.json
func ExampleProvidersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().Update(ctx, "MyResourceGroup", "myattestationprovider", armattestation.ServicePatchParams{
		Tags: map[string]*string{
			"Property1": to.Ptr("Value1"),
			"Property2": to.Ptr("Value2"),
			"Property3": to.Ptr("Value3"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armattestation.Provider{
	// 	Name: to.Ptr("myattestationprovider"),
	// 	Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"Property1": to.Ptr("Value1"),
	// 		"Property2": to.Ptr("Value2"),
	// 		"Property3": to.Ptr("Value3"),
	// 	},
	// 	Properties: &armattestation.StatusResult{
	// 		AttestURI: to.Ptr("https://superservice.attestation.azure.net"),
	// 		Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 		TrustModel: to.Ptr("Isolated"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Delete_AttestationProvider.json
func ExampleProvidersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProvidersClient().Delete(ctx, "sample-resource-group", "myattestationprovider", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvidersList.json
func ExampleProvidersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderListResult = armattestation.ProviderListResult{
	// 	Value: []*armattestation.Provider{
	// 		{
	// 			Name: to.Ptr("myattestationprovider"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("subscriptions/6c96b33e-f5b8-40a6-9011-5cb1c58b0915/resourceGroups/testrg1/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armattestation.StatusResult{
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("codes2"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("subscriptions/6c96b33e-f5b8-40a6-9011-5cb1c58b0915/resourceGroups/testrg2/providers/Microsoft.Attestation/attestationProviders/codes2"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armattestation.StatusResult{
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_AttestationProvidersListByResourceGroup.json
func ExampleProvidersClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().ListByResourceGroup(ctx, "testrg1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderListResult = armattestation.ProviderListResult{
	// 	Value: []*armattestation.Provider{
	// 		{
	// 			Name: to.Ptr("myattestationprovider"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("subscriptions/6c96b33e-f5b8-40a6-9011-5cb1c58b0915/resourceGroups/testrg1/providers/Microsoft.Attestation/attestationProviders/myattestationprovider"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armattestation.StatusResult{
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("codes2"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("subscriptions/6c96b33e-f5b8-40a6-9011-5cb1c58b0915/resourceGroups/testrg1/providers/Microsoft.Attestation/attestationProviders/codes2"),
	// 			Location: to.Ptr("East US"),
	// 			Properties: &armattestation.StatusResult{
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_DefaultProviders.json
func ExampleProvidersClient_ListDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().ListDefault(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderListResult = armattestation.ProviderListResult{
	// 	Value: []*armattestation.Provider{
	// 		{
	// 			Name: to.Ptr("sharedcus"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/sharedcus"),
	// 			Location: to.Ptr("Central US"),
	// 			Properties: &armattestation.StatusResult{
	// 				AttestURI: to.Ptr("https://sharedcus.cus.attest.azure.net"),
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 				TrustModel: to.Ptr("AAD"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("shareduks"),
	// 			Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 			ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/shareduks"),
	// 			Location: to.Ptr("UK South"),
	// 			Properties: &armattestation.StatusResult{
	// 				AttestURI: to.Ptr("https://shareduks.uks.attest.azure.net"),
	// 				Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 				TrustModel: to.Ptr("AAD"),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/attestation/resource-manager/Microsoft.Attestation/stable/2020-10-01/examples/Get_DefaultProviderByLocation.json
func ExampleProvidersClient_GetDefaultByLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armattestation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().GetDefaultByLocation(ctx, "Central US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armattestation.Provider{
	// 	Name: to.Ptr("sharedcus"),
	// 	Type: to.Ptr("Microsoft.Attestation/attestationProviders"),
	// 	ID: to.Ptr("providers/Microsoft.Attestation/attestationProviders/sharedcus"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armattestation.StatusResult{
	// 		AttestURI: to.Ptr("https://sharedcus.cus.attest.azure.net"),
	// 		Status: to.Ptr(armattestation.AttestationServiceStatusReady),
	// 		TrustModel: to.Ptr("AAD"),
	// 	},
	// }
}