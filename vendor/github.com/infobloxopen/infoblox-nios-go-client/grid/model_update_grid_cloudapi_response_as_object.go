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

// checks if the UpdateGridCloudapiResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGridCloudapiResponseAsObject{}

// UpdateGridCloudapiResponseAsObject The response format to update __GridCloudapi__ in object format.
type UpdateGridCloudapiResponseAsObject struct {
	Result               *GridCloudapi `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateGridCloudapiResponseAsObject UpdateGridCloudapiResponseAsObject

// NewUpdateGridCloudapiResponseAsObject instantiates a new UpdateGridCloudapiResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGridCloudapiResponseAsObject() *UpdateGridCloudapiResponseAsObject {
	this := UpdateGridCloudapiResponseAsObject{}
	return &this
}

// NewUpdateGridCloudapiResponseAsObjectWithDefaults instantiates a new UpdateGridCloudapiResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGridCloudapiResponseAsObjectWithDefaults() *UpdateGridCloudapiResponseAsObject {
	this := UpdateGridCloudapiResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateGridCloudapiResponseAsObject) GetResult() GridCloudapi {
	if o == nil || IsNil(o.Result) {
		var ret GridCloudapi
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGridCloudapiResponseAsObject) GetResultOk() (*GridCloudapi, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateGridCloudapiResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GridCloudapi and assigns it to the Result field.
func (o *UpdateGridCloudapiResponseAsObject) SetResult(v GridCloudapi) {
	o.Result = &v
}

func (o UpdateGridCloudapiResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGridCloudapiResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateGridCloudapiResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateGridCloudapiResponseAsObject := _UpdateGridCloudapiResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateGridCloudapiResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateGridCloudapiResponseAsObject(varUpdateGridCloudapiResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateGridCloudapiResponseAsObject struct {
	value *UpdateGridCloudapiResponseAsObject
	isSet bool
}

func (v NullableUpdateGridCloudapiResponseAsObject) Get() *UpdateGridCloudapiResponseAsObject {
	return v.value
}

func (v *NullableUpdateGridCloudapiResponseAsObject) Set(val *UpdateGridCloudapiResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGridCloudapiResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGridCloudapiResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGridCloudapiResponseAsObject(val *UpdateGridCloudapiResponseAsObject) *NullableUpdateGridCloudapiResponseAsObject {
	return &NullableUpdateGridCloudapiResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateGridCloudapiResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGridCloudapiResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
