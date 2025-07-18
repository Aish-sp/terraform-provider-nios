/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the UpdateNetworktemplateResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworktemplateResponseAsObject{}

// UpdateNetworktemplateResponseAsObject The response format to update __Networktemplate__ in object format.
type UpdateNetworktemplateResponseAsObject struct {
	Result               *Networktemplate `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworktemplateResponseAsObject UpdateNetworktemplateResponseAsObject

// NewUpdateNetworktemplateResponseAsObject instantiates a new UpdateNetworktemplateResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworktemplateResponseAsObject() *UpdateNetworktemplateResponseAsObject {
	this := UpdateNetworktemplateResponseAsObject{}
	return &this
}

// NewUpdateNetworktemplateResponseAsObjectWithDefaults instantiates a new UpdateNetworktemplateResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworktemplateResponseAsObjectWithDefaults() *UpdateNetworktemplateResponseAsObject {
	this := UpdateNetworktemplateResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateNetworktemplateResponseAsObject) GetResult() Networktemplate {
	if o == nil || IsNil(o.Result) {
		var ret Networktemplate
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworktemplateResponseAsObject) GetResultOk() (*Networktemplate, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateNetworktemplateResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Networktemplate and assigns it to the Result field.
func (o *UpdateNetworktemplateResponseAsObject) SetResult(v Networktemplate) {
	o.Result = &v
}

func (o UpdateNetworktemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworktemplateResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworktemplateResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateNetworktemplateResponseAsObject := _UpdateNetworktemplateResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateNetworktemplateResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateNetworktemplateResponseAsObject(varUpdateNetworktemplateResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworktemplateResponseAsObject struct {
	value *UpdateNetworktemplateResponseAsObject
	isSet bool
}

func (v NullableUpdateNetworktemplateResponseAsObject) Get() *UpdateNetworktemplateResponseAsObject {
	return v.value
}

func (v *NullableUpdateNetworktemplateResponseAsObject) Set(val *UpdateNetworktemplateResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworktemplateResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworktemplateResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworktemplateResponseAsObject(val *UpdateNetworktemplateResponseAsObject) *NullableUpdateNetworktemplateResponseAsObject {
	return &NullableUpdateNetworktemplateResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateNetworktemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworktemplateResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
