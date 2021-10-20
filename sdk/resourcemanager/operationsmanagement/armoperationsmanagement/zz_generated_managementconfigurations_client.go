//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationsmanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ManagementConfigurationsClient contains the methods for the ManagementConfigurations group.
// Don't use this type directly, use NewManagementConfigurationsClient() instead.
type ManagementConfigurationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagementConfigurationsClient creates a new instance of ManagementConfigurationsClient with the specified values.
func NewManagementConfigurationsClient(con *arm.Connection, subscriptionID string) *ManagementConfigurationsClient {
	return &ManagementConfigurationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates the ManagementConfiguration.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managementConfigurationName string, parameters ManagementConfiguration, options *ManagementConfigurationsCreateOrUpdateOptions) (ManagementConfigurationsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managementConfigurationName, parameters, options)
	if err != nil {
		return ManagementConfigurationsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementConfigurationsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementConfigurationsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagementConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managementConfigurationName string, parameters ManagementConfiguration, options *ManagementConfigurationsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managementConfigurationName == "" {
		return nil, errors.New("parameter managementConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementConfigurationName}", url.PathEscape(managementConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagementConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagementConfigurationsCreateOrUpdateResponse, error) {
	result := ManagementConfigurationsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementConfiguration); err != nil {
		return ManagementConfigurationsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagementConfigurationsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the ManagementConfiguration in the subscription.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, managementConfigurationName string, options *ManagementConfigurationsDeleteOptions) (ManagementConfigurationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managementConfigurationName, options)
	if err != nil {
		return ManagementConfigurationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementConfigurationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementConfigurationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ManagementConfigurationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managementConfigurationName string, options *ManagementConfigurationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managementConfigurationName == "" {
		return nil, errors.New("parameter managementConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementConfigurationName}", url.PathEscape(managementConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ManagementConfigurationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the user ManagementConfiguration.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementConfigurationsClient) Get(ctx context.Context, resourceGroupName string, managementConfigurationName string, options *ManagementConfigurationsGetOptions) (ManagementConfigurationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managementConfigurationName, options)
	if err != nil {
		return ManagementConfigurationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementConfigurationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementConfigurationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagementConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managementConfigurationName string, options *ManagementConfigurationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managementConfigurationName == "" {
		return nil, errors.New("parameter managementConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementConfigurationName}", url.PathEscape(managementConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementConfigurationsClient) getHandleResponse(resp *http.Response) (ManagementConfigurationsGetResponse, error) {
	result := ManagementConfigurationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementConfiguration); err != nil {
		return ManagementConfigurationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagementConfigurationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Retrieves the ManagementConfigurations list.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementConfigurationsClient) ListBySubscription(ctx context.Context, options *ManagementConfigurationsListBySubscriptionOptions) (ManagementConfigurationsListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ManagementConfigurationsListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementConfigurationsListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementConfigurationsListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagementConfigurationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagementConfigurationsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagementConfigurationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagementConfigurationsListBySubscriptionResponse, error) {
	result := ManagementConfigurationsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementConfigurationPropertiesList); err != nil {
		return ManagementConfigurationsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ManagementConfigurationsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
