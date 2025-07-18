/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type RecordMxAPI interface {
	/*
		Create Create a record:mx object

		Creates a new record:mx object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordMxAPICreateRequest
	*/
	Create(ctx context.Context) RecordMxAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateRecordMxResponse
	CreateExecute(r RecordMxAPICreateRequest) (*CreateRecordMxResponse, *http.Response, error)
	/*
		Delete Delete a record:mx object

		Deletes a specific record:mx object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:mx object
		@return RecordMxAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) RecordMxAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r RecordMxAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve record:mx objects

		Returns a list of record:mx objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordMxAPIListRequest
	*/
	List(ctx context.Context) RecordMxAPIListRequest

	// ListExecute executes the request
	//  @return ListRecordMxResponse
	ListExecute(r RecordMxAPIListRequest) (*ListRecordMxResponse, *http.Response, error)
	/*
		Read Get a specific record:mx object

		Returns a specific record:mx object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:mx object
		@return RecordMxAPIReadRequest
	*/
	Read(ctx context.Context, reference string) RecordMxAPIReadRequest

	// ReadExecute executes the request
	//  @return GetRecordMxResponse
	ReadExecute(r RecordMxAPIReadRequest) (*GetRecordMxResponse, *http.Response, error)
	/*
		Update Update a record:mx object

		Updates a specific record:mx object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the record:mx object
		@return RecordMxAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) RecordMxAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateRecordMxResponse
	UpdateExecute(r RecordMxAPIUpdateRequest) (*UpdateRecordMxResponse, *http.Response, error)
}

// RecordMxAPIService RecordMxAPI service
type RecordMxAPIService internal.Service

type RecordMxAPICreateRequest struct {
	ctx              context.Context
	ApiService       RecordMxAPI
	recordMx         *RecordMx
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r RecordMxAPICreateRequest) RecordMx(recordMx RecordMx) RecordMxAPICreateRequest {
	r.recordMx = &recordMx
	return r
}

// Enter the field names followed by comma
func (r RecordMxAPICreateRequest) ReturnFields(returnFields string) RecordMxAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordMxAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordMxAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordMxAPICreateRequest) ReturnAsObject(returnAsObject int32) RecordMxAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordMxAPICreateRequest) Execute() (*CreateRecordMxResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a record:mx object

Creates a new record:mx object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordMxAPICreateRequest
*/
func (a *RecordMxAPIService) Create(ctx context.Context) RecordMxAPICreateRequest {
	return RecordMxAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRecordMxResponse
func (a *RecordMxAPIService) CreateExecute(r RecordMxAPICreateRequest) (*CreateRecordMxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateRecordMxResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordMxAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:mx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordMx == nil {
		return localVarReturnValue, nil, internal.ReportError("recordMx is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordMx != nil {
		if r.recordMx.ExtAttrs == nil {
			r.recordMx.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordMx.ExtAttrs)[k]; !ok {
				(*r.recordMx.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.recordMx
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordMxAPIDeleteRequest struct {
	ctx        context.Context
	ApiService RecordMxAPI
	reference  string
}

func (r RecordMxAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a record:mx object

Deletes a specific record:mx object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:mx object
	@return RecordMxAPIDeleteRequest
*/
func (a *RecordMxAPIService) Delete(ctx context.Context, reference string) RecordMxAPIDeleteRequest {
	return RecordMxAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *RecordMxAPIService) DeleteExecute(r RecordMxAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordMxAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:mx/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RecordMxAPIListRequest struct {
	ctx              context.Context
	ApiService       RecordMxAPI
	returnFields     *string
	returnFieldsPlus *string
	maxResults       *int32
	returnAsObject   *int32
	paging           *int32
	pageId           *string
	filters          *map[string]interface{}
	extattrfilter    *map[string]interface{}
}

// Enter the field names followed by comma
func (r RecordMxAPIListRequest) ReturnFields(returnFields string) RecordMxAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordMxAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordMxAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r RecordMxAPIListRequest) MaxResults(maxResults int32) RecordMxAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r RecordMxAPIListRequest) ReturnAsObject(returnAsObject int32) RecordMxAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r RecordMxAPIListRequest) Paging(paging int32) RecordMxAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r RecordMxAPIListRequest) PageId(pageId string) RecordMxAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r RecordMxAPIListRequest) Filters(filters map[string]interface{}) RecordMxAPIListRequest {
	r.filters = &filters
	return r
}

func (r RecordMxAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) RecordMxAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r RecordMxAPIListRequest) Execute() (*ListRecordMxResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve record:mx objects

Returns a list of record:mx objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordMxAPIListRequest
*/
func (a *RecordMxAPIService) List(ctx context.Context) RecordMxAPIListRequest {
	return RecordMxAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListRecordMxResponse
func (a *RecordMxAPIService) ListExecute(r RecordMxAPIListRequest) (*ListRecordMxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListRecordMxResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordMxAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:mx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.maxResults != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_max_results", r.maxResults, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	if r.paging != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_paging", r.paging, "form", "")
	}
	if r.pageId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_id", r.pageId, "form", "")
	}
	if r.filters != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "filters", r.filters, "form", "")
	}
	if r.extattrfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "extattrfilter", r.extattrfilter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordMxAPIReadRequest struct {
	ctx              context.Context
	ApiService       RecordMxAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r RecordMxAPIReadRequest) ReturnFields(returnFields string) RecordMxAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordMxAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordMxAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordMxAPIReadRequest) ReturnAsObject(returnAsObject int32) RecordMxAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordMxAPIReadRequest) Execute() (*GetRecordMxResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific record:mx object

Returns a specific record:mx object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:mx object
	@return RecordMxAPIReadRequest
*/
func (a *RecordMxAPIService) Read(ctx context.Context, reference string) RecordMxAPIReadRequest {
	return RecordMxAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetRecordMxResponse
func (a *RecordMxAPIService) ReadExecute(r RecordMxAPIReadRequest) (*GetRecordMxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetRecordMxResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordMxAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:mx/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordMxAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       RecordMxAPI
	reference        string
	recordMx         *RecordMx
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r RecordMxAPIUpdateRequest) RecordMx(recordMx RecordMx) RecordMxAPIUpdateRequest {
	r.recordMx = &recordMx
	return r
}

// Enter the field names followed by comma
func (r RecordMxAPIUpdateRequest) ReturnFields(returnFields string) RecordMxAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordMxAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) RecordMxAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r RecordMxAPIUpdateRequest) ReturnAsObject(returnAsObject int32) RecordMxAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordMxAPIUpdateRequest) Execute() (*UpdateRecordMxResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a record:mx object

Updates a specific record:mx object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the record:mx object
	@return RecordMxAPIUpdateRequest
*/
func (a *RecordMxAPIService) Update(ctx context.Context, reference string) RecordMxAPIUpdateRequest {
	return RecordMxAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateRecordMxResponse
func (a *RecordMxAPIService) UpdateExecute(r RecordMxAPIUpdateRequest) (*UpdateRecordMxResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateRecordMxResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordMxAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:mx/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordMx == nil {
		return localVarReturnValue, nil, internal.ReportError("recordMx is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "form", "")
	}
	if r.returnFieldsPlus != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFieldsPlus, "form", "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.recordMx != nil {
		if r.recordMx.ExtAttrs == nil {
			r.recordMx.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.recordMx.ExtAttrs)[k]; !ok {
				(*r.recordMx.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.recordMx
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
