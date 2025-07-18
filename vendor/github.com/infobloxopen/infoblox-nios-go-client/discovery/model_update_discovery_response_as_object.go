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

// checks if the UpdateDiscoveryResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDiscoveryResponseAsObject{}

// UpdateDiscoveryResponseAsObject The response format to update __Discovery__ in object format.
type UpdateDiscoveryResponseAsObject struct {
	Result               *Discovery `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDiscoveryResponseAsObject UpdateDiscoveryResponseAsObject

// NewUpdateDiscoveryResponseAsObject instantiates a new UpdateDiscoveryResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDiscoveryResponseAsObject() *UpdateDiscoveryResponseAsObject {
	this := UpdateDiscoveryResponseAsObject{}
	return &this
}

// NewUpdateDiscoveryResponseAsObjectWithDefaults instantiates a new UpdateDiscoveryResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDiscoveryResponseAsObjectWithDefaults() *UpdateDiscoveryResponseAsObject {
	this := UpdateDiscoveryResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateDiscoveryResponseAsObject) GetResult() Discovery {
	if o == nil || IsNil(o.Result) {
		var ret Discovery
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscoveryResponseAsObject) GetResultOk() (*Discovery, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateDiscoveryResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Discovery and assigns it to the Result field.
func (o *UpdateDiscoveryResponseAsObject) SetResult(v Discovery) {
	o.Result = &v
}

func (o UpdateDiscoveryResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDiscoveryResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDiscoveryResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateDiscoveryResponseAsObject := _UpdateDiscoveryResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateDiscoveryResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateDiscoveryResponseAsObject(varUpdateDiscoveryResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDiscoveryResponseAsObject struct {
	value *UpdateDiscoveryResponseAsObject
	isSet bool
}

func (v NullableUpdateDiscoveryResponseAsObject) Get() *UpdateDiscoveryResponseAsObject {
	return v.value
}

func (v *NullableUpdateDiscoveryResponseAsObject) Set(val *UpdateDiscoveryResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiscoveryResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiscoveryResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiscoveryResponseAsObject(val *UpdateDiscoveryResponseAsObject) *NullableUpdateDiscoveryResponseAsObject {
	return &NullableUpdateDiscoveryResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateDiscoveryResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiscoveryResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
