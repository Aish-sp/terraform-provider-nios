/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

type DtcServerAPI interface {
	/*
		Create Create a dtc:server object

		Creates a new dtc:server object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcServerAPICreateRequest
	*/
	Create(ctx context.Context) DtcServerAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateDtcServerResponse
	CreateExecute(r DtcServerAPICreateRequest) (*CreateDtcServerResponse, *http.Response, error)
	/*
		Delete Delete a dtc:server object

		Deletes a specific dtc:server object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:server object
		@return DtcServerAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) DtcServerAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r DtcServerAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve dtc:server objects

		Returns a list of dtc:server objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return DtcServerAPIListRequest
	*/
	List(ctx context.Context) DtcServerAPIListRequest

	// ListExecute executes the request
	//  @return ListDtcServerResponse
	ListExecute(r DtcServerAPIListRequest) (*ListDtcServerResponse, *http.Response, error)
	/*
		Read Get a specific dtc:server object

		Returns a specific dtc:server object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:server object
		@return DtcServerAPIReadRequest
	*/
	Read(ctx context.Context, reference string) DtcServerAPIReadRequest

	// ReadExecute executes the request
	//  @return GetDtcServerResponse
	ReadExecute(r DtcServerAPIReadRequest) (*GetDtcServerResponse, *http.Response, error)
	/*
		Update Update a dtc:server object

		Updates a specific dtc:server object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the dtc:server object
		@return DtcServerAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) DtcServerAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateDtcServerResponse
	UpdateExecute(r DtcServerAPIUpdateRequest) (*UpdateDtcServerResponse, *http.Response, error)
}

// DtcServerAPIService DtcServerAPI service
type DtcServerAPIService internal.Service

type DtcServerAPICreateRequest struct {
	ctx              context.Context
	ApiService       DtcServerAPI
	dtcServer        *DtcServer
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r DtcServerAPICreateRequest) DtcServer(dtcServer DtcServer) DtcServerAPICreateRequest {
	r.dtcServer = &dtcServer
	return r
}

// Enter the field names followed by comma
func (r DtcServerAPICreateRequest) ReturnFields(returnFields string) DtcServerAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcServerAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcServerAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcServerAPICreateRequest) ReturnAsObject(returnAsObject int32) DtcServerAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcServerAPICreateRequest) Execute() (*CreateDtcServerResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a dtc:server object

Creates a new dtc:server object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcServerAPICreateRequest
*/
func (a *DtcServerAPIService) Create(ctx context.Context) DtcServerAPICreateRequest {
	return DtcServerAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateDtcServerResponse
func (a *DtcServerAPIService) CreateExecute(r DtcServerAPICreateRequest) (*CreateDtcServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateDtcServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcServerAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcServer == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcServer is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcServer != nil {
		if r.dtcServer.ExtAttrs == nil {
			r.dtcServer.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcServer.ExtAttrs)[k]; !ok {
				(*r.dtcServer.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcServer
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

type DtcServerAPIDeleteRequest struct {
	ctx        context.Context
	ApiService DtcServerAPI
	reference  string
}

func (r DtcServerAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a dtc:server object

Deletes a specific dtc:server object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:server object
	@return DtcServerAPIDeleteRequest
*/
func (a *DtcServerAPIService) Delete(ctx context.Context, reference string) DtcServerAPIDeleteRequest {
	return DtcServerAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *DtcServerAPIService) DeleteExecute(r DtcServerAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcServerAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:server/{reference}"
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

type DtcServerAPIListRequest struct {
	ctx              context.Context
	ApiService       DtcServerAPI
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
func (r DtcServerAPIListRequest) ReturnFields(returnFields string) DtcServerAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcServerAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcServerAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r DtcServerAPIListRequest) MaxResults(maxResults int32) DtcServerAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r DtcServerAPIListRequest) ReturnAsObject(returnAsObject int32) DtcServerAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r DtcServerAPIListRequest) Paging(paging int32) DtcServerAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r DtcServerAPIListRequest) PageId(pageId string) DtcServerAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r DtcServerAPIListRequest) Filters(filters map[string]interface{}) DtcServerAPIListRequest {
	r.filters = &filters
	return r
}

func (r DtcServerAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) DtcServerAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r DtcServerAPIListRequest) Execute() (*ListDtcServerResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve dtc:server objects

Returns a list of dtc:server objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DtcServerAPIListRequest
*/
func (a *DtcServerAPIService) List(ctx context.Context) DtcServerAPIListRequest {
	return DtcServerAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListDtcServerResponse
func (a *DtcServerAPIService) ListExecute(r DtcServerAPIListRequest) (*ListDtcServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListDtcServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcServerAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:server"

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

type DtcServerAPIReadRequest struct {
	ctx              context.Context
	ApiService       DtcServerAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r DtcServerAPIReadRequest) ReturnFields(returnFields string) DtcServerAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcServerAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcServerAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcServerAPIReadRequest) ReturnAsObject(returnAsObject int32) DtcServerAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcServerAPIReadRequest) Execute() (*GetDtcServerResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific dtc:server object

Returns a specific dtc:server object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:server object
	@return DtcServerAPIReadRequest
*/
func (a *DtcServerAPIService) Read(ctx context.Context, reference string) DtcServerAPIReadRequest {
	return DtcServerAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetDtcServerResponse
func (a *DtcServerAPIService) ReadExecute(r DtcServerAPIReadRequest) (*GetDtcServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetDtcServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcServerAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:server/{reference}"
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

type DtcServerAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       DtcServerAPI
	reference        string
	dtcServer        *DtcServer
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r DtcServerAPIUpdateRequest) DtcServer(dtcServer DtcServer) DtcServerAPIUpdateRequest {
	r.dtcServer = &dtcServer
	return r
}

// Enter the field names followed by comma
func (r DtcServerAPIUpdateRequest) ReturnFields(returnFields string) DtcServerAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r DtcServerAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) DtcServerAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r DtcServerAPIUpdateRequest) ReturnAsObject(returnAsObject int32) DtcServerAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r DtcServerAPIUpdateRequest) Execute() (*UpdateDtcServerResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a dtc:server object

Updates a specific dtc:server object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the dtc:server object
	@return DtcServerAPIUpdateRequest
*/
func (a *DtcServerAPIService) Update(ctx context.Context, reference string) DtcServerAPIUpdateRequest {
	return DtcServerAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateDtcServerResponse
func (a *DtcServerAPIService) UpdateExecute(r DtcServerAPIUpdateRequest) (*UpdateDtcServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateDtcServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DtcServerAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dtc:server/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dtcServer == nil {
		return localVarReturnValue, nil, internal.ReportError("dtcServer is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.dtcServer != nil {
		if r.dtcServer.ExtAttrs == nil {
			r.dtcServer.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.dtcServer.ExtAttrs)[k]; !ok {
				(*r.dtcServer.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.dtcServer
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
