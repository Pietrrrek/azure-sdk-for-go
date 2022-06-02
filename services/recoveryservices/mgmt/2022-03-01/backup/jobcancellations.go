package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// JobCancellationsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type JobCancellationsClient struct {
	BaseClient
}

// NewJobCancellationsClient creates an instance of the JobCancellationsClient client.
func NewJobCancellationsClient(subscriptionID string) JobCancellationsClient {
	return NewJobCancellationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobCancellationsClientWithBaseURI creates an instance of the JobCancellationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewJobCancellationsClientWithBaseURI(baseURI string, subscriptionID string) JobCancellationsClient {
	return JobCancellationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger cancels a job. This is an asynchronous operation. To know the status of the cancellation, call
// GetCancelOperationResult API.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// jobName - name of the job to cancel.
func (client JobCancellationsClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCancellationsClient.Trigger")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerPreparer(ctx, vaultName, resourceGroupName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.JobCancellationsClient", "Trigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.TriggerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "backup.JobCancellationsClient", "Trigger", resp, "Failure sending request")
		return
	}

	result, err = client.TriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.JobCancellationsClient", "Trigger", resp, "Failure responding to request")
		return
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client JobCancellationsClient) TriggerPreparer(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2022-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client JobCancellationsClient) TriggerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client JobCancellationsClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}