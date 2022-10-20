//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// RoleEligibilityScheduleRequestsClient contains the methods for the RoleEligibilityScheduleRequests group.
// Don't use this type directly, use NewRoleEligibilityScheduleRequestsClient() instead.
type RoleEligibilityScheduleRequestsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRoleEligibilityScheduleRequestsClient creates a new instance of RoleEligibilityScheduleRequestsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleEligibilityScheduleRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleEligibilityScheduleRequestsClient, error) {
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
	client := &RoleEligibilityScheduleRequestsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Cancel - Cancels a pending role eligibility schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility request to cancel.
// roleEligibilityScheduleRequestName - The name of the role eligibility request to cancel.
// options - RoleEligibilityScheduleRequestsClientCancelOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Cancel
// method.
func (client *RoleEligibilityScheduleRequestsClient) Cancel(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsClientCancelOptions) (RoleEligibilityScheduleRequestsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleRequestsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return RoleEligibilityScheduleRequestsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RoleEligibilityScheduleRequestsClient) cancelCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Create - Creates a role eligibility schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility schedule request to create. The scope can be any REST resource instance. For
// example, use '/subscriptions/{subscription-id}/' for a subscription,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource.
// roleEligibilityScheduleRequestName - The name of the role eligibility to create. It can be any valid GUID.
// parameters - Parameters for the role eligibility schedule request.
// options - RoleEligibilityScheduleRequestsClientCreateOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Create
// method.
func (client *RoleEligibilityScheduleRequestsClient) Create(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsClientCreateOptions) (RoleEligibilityScheduleRequestsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, parameters, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleEligibilityScheduleRequestsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleEligibilityScheduleRequestsClient) createCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleEligibilityScheduleRequestsClient) createHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsClientCreateResponse, error) {
	result := RoleEligibilityScheduleRequestsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequest); err != nil {
		return RoleEligibilityScheduleRequestsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get the specified role eligibility schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility schedule request.
// roleEligibilityScheduleRequestName - The name (guid) of the role eligibility schedule request to get.
// options - RoleEligibilityScheduleRequestsClientGetOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Get
// method.
func (client *RoleEligibilityScheduleRequestsClient) Get(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsClientGetOptions) (RoleEligibilityScheduleRequestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleRequestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilityScheduleRequestsClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilityScheduleRequestsClient) getHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsClientGetResponse, error) {
	result := RoleEligibilityScheduleRequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequest); err != nil {
		return RoleEligibilityScheduleRequestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role eligibility schedule requests for a scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility schedule requests.
// options - RoleEligibilityScheduleRequestsClientListForScopeOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.ListForScope
// method.
func (client *RoleEligibilityScheduleRequestsClient) NewListForScopePager(scope string, options *RoleEligibilityScheduleRequestsClientListForScopeOptions) *runtime.Pager[RoleEligibilityScheduleRequestsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleEligibilityScheduleRequestsClientListForScopeResponse]{
		More: func(page RoleEligibilityScheduleRequestsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleEligibilityScheduleRequestsClientListForScopeResponse) (RoleEligibilityScheduleRequestsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleEligibilityScheduleRequestsClientListForScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleEligibilityScheduleRequestsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleEligibilityScheduleRequestsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilityScheduleRequestsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilityScheduleRequestsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilityScheduleRequestsClient) listForScopeHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsClientListForScopeResponse, error) {
	result := RoleEligibilityScheduleRequestsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequestListResult); err != nil {
		return RoleEligibilityScheduleRequestsClientListForScopeResponse{}, err
	}
	return result, nil
}

// Validate - Validates a new role eligibility schedule request.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility request to validate.
// roleEligibilityScheduleRequestName - The name of the role eligibility request to validate.
// parameters - Parameters for the role eligibility schedule request.
// options - RoleEligibilityScheduleRequestsClientValidateOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Validate
// method.
func (client *RoleEligibilityScheduleRequestsClient) Validate(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsClientValidateOptions) (RoleEligibilityScheduleRequestsClientValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, parameters, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleRequestsClientValidateResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *RoleEligibilityScheduleRequestsClient) validateCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}/validate"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateHandleResponse handles the Validate response.
func (client *RoleEligibilityScheduleRequestsClient) validateHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsClientValidateResponse, error) {
	result := RoleEligibilityScheduleRequestsClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequest); err != nil {
		return RoleEligibilityScheduleRequestsClientValidateResponse{}, err
	}
	return result, nil
}