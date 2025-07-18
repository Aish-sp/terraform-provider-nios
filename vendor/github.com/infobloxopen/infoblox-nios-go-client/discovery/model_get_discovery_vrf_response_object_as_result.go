/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
)

// checks if the GetDiscoveryVrfResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDiscoveryVrfResponseObjectAsResult{}

// GetDiscoveryVrfResponseObjectAsResult The response format to retrieve __DiscoveryVrf__ objects.
type GetDiscoveryVrfResponseObjectAsResult struct {
	Result *DiscoveryVrf `json:"result,omitempty"`
}

// NewGetDiscoveryVrfResponseObjectAsResult instantiates a new GetDiscoveryVrfResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoveryVrfResponseObjectAsResult() *GetDiscoveryVrfResponseObjectAsResult {
	this := GetDiscoveryVrfResponseObjectAsResult{}
	return &this
}

// NewGetDiscoveryVrfResponseObjectAsResultWithDefaults instantiates a new GetDiscoveryVrfResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoveryVrfResponseObjectAsResultWithDefaults() *GetDiscoveryVrfResponseObjectAsResult {
	this := GetDiscoveryVrfResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDiscoveryVrfResponseObjectAsResult) GetResult() DiscoveryVrf {
	if o == nil || IsNil(o.Result) {
		var ret DiscoveryVrf
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveryVrfResponseObjectAsResult) GetResultOk() (*DiscoveryVrf, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDiscoveryVrfResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DiscoveryVrf and assigns it to the Result field.
func (o *GetDiscoveryVrfResponseObjectAsResult) SetResult(v DiscoveryVrf) {
	o.Result = &v
}

func (o GetDiscoveryVrfResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDiscoveryVrfResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDiscoveryVrfResponseObjectAsResult struct {
	value *GetDiscoveryVrfResponseObjectAsResult
	isSet bool
}

func (v NullableGetDiscoveryVrfResponseObjectAsResult) Get() *GetDiscoveryVrfResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDiscoveryVrfResponseObjectAsResult) Set(val *GetDiscoveryVrfResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveryVrfResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveryVrfResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveryVrfResponseObjectAsResult(val *GetDiscoveryVrfResponseObjectAsResult) *NullableGetDiscoveryVrfResponseObjectAsResult {
	return &NullableGetDiscoveryVrfResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDiscoveryVrfResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveryVrfResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
