/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the GetFiltermacResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFiltermacResponseObjectAsResult{}

// GetFiltermacResponseObjectAsResult The response format to retrieve __Filtermac__ objects.
type GetFiltermacResponseObjectAsResult struct {
	Result *Filtermac `json:"result,omitempty"`
}

// NewGetFiltermacResponseObjectAsResult instantiates a new GetFiltermacResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiltermacResponseObjectAsResult() *GetFiltermacResponseObjectAsResult {
	this := GetFiltermacResponseObjectAsResult{}
	return &this
}

// NewGetFiltermacResponseObjectAsResultWithDefaults instantiates a new GetFiltermacResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiltermacResponseObjectAsResultWithDefaults() *GetFiltermacResponseObjectAsResult {
	this := GetFiltermacResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetFiltermacResponseObjectAsResult) GetResult() Filtermac {
	if o == nil || IsNil(o.Result) {
		var ret Filtermac
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiltermacResponseObjectAsResult) GetResultOk() (*Filtermac, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetFiltermacResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Filtermac and assigns it to the Result field.
func (o *GetFiltermacResponseObjectAsResult) SetResult(v Filtermac) {
	o.Result = &v
}

func (o GetFiltermacResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiltermacResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetFiltermacResponseObjectAsResult struct {
	value *GetFiltermacResponseObjectAsResult
	isSet bool
}

func (v NullableGetFiltermacResponseObjectAsResult) Get() *GetFiltermacResponseObjectAsResult {
	return v.value
}

func (v *NullableGetFiltermacResponseObjectAsResult) Set(val *GetFiltermacResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiltermacResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiltermacResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiltermacResponseObjectAsResult(val *GetFiltermacResponseObjectAsResult) *NullableGetFiltermacResponseObjectAsResult {
	return &NullableGetFiltermacResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetFiltermacResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiltermacResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
