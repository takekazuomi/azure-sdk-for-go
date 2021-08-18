// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// CosmosContainerResponse represents the response from a container request.
type CosmosContainerResponse struct {
	// ContainerProperties contains the unmarshalled response body in CosmosContainerProperties format.
	ContainerProperties *CosmosContainerProperties
	cosmosResponse
}

func newCosmosContainerResponse(resp *azcore.Response, container *CosmosContainer) (CosmosContainerResponse, error) {
	response := CosmosContainerResponse{}
	response.RawResponse = resp.Response
	properties := &CosmosContainerProperties{}
	err := resp.UnmarshalAsJSON(properties)
	if err != nil {
		return response, err
	}
	response.ContainerProperties = properties
	response.ContainerProperties.Container = container
	return response, nil
}
