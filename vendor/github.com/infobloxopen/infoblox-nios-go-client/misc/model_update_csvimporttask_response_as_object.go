/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the UpdateCsvimporttaskResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCsvimporttaskResponseAsObject{}

// UpdateCsvimporttaskResponseAsObject The response format to update __Csvimporttask__ in object format.
type UpdateCsvimporttaskResponseAsObject struct {
	Result               *Csvimporttask `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateCsvimporttaskResponseAsObject UpdateCsvimporttaskResponseAsObject

// NewUpdateCsvimporttaskResponseAsObject instantiates a new UpdateCsvimporttaskResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCsvimporttaskResponseAsObject() *UpdateCsvimporttaskResponseAsObject {
	this := UpdateCsvimporttaskResponseAsObject{}
	return &this
}

// NewUpdateCsvimporttaskResponseAsObjectWithDefaults instantiates a new UpdateCsvimporttaskResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCsvimporttaskResponseAsObjectWithDefaults() *UpdateCsvimporttaskResponseAsObject {
	this := UpdateCsvimporttaskResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateCsvimporttaskResponseAsObject) GetResult() Csvimporttask {
	if o == nil || IsNil(o.Result) {
		var ret Csvimporttask
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCsvimporttaskResponseAsObject) GetResultOk() (*Csvimporttask, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateCsvimporttaskResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Csvimporttask and assigns it to the Result field.
func (o *UpdateCsvimporttaskResponseAsObject) SetResult(v Csvimporttask) {
	o.Result = &v
}

func (o UpdateCsvimporttaskResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCsvimporttaskResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateCsvimporttaskResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateCsvimporttaskResponseAsObject := _UpdateCsvimporttaskResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateCsvimporttaskResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateCsvimporttaskResponseAsObject(varUpdateCsvimporttaskResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateCsvimporttaskResponseAsObject struct {
	value *UpdateCsvimporttaskResponseAsObject
	isSet bool
}

func (v NullableUpdateCsvimporttaskResponseAsObject) Get() *UpdateCsvimporttaskResponseAsObject {
	return v.value
}

func (v *NullableUpdateCsvimporttaskResponseAsObject) Set(val *UpdateCsvimporttaskResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCsvimporttaskResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCsvimporttaskResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCsvimporttaskResponseAsObject(val *UpdateCsvimporttaskResponseAsObject) *NullableUpdateCsvimporttaskResponseAsObject {
	return &NullableUpdateCsvimporttaskResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateCsvimporttaskResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCsvimporttaskResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
