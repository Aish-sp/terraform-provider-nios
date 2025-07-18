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

// checks if the CreateRecordAliasResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordAliasResponseAsObject{}

// CreateRecordAliasResponseAsObject The response format to create __RecordAlias__ in object format.
type CreateRecordAliasResponseAsObject struct {
	Result               *RecordAlias `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecordAliasResponseAsObject CreateRecordAliasResponseAsObject

// NewCreateRecordAliasResponseAsObject instantiates a new CreateRecordAliasResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordAliasResponseAsObject() *CreateRecordAliasResponseAsObject {
	this := CreateRecordAliasResponseAsObject{}
	return &this
}

// NewCreateRecordAliasResponseAsObjectWithDefaults instantiates a new CreateRecordAliasResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordAliasResponseAsObjectWithDefaults() *CreateRecordAliasResponseAsObject {
	this := CreateRecordAliasResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateRecordAliasResponseAsObject) GetResult() RecordAlias {
	if o == nil || IsNil(o.Result) {
		var ret RecordAlias
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordAliasResponseAsObject) GetResultOk() (*RecordAlias, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateRecordAliasResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordAlias and assigns it to the Result field.
func (o *CreateRecordAliasResponseAsObject) SetResult(v RecordAlias) {
	o.Result = &v
}

func (o CreateRecordAliasResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordAliasResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecordAliasResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateRecordAliasResponseAsObject := _CreateRecordAliasResponseAsObject{}

	err = json.Unmarshal(data, &varCreateRecordAliasResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateRecordAliasResponseAsObject(varCreateRecordAliasResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecordAliasResponseAsObject struct {
	value *CreateRecordAliasResponseAsObject
	isSet bool
}

func (v NullableCreateRecordAliasResponseAsObject) Get() *CreateRecordAliasResponseAsObject {
	return v.value
}

func (v *NullableCreateRecordAliasResponseAsObject) Set(val *CreateRecordAliasResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordAliasResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordAliasResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordAliasResponseAsObject(val *CreateRecordAliasResponseAsObject) *NullableCreateRecordAliasResponseAsObject {
	return &NullableCreateRecordAliasResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateRecordAliasResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordAliasResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
