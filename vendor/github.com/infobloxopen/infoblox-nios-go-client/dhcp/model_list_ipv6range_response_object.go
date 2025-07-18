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

// checks if the ListIpv6rangeResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIpv6rangeResponseObject{}

// ListIpv6rangeResponseObject The response format to retrieve __Ipv6range__ objects.
type ListIpv6rangeResponseObject struct {
	Result               []Ipv6range `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListIpv6rangeResponseObject ListIpv6rangeResponseObject

// NewListIpv6rangeResponseObject instantiates a new ListIpv6rangeResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIpv6rangeResponseObject() *ListIpv6rangeResponseObject {
	this := ListIpv6rangeResponseObject{}
	return &this
}

// NewListIpv6rangeResponseObjectWithDefaults instantiates a new ListIpv6rangeResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIpv6rangeResponseObjectWithDefaults() *ListIpv6rangeResponseObject {
	this := ListIpv6rangeResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListIpv6rangeResponseObject) GetResult() []Ipv6range {
	if o == nil || IsNil(o.Result) {
		var ret []Ipv6range
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIpv6rangeResponseObject) GetResultOk() ([]Ipv6range, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListIpv6rangeResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Ipv6range and assigns it to the Result field.
func (o *ListIpv6rangeResponseObject) SetResult(v []Ipv6range) {
	o.Result = v
}

func (o ListIpv6rangeResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIpv6rangeResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListIpv6rangeResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListIpv6rangeResponseObject := _ListIpv6rangeResponseObject{}

	err = json.Unmarshal(data, &varListIpv6rangeResponseObject)

	if err != nil {
		return err
	}

	*o = ListIpv6rangeResponseObject(varListIpv6rangeResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListIpv6rangeResponseObject struct {
	value *ListIpv6rangeResponseObject
	isSet bool
}

func (v NullableListIpv6rangeResponseObject) Get() *ListIpv6rangeResponseObject {
	return v.value
}

func (v *NullableListIpv6rangeResponseObject) Set(val *ListIpv6rangeResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListIpv6rangeResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListIpv6rangeResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIpv6rangeResponseObject(val *ListIpv6rangeResponseObject) *NullableListIpv6rangeResponseObject {
	return &NullableListIpv6rangeResponseObject{value: val, isSet: true}
}

func (v NullableListIpv6rangeResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIpv6rangeResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
