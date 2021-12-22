//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Get.json
func ExampleIntegrationRuntimeNodesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewIntegrationRuntimeNodesClient("<subscription-id>", cred, nil)
	_, err = client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<integration-runtime-name>",
		"<node-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Update.json
func ExampleIntegrationRuntimeNodesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewIntegrationRuntimeNodesClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<integration-runtime-name>",
		"<node-name>",
		armsynapse.UpdateIntegrationRuntimeNodeRequest{
			ConcurrentJobsLimit: to.Int32Ptr(2),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Delete.json
func ExampleIntegrationRuntimeNodesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewIntegrationRuntimeNodesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<integration-runtime-name>",
		"<node-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
