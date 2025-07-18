/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the GetRecordRrsigResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordRrsigResponseObjectAsResult{}

// GetRecordRrsigResponseObjectAsResult The response format to retrieve __RecordRrsig__ objects.
type GetRecordRrsigResponseObjectAsResult struct {
	Result *RecordRrsig `json:"result,omitempty"`
}

// NewGetRecordRrsigResponseObjectAsResult instantiates a new GetRecordRrsigResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordRrsigResponseObjectAsResult() *GetRecordRrsigResponseObjectAsResult {
	this := GetRecordRrsigResponseObjectAsResult{}
	return &this
}

// NewGetRecordRrsigResponseObjectAsResultWithDefaults instantiates a new GetRecordRrsigResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordRrsigResponseObjectAsResultWithDefaults() *GetRecordRrsigResponseObjectAsResult {
	this := GetRecordRrsigResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetRecordRrsigResponseObjectAsResult) GetResult() RecordRrsig {
	if o == nil || IsNil(o.Result) {
		var ret RecordRrsig
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordRrsigResponseObjectAsResult) GetResultOk() (*RecordRrsig, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetRecordRrsigResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordRrsig and assigns it to the Result field.
func (o *GetRecordRrsigResponseObjectAsResult) SetResult(v RecordRrsig) {
	o.Result = &v
}

func (o GetRecordRrsigResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordRrsigResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetRecordRrsigResponseObjectAsResult struct {
	value *GetRecordRrsigResponseObjectAsResult
	isSet bool
}

func (v NullableGetRecordRrsigResponseObjectAsResult) Get() *GetRecordRrsigResponseObjectAsResult {
	return v.value
}

func (v *NullableGetRecordRrsigResponseObjectAsResult) Set(val *GetRecordRrsigResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordRrsigResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordRrsigResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordRrsigResponseObjectAsResult(val *GetRecordRrsigResponseObjectAsResult) *NullableGetRecordRrsigResponseObjectAsResult {
	return &NullableGetRecordRrsigResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetRecordRrsigResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordRrsigResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
