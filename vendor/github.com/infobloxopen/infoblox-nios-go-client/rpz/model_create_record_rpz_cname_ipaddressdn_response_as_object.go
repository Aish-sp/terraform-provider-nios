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

// checks if the CreateRecordRpzCnameIpaddressdnResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordRpzCnameIpaddressdnResponseAsObject{}

// CreateRecordRpzCnameIpaddressdnResponseAsObject The response format to create __RecordRpzCnameIpaddressdn__ in object format.
type CreateRecordRpzCnameIpaddressdnResponseAsObject struct {
	Result               *RecordRpzCnameIpaddressdn `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecordRpzCnameIpaddressdnResponseAsObject CreateRecordRpzCnameIpaddressdnResponseAsObject

// NewCreateRecordRpzCnameIpaddressdnResponseAsObject instantiates a new CreateRecordRpzCnameIpaddressdnResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordRpzCnameIpaddressdnResponseAsObject() *CreateRecordRpzCnameIpaddressdnResponseAsObject {
	this := CreateRecordRpzCnameIpaddressdnResponseAsObject{}
	return &this
}

// NewCreateRecordRpzCnameIpaddressdnResponseAsObjectWithDefaults instantiates a new CreateRecordRpzCnameIpaddressdnResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordRpzCnameIpaddressdnResponseAsObjectWithDefaults() *CreateRecordRpzCnameIpaddressdnResponseAsObject {
	this := CreateRecordRpzCnameIpaddressdnResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateRecordRpzCnameIpaddressdnResponseAsObject) GetResult() RecordRpzCnameIpaddressdn {
	if o == nil || IsNil(o.Result) {
		var ret RecordRpzCnameIpaddressdn
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRpzCnameIpaddressdnResponseAsObject) GetResultOk() (*RecordRpzCnameIpaddressdn, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateRecordRpzCnameIpaddressdnResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordRpzCnameIpaddressdn and assigns it to the Result field.
func (o *CreateRecordRpzCnameIpaddressdnResponseAsObject) SetResult(v RecordRpzCnameIpaddressdn) {
	o.Result = &v
}

func (o CreateRecordRpzCnameIpaddressdnResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordRpzCnameIpaddressdnResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecordRpzCnameIpaddressdnResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateRecordRpzCnameIpaddressdnResponseAsObject := _CreateRecordRpzCnameIpaddressdnResponseAsObject{}

	err = json.Unmarshal(data, &varCreateRecordRpzCnameIpaddressdnResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateRecordRpzCnameIpaddressdnResponseAsObject(varCreateRecordRpzCnameIpaddressdnResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecordRpzCnameIpaddressdnResponseAsObject struct {
	value *CreateRecordRpzCnameIpaddressdnResponseAsObject
	isSet bool
}

func (v NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) Get() *CreateRecordRpzCnameIpaddressdnResponseAsObject {
	return v.value
}

func (v *NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) Set(val *CreateRecordRpzCnameIpaddressdnResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordRpzCnameIpaddressdnResponseAsObject(val *CreateRecordRpzCnameIpaddressdnResponseAsObject) *NullableCreateRecordRpzCnameIpaddressdnResponseAsObject {
	return &NullableCreateRecordRpzCnameIpaddressdnResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordRpzCnameIpaddressdnResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
