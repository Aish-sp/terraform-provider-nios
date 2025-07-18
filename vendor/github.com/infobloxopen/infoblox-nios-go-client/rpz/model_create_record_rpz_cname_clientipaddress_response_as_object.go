/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
)

// checks if the CreateRecordRpzCnameClientipaddressResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordRpzCnameClientipaddressResponseAsObject{}

// CreateRecordRpzCnameClientipaddressResponseAsObject The response format to create __RecordRpzCnameClientipaddress__ in object format.
type CreateRecordRpzCnameClientipaddressResponseAsObject struct {
	Result               *RecordRpzCnameClientipaddress `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecordRpzCnameClientipaddressResponseAsObject CreateRecordRpzCnameClientipaddressResponseAsObject

// NewCreateRecordRpzCnameClientipaddressResponseAsObject instantiates a new CreateRecordRpzCnameClientipaddressResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordRpzCnameClientipaddressResponseAsObject() *CreateRecordRpzCnameClientipaddressResponseAsObject {
	this := CreateRecordRpzCnameClientipaddressResponseAsObject{}
	return &this
}

// NewCreateRecordRpzCnameClientipaddressResponseAsObjectWithDefaults instantiates a new CreateRecordRpzCnameClientipaddressResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordRpzCnameClientipaddressResponseAsObjectWithDefaults() *CreateRecordRpzCnameClientipaddressResponseAsObject {
	this := CreateRecordRpzCnameClientipaddressResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateRecordRpzCnameClientipaddressResponseAsObject) GetResult() RecordRpzCnameClientipaddress {
	if o == nil || IsNil(o.Result) {
		var ret RecordRpzCnameClientipaddress
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRpzCnameClientipaddressResponseAsObject) GetResultOk() (*RecordRpzCnameClientipaddress, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateRecordRpzCnameClientipaddressResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordRpzCnameClientipaddress and assigns it to the Result field.
func (o *CreateRecordRpzCnameClientipaddressResponseAsObject) SetResult(v RecordRpzCnameClientipaddress) {
	o.Result = &v
}

func (o CreateRecordRpzCnameClientipaddressResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordRpzCnameClientipaddressResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecordRpzCnameClientipaddressResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateRecordRpzCnameClientipaddressResponseAsObject := _CreateRecordRpzCnameClientipaddressResponseAsObject{}

	err = json.Unmarshal(data, &varCreateRecordRpzCnameClientipaddressResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateRecordRpzCnameClientipaddressResponseAsObject(varCreateRecordRpzCnameClientipaddressResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecordRpzCnameClientipaddressResponseAsObject struct {
	value *CreateRecordRpzCnameClientipaddressResponseAsObject
	isSet bool
}

func (v NullableCreateRecordRpzCnameClientipaddressResponseAsObject) Get() *CreateRecordRpzCnameClientipaddressResponseAsObject {
	return v.value
}

func (v *NullableCreateRecordRpzCnameClientipaddressResponseAsObject) Set(val *CreateRecordRpzCnameClientipaddressResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordRpzCnameClientipaddressResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordRpzCnameClientipaddressResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordRpzCnameClientipaddressResponseAsObject(val *CreateRecordRpzCnameClientipaddressResponseAsObject) *NullableCreateRecordRpzCnameClientipaddressResponseAsObject {
	return &NullableCreateRecordRpzCnameClientipaddressResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateRecordRpzCnameClientipaddressResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordRpzCnameClientipaddressResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
