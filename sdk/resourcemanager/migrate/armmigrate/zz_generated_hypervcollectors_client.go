//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HyperVCollectorsClient contains the methods for the HyperVCollectors group.
// Don't use this type directly, use NewHyperVCollectorsClient() instead.
type HyperVCollectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHyperVCollectorsClient creates a new instance of HyperVCollectorsClient with the specified values.
// subscriptionID - Azure Subscription Id in which project was created.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHyperVCollectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HyperVCollectorsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HyperVCollectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create or Update Hyper-V collector
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// hyperVCollectorName - Unique name of a Hyper-V collector within a project.
// options - HyperVCollectorsClientCreateOptions contains the optional parameters for the HyperVCollectorsClient.Create method.
func (client *HyperVCollectorsClient) Create(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientCreateOptions) (HyperVCollectorsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, hyperVCollectorName, options)
	if err != nil {
		return HyperVCollectorsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVCollectorsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HyperVCollectorsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *HyperVCollectorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if hyperVCollectorName == "" {
		return nil, errors.New("parameter hyperVCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hyperVCollectorName}", url.PathEscape(hyperVCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CollectorBody != nil {
		return req, runtime.MarshalAsJSON(req, *options.CollectorBody)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *HyperVCollectorsClient) createHandleResponse(resp *http.Response) (HyperVCollectorsClientCreateResponse, error) {
	result := HyperVCollectorsClientCreateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVCollector); err != nil {
		return HyperVCollectorsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Hyper-V collector from the project.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// hyperVCollectorName - Unique name of a Hyper-V collector within a project.
// options - HyperVCollectorsClientDeleteOptions contains the optional parameters for the HyperVCollectorsClient.Delete method.
func (client *HyperVCollectorsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientDeleteOptions) (HyperVCollectorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, hyperVCollectorName, options)
	if err != nil {
		return HyperVCollectorsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVCollectorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HyperVCollectorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *HyperVCollectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if hyperVCollectorName == "" {
		return nil, errors.New("parameter hyperVCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hyperVCollectorName}", url.PathEscape(hyperVCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *HyperVCollectorsClient) deleteHandleResponse(resp *http.Response) (HyperVCollectorsClientDeleteResponse, error) {
	result := HyperVCollectorsClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// Get - Get a Hyper-V collector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// hyperVCollectorName - Unique name of a Hyper-V collector within a project.
// options - HyperVCollectorsClientGetOptions contains the optional parameters for the HyperVCollectorsClient.Get method.
func (client *HyperVCollectorsClient) Get(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientGetOptions) (HyperVCollectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, hyperVCollectorName, options)
	if err != nil {
		return HyperVCollectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVCollectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HyperVCollectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HyperVCollectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, hyperVCollectorName string, options *HyperVCollectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if hyperVCollectorName == "" {
		return nil, errors.New("parameter hyperVCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hyperVCollectorName}", url.PathEscape(hyperVCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HyperVCollectorsClient) getHandleResponse(resp *http.Response) (HyperVCollectorsClientGetResponse, error) {
	result := HyperVCollectorsClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVCollector); err != nil {
		return HyperVCollectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - Get a list of Hyper-V collector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// resourceGroupName - Name of the Azure Resource Group that project is part of.
// projectName - Name of the Azure Migrate project.
// options - HyperVCollectorsClientListByProjectOptions contains the optional parameters for the HyperVCollectorsClient.ListByProject
// method.
func (client *HyperVCollectorsClient) NewListByProjectPager(resourceGroupName string, projectName string, options *HyperVCollectorsClientListByProjectOptions) *runtime.Pager[HyperVCollectorsClientListByProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[HyperVCollectorsClientListByProjectResponse]{
		More: func(page HyperVCollectorsClientListByProjectResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *HyperVCollectorsClientListByProjectResponse) (HyperVCollectorsClientListByProjectResponse, error) {
			req, err := client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			if err != nil {
				return HyperVCollectorsClientListByProjectResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HyperVCollectorsClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HyperVCollectorsClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *HyperVCollectorsClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *HyperVCollectorsClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *HyperVCollectorsClient) listByProjectHandleResponse(resp *http.Response) (HyperVCollectorsClientListByProjectResponse, error) {
	result := HyperVCollectorsClientListByProjectResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVCollectorList); err != nil {
		return HyperVCollectorsClientListByProjectResponse{}, err
	}
	return result, nil
}