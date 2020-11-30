// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ListBlobsFlatSegmentResponsePager provides iteration over ListBlobsFlatSegmentResponse pages.
type ListBlobsFlatSegmentResponsePager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ListBlobsFlatSegmentResponseResponse.
	PageResponse() *ListBlobsFlatSegmentResponseResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type listBlobsFlatSegmentResponseCreateRequest func(context.Context) (*azcore.Request, error)

type listBlobsFlatSegmentResponseHandleError func(*azcore.Response) error

type listBlobsFlatSegmentResponseHandleResponse func(*azcore.Response) (*ListBlobsFlatSegmentResponseResponse, error)

type listBlobsFlatSegmentResponseAdvancePage func(context.Context, *ListBlobsFlatSegmentResponseResponse) (*azcore.Request, error)

type listBlobsFlatSegmentResponsePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listBlobsFlatSegmentResponseCreateRequest
	// callback for handling response errors
	errorer listBlobsFlatSegmentResponseHandleError
	// callback for handling the HTTP response
	responder listBlobsFlatSegmentResponseHandleResponse
	// callback for advancing to the next page
	advancer listBlobsFlatSegmentResponseAdvancePage
	// contains the current response
	current *ListBlobsFlatSegmentResponseResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listBlobsFlatSegmentResponsePager) Err() error {
	return p.err
}

func (p *listBlobsFlatSegmentResponsePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.EnumerationResults.NextMarker == nil || len(*p.current.EnumerationResults.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
  } else {
		req, err = p.requester(ctx)
  }
	if err != nil {
		p.err = err
		return false
	}
  	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
	p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listBlobsFlatSegmentResponsePager) PageResponse() *ListBlobsFlatSegmentResponseResponse {
	return p.current
}

// ListBlobsHierarchySegmentResponsePager provides iteration over ListBlobsHierarchySegmentResponse pages.
type ListBlobsHierarchySegmentResponsePager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ListBlobsHierarchySegmentResponseResponse.
	PageResponse() *ListBlobsHierarchySegmentResponseResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type listBlobsHierarchySegmentResponseCreateRequest func(context.Context) (*azcore.Request, error)

type listBlobsHierarchySegmentResponseHandleError func(*azcore.Response) error

type listBlobsHierarchySegmentResponseHandleResponse func(*azcore.Response) (*ListBlobsHierarchySegmentResponseResponse, error)

type listBlobsHierarchySegmentResponseAdvancePage func(context.Context, *ListBlobsHierarchySegmentResponseResponse) (*azcore.Request, error)

type listBlobsHierarchySegmentResponsePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listBlobsHierarchySegmentResponseCreateRequest
	// callback for handling response errors
	errorer listBlobsHierarchySegmentResponseHandleError
	// callback for handling the HTTP response
	responder listBlobsHierarchySegmentResponseHandleResponse
	// callback for advancing to the next page
	advancer listBlobsHierarchySegmentResponseAdvancePage
	// contains the current response
	current *ListBlobsHierarchySegmentResponseResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listBlobsHierarchySegmentResponsePager) Err() error {
	return p.err
}

func (p *listBlobsHierarchySegmentResponsePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.EnumerationResults.NextMarker == nil || len(*p.current.EnumerationResults.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
  } else {
		req, err = p.requester(ctx)
  }
	if err != nil {
		p.err = err
		return false
	}
  	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
	p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listBlobsHierarchySegmentResponsePager) PageResponse() *ListBlobsHierarchySegmentResponseResponse {
	return p.current
}

// ListContainersSegmentResponsePager provides iteration over ListContainersSegmentResponse pages.
type ListContainersSegmentResponsePager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ListContainersSegmentResponseResponse.
	PageResponse() *ListContainersSegmentResponseResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type listContainersSegmentResponseCreateRequest func(context.Context) (*azcore.Request, error)

type listContainersSegmentResponseHandleError func(*azcore.Response) error

type listContainersSegmentResponseHandleResponse func(*azcore.Response) (*ListContainersSegmentResponseResponse, error)

type listContainersSegmentResponseAdvancePage func(context.Context, *ListContainersSegmentResponseResponse) (*azcore.Request, error)

type listContainersSegmentResponsePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester listContainersSegmentResponseCreateRequest
	// callback for handling response errors
	errorer listContainersSegmentResponseHandleError
	// callback for handling the HTTP response
	responder listContainersSegmentResponseHandleResponse
	// callback for advancing to the next page
	advancer listContainersSegmentResponseAdvancePage
	// contains the current response
	current *ListContainersSegmentResponseResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *listContainersSegmentResponsePager) Err() error {
	return p.err
}

func (p *listContainersSegmentResponsePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if p.current != nil {
		if p.current.EnumerationResults.NextMarker == nil || len(*p.current.EnumerationResults.NextMarker) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
  } else {
		req, err = p.requester(ctx)
  }
	if err != nil {
		p.err = err
		return false
	}
  	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
	p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *listContainersSegmentResponsePager) PageResponse() *ListContainersSegmentResponseResponse {
	return p.current
}

