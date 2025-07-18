/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the UpdateNatgroupResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNatgroupResponseAsObject{}

// UpdateNatgroupResponseAsObject The response format to update __Natgroup__ in object format.
type UpdateNatgroupResponseAsObject struct {
	Result               *Natgroup `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNatgroupResponseAsObject UpdateNatgroupResponseAsObject

// NewUpdateNatgroupResponseAsObject instantiates a new UpdateNatgroupResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNatgroupResponseAsObject() *UpdateNatgroupResponseAsObject {
	this := UpdateNatgroupResponseAsObject{}
	return &this
}

// NewUpdateNatgroupResponseAsObjectWithDefaults instantiates a new UpdateNatgroupResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNatgroupResponseAsObjectWithDefaults() *UpdateNatgroupResponseAsObject {
	this := UpdateNatgroupResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateNatgroupResponseAsObject) GetResult() Natgroup {
	if o == nil || IsNil(o.Result) {
		var ret Natgroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNatgroupResponseAsObject) GetResultOk() (*Natgroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateNatgroupResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Natgroup and assigns it to the Result field.
func (o *UpdateNatgroupResponseAsObject) SetResult(v Natgroup) {
	o.Result = &v
}

func (o UpdateNatgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNatgroupResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNatgroupResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateNatgroupResponseAsObject := _UpdateNatgroupResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateNatgroupResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateNatgroupResponseAsObject(varUpdateNatgroupResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNatgroupResponseAsObject struct {
	value *UpdateNatgroupResponseAsObject
	isSet bool
}

func (v NullableUpdateNatgroupResponseAsObject) Get() *UpdateNatgroupResponseAsObject {
	return v.value
}

func (v *NullableUpdateNatgroupResponseAsObject) Set(val *UpdateNatgroupResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNatgroupResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNatgroupResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNatgroupResponseAsObject(val *UpdateNatgroupResponseAsObject) *NullableUpdateNatgroupResponseAsObject {
	return &NullableUpdateNatgroupResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateNatgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNatgroupResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
