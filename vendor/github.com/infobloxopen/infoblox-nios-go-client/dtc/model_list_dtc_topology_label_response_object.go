/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the ListDtcTopologyLabelResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDtcTopologyLabelResponseObject{}

// ListDtcTopologyLabelResponseObject The response format to retrieve __DtcTopologyLabel__ objects.
type ListDtcTopologyLabelResponseObject struct {
	Result               []DtcTopologyLabel `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDtcTopologyLabelResponseObject ListDtcTopologyLabelResponseObject

// NewListDtcTopologyLabelResponseObject instantiates a new ListDtcTopologyLabelResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDtcTopologyLabelResponseObject() *ListDtcTopologyLabelResponseObject {
	this := ListDtcTopologyLabelResponseObject{}
	return &this
}

// NewListDtcTopologyLabelResponseObjectWithDefaults instantiates a new ListDtcTopologyLabelResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDtcTopologyLabelResponseObjectWithDefaults() *ListDtcTopologyLabelResponseObject {
	this := ListDtcTopologyLabelResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListDtcTopologyLabelResponseObject) GetResult() []DtcTopologyLabel {
	if o == nil || IsNil(o.Result) {
		var ret []DtcTopologyLabel
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDtcTopologyLabelResponseObject) GetResultOk() ([]DtcTopologyLabel, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListDtcTopologyLabelResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []DtcTopologyLabel and assigns it to the Result field.
func (o *ListDtcTopologyLabelResponseObject) SetResult(v []DtcTopologyLabel) {
	o.Result = v
}

func (o ListDtcTopologyLabelResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDtcTopologyLabelResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDtcTopologyLabelResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListDtcTopologyLabelResponseObject := _ListDtcTopologyLabelResponseObject{}

	err = json.Unmarshal(data, &varListDtcTopologyLabelResponseObject)

	if err != nil {
		return err
	}

	*o = ListDtcTopologyLabelResponseObject(varListDtcTopologyLabelResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDtcTopologyLabelResponseObject struct {
	value *ListDtcTopologyLabelResponseObject
	isSet bool
}

func (v NullableListDtcTopologyLabelResponseObject) Get() *ListDtcTopologyLabelResponseObject {
	return v.value
}

func (v *NullableListDtcTopologyLabelResponseObject) Set(val *ListDtcTopologyLabelResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListDtcTopologyLabelResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListDtcTopologyLabelResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDtcTopologyLabelResponseObject(val *ListDtcTopologyLabelResponseObject) *NullableListDtcTopologyLabelResponseObject {
	return &NullableListDtcTopologyLabelResponseObject{value: val, isSet: true}
}

func (v NullableListDtcTopologyLabelResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDtcTopologyLabelResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
