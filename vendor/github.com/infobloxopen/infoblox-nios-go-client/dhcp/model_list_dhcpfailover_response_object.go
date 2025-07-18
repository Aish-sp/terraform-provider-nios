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

// checks if the ListDhcpfailoverResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDhcpfailoverResponseObject{}

// ListDhcpfailoverResponseObject The response format to retrieve __Dhcpfailover__ objects.
type ListDhcpfailoverResponseObject struct {
	Result               []Dhcpfailover `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDhcpfailoverResponseObject ListDhcpfailoverResponseObject

// NewListDhcpfailoverResponseObject instantiates a new ListDhcpfailoverResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDhcpfailoverResponseObject() *ListDhcpfailoverResponseObject {
	this := ListDhcpfailoverResponseObject{}
	return &this
}

// NewListDhcpfailoverResponseObjectWithDefaults instantiates a new ListDhcpfailoverResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDhcpfailoverResponseObjectWithDefaults() *ListDhcpfailoverResponseObject {
	this := ListDhcpfailoverResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListDhcpfailoverResponseObject) GetResult() []Dhcpfailover {
	if o == nil || IsNil(o.Result) {
		var ret []Dhcpfailover
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDhcpfailoverResponseObject) GetResultOk() ([]Dhcpfailover, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListDhcpfailoverResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Dhcpfailover and assigns it to the Result field.
func (o *ListDhcpfailoverResponseObject) SetResult(v []Dhcpfailover) {
	o.Result = v
}

func (o ListDhcpfailoverResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDhcpfailoverResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDhcpfailoverResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListDhcpfailoverResponseObject := _ListDhcpfailoverResponseObject{}

	err = json.Unmarshal(data, &varListDhcpfailoverResponseObject)

	if err != nil {
		return err
	}

	*o = ListDhcpfailoverResponseObject(varListDhcpfailoverResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDhcpfailoverResponseObject struct {
	value *ListDhcpfailoverResponseObject
	isSet bool
}

func (v NullableListDhcpfailoverResponseObject) Get() *ListDhcpfailoverResponseObject {
	return v.value
}

func (v *NullableListDhcpfailoverResponseObject) Set(val *ListDhcpfailoverResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListDhcpfailoverResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListDhcpfailoverResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDhcpfailoverResponseObject(val *ListDhcpfailoverResponseObject) *NullableListDhcpfailoverResponseObject {
	return &NullableListDhcpfailoverResponseObject{value: val, isSet: true}
}

func (v NullableListDhcpfailoverResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDhcpfailoverResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
