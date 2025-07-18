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

type ZoneAuthAPI interface {
	/*
		Create Create a zone_auth object

		Creates a new zone_auth object

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ZoneAuthAPICreateRequest
	*/
	Create(ctx context.Context) ZoneAuthAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateZoneAuthResponse
	CreateExecute(r ZoneAuthAPICreateRequest) (*CreateZoneAuthResponse, *http.Response, error)
	/*
		Delete Delete a zone_auth object

		Deletes a specific zone_auth object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the zone_auth object
		@return ZoneAuthAPIDeleteRequest
	*/
	Delete(ctx context.Context, reference string) ZoneAuthAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r ZoneAuthAPIDeleteRequest) (*http.Response, error)
	/*
		List Retrieve zone_auth objects

		Returns a list of zone_auth objects matching the search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ZoneAuthAPIListRequest
	*/
	List(ctx context.Context) ZoneAuthAPIListRequest

	// ListExecute executes the request
	//  @return ListZoneAuthResponse
	ListExecute(r ZoneAuthAPIListRequest) (*ListZoneAuthResponse, *http.Response, error)
	/*
		Read Get a specific zone_auth object

		Returns a specific zone_auth object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the zone_auth object
		@return ZoneAuthAPIReadRequest
	*/
	Read(ctx context.Context, reference string) ZoneAuthAPIReadRequest

	// ReadExecute executes the request
	//  @return GetZoneAuthResponse
	ReadExecute(r ZoneAuthAPIReadRequest) (*GetZoneAuthResponse, *http.Response, error)
	/*
		Update Update a zone_auth object

		Updates a specific zone_auth object by reference

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reference Reference of the zone_auth object
		@return ZoneAuthAPIUpdateRequest
	*/
	Update(ctx context.Context, reference string) ZoneAuthAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateZoneAuthResponse
	UpdateExecute(r ZoneAuthAPIUpdateRequest) (*UpdateZoneAuthResponse, *http.Response, error)
}

// ZoneAuthAPIService ZoneAuthAPI service
type ZoneAuthAPIService internal.Service

type ZoneAuthAPICreateRequest struct {
	ctx              context.Context
	ApiService       ZoneAuthAPI
	zoneAuth         *ZoneAuth
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to create
func (r ZoneAuthAPICreateRequest) ZoneAuth(zoneAuth ZoneAuth) ZoneAuthAPICreateRequest {
	r.zoneAuth = &zoneAuth
	return r
}

// Enter the field names followed by comma
func (r ZoneAuthAPICreateRequest) ReturnFields(returnFields string) ZoneAuthAPICreateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ZoneAuthAPICreateRequest) ReturnFieldsPlus(returnFieldsPlus string) ZoneAuthAPICreateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r ZoneAuthAPICreateRequest) ReturnAsObject(returnAsObject int32) ZoneAuthAPICreateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r ZoneAuthAPICreateRequest) Execute() (*CreateZoneAuthResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a zone_auth object

Creates a new zone_auth object

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ZoneAuthAPICreateRequest
*/
func (a *ZoneAuthAPIService) Create(ctx context.Context) ZoneAuthAPICreateRequest {
	return ZoneAuthAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateZoneAuthResponse
func (a *ZoneAuthAPIService) CreateExecute(r ZoneAuthAPICreateRequest) (*CreateZoneAuthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateZoneAuthResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ZoneAuthAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/zone_auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneAuth == nil {
		return localVarReturnValue, nil, internal.ReportError("zoneAuth is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.zoneAuth != nil {
		if r.zoneAuth.ExtAttrs == nil {
			r.zoneAuth.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.zoneAuth.ExtAttrs)[k]; !ok {
				(*r.zoneAuth.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.zoneAuth
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

type ZoneAuthAPIDeleteRequest struct {
	ctx            context.Context
	ApiService     ZoneAuthAPI
	reference      string
	removeSubzones *bool
}

// Remove subzones delete option. Determines whether all child objects should be removed alongside with the parent zone or child objects should be assigned to another parental zone. By default child objects are deleted with the parent zone.
func (r ZoneAuthAPIDeleteRequest) RemoveSubzones(removeSubzones bool) ZoneAuthAPIDeleteRequest {
	r.removeSubzones = &removeSubzones
	return r
}

func (r ZoneAuthAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete a zone_auth object

Deletes a specific zone_auth object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the zone_auth object
	@return ZoneAuthAPIDeleteRequest
*/
func (a *ZoneAuthAPIService) Delete(ctx context.Context, reference string) ZoneAuthAPIDeleteRequest {
	return ZoneAuthAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
func (a *ZoneAuthAPIService) DeleteExecute(r ZoneAuthAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ZoneAuthAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/zone_auth/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removeSubzones != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "remove_subzones", r.removeSubzones, "form", "")
	}
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

type ZoneAuthAPIListRequest struct {
	ctx              context.Context
	ApiService       ZoneAuthAPI
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
func (r ZoneAuthAPIListRequest) ReturnFields(returnFields string) ZoneAuthAPIListRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ZoneAuthAPIListRequest) ReturnFieldsPlus(returnFieldsPlus string) ZoneAuthAPIListRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Enter the number of results to be fetched
func (r ZoneAuthAPIListRequest) MaxResults(maxResults int32) ZoneAuthAPIListRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r ZoneAuthAPIListRequest) ReturnAsObject(returnAsObject int32) ZoneAuthAPIListRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Control paging of results
func (r ZoneAuthAPIListRequest) Paging(paging int32) ZoneAuthAPIListRequest {
	r.paging = &paging
	return r
}

// Page id for retrieving next page of results
func (r ZoneAuthAPIListRequest) PageId(pageId string) ZoneAuthAPIListRequest {
	r.pageId = &pageId
	return r
}

func (r ZoneAuthAPIListRequest) Filters(filters map[string]interface{}) ZoneAuthAPIListRequest {
	r.filters = &filters
	return r
}

func (r ZoneAuthAPIListRequest) Extattrfilter(extattrfilter map[string]interface{}) ZoneAuthAPIListRequest {
	r.extattrfilter = &extattrfilter
	return r
}

func (r ZoneAuthAPIListRequest) Execute() (*ListZoneAuthResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve zone_auth objects

Returns a list of zone_auth objects matching the search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ZoneAuthAPIListRequest
*/
func (a *ZoneAuthAPIService) List(ctx context.Context) ZoneAuthAPIListRequest {
	return ZoneAuthAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListZoneAuthResponse
func (a *ZoneAuthAPIService) ListExecute(r ZoneAuthAPIListRequest) (*ListZoneAuthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListZoneAuthResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ZoneAuthAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/zone_auth"

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

type ZoneAuthAPIReadRequest struct {
	ctx              context.Context
	ApiService       ZoneAuthAPI
	reference        string
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Enter the field names followed by comma
func (r ZoneAuthAPIReadRequest) ReturnFields(returnFields string) ZoneAuthAPIReadRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ZoneAuthAPIReadRequest) ReturnFieldsPlus(returnFieldsPlus string) ZoneAuthAPIReadRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r ZoneAuthAPIReadRequest) ReturnAsObject(returnAsObject int32) ZoneAuthAPIReadRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r ZoneAuthAPIReadRequest) Execute() (*GetZoneAuthResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Get a specific zone_auth object

Returns a specific zone_auth object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the zone_auth object
	@return ZoneAuthAPIReadRequest
*/
func (a *ZoneAuthAPIService) Read(ctx context.Context, reference string) ZoneAuthAPIReadRequest {
	return ZoneAuthAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return GetZoneAuthResponse
func (a *ZoneAuthAPIService) ReadExecute(r ZoneAuthAPIReadRequest) (*GetZoneAuthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *GetZoneAuthResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ZoneAuthAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/zone_auth/{reference}"
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

type ZoneAuthAPIUpdateRequest struct {
	ctx              context.Context
	ApiService       ZoneAuthAPI
	reference        string
	zoneAuth         *ZoneAuth
	returnFields     *string
	returnFieldsPlus *string
	returnAsObject   *int32
}

// Object data to update
func (r ZoneAuthAPIUpdateRequest) ZoneAuth(zoneAuth ZoneAuth) ZoneAuthAPIUpdateRequest {
	r.zoneAuth = &zoneAuth
	return r
}

// Enter the field names followed by comma
func (r ZoneAuthAPIUpdateRequest) ReturnFields(returnFields string) ZoneAuthAPIUpdateRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r ZoneAuthAPIUpdateRequest) ReturnFieldsPlus(returnFieldsPlus string) ZoneAuthAPIUpdateRequest {
	r.returnFieldsPlus = &returnFieldsPlus
	return r
}

// Select 1 if result is required as an object
func (r ZoneAuthAPIUpdateRequest) ReturnAsObject(returnAsObject int32) ZoneAuthAPIUpdateRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r ZoneAuthAPIUpdateRequest) Execute() (*UpdateZoneAuthResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a zone_auth object

Updates a specific zone_auth object by reference

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reference Reference of the zone_auth object
	@return ZoneAuthAPIUpdateRequest
*/
func (a *ZoneAuthAPIService) Update(ctx context.Context, reference string) ZoneAuthAPIUpdateRequest {
	return ZoneAuthAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		reference:  reference,
	}
}

// Execute executes the request
//
//	@return UpdateZoneAuthResponse
func (a *ZoneAuthAPIService) UpdateExecute(r ZoneAuthAPIUpdateRequest) (*UpdateZoneAuthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateZoneAuthResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ZoneAuthAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/zone_auth/{reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"reference"+"}", url.PathEscape(internal.ParameterValueToString(r.reference, "reference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneAuth == nil {
		return localVarReturnValue, nil, internal.ReportError("zoneAuth is required and must be specified")
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
	if len(a.Client.Cfg.DefaultExtAttrs) > 0 && r.zoneAuth != nil {
		if r.zoneAuth.ExtAttrs == nil {
			r.zoneAuth.ExtAttrs = &map[string]ExtAttrs{}
		}
		for k, v := range a.Client.Cfg.DefaultExtAttrs {
			if _, ok := (*r.zoneAuth.ExtAttrs)[k]; !ok {
				(*r.zoneAuth.ExtAttrs)[k] = ExtAttrs{
					Value: v.Value,
				}
			}
		}
	}
	// body params
	localVarPostBody = r.zoneAuth
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
