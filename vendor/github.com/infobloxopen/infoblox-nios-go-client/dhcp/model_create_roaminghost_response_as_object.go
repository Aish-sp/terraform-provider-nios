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

// checks if the CreateRoaminghostResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoaminghostResponseAsObject{}

// CreateRoaminghostResponseAsObject The response format to create __Roaminghost__ in object format.
type CreateRoaminghostResponseAsObject struct {
	Result               *Roaminghost `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRoaminghostResponseAsObject CreateRoaminghostResponseAsObject

// NewCreateRoaminghostResponseAsObject instantiates a new CreateRoaminghostResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoaminghostResponseAsObject() *CreateRoaminghostResponseAsObject {
	this := CreateRoaminghostResponseAsObject{}
	return &this
}

// NewCreateRoaminghostResponseAsObjectWithDefaults instantiates a new CreateRoaminghostResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoaminghostResponseAsObjectWithDefaults() *CreateRoaminghostResponseAsObject {
	this := CreateRoaminghostResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateRoaminghostResponseAsObject) GetResult() Roaminghost {
	if o == nil || IsNil(o.Result) {
		var ret Roaminghost
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoaminghostResponseAsObject) GetResultOk() (*Roaminghost, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateRoaminghostResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Roaminghost and assigns it to the Result field.
func (o *CreateRoaminghostResponseAsObject) SetResult(v Roaminghost) {
	o.Result = &v
}

func (o CreateRoaminghostResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoaminghostResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRoaminghostResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateRoaminghostResponseAsObject := _CreateRoaminghostResponseAsObject{}

	err = json.Unmarshal(data, &varCreateRoaminghostResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateRoaminghostResponseAsObject(varCreateRoaminghostResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRoaminghostResponseAsObject struct {
	value *CreateRoaminghostResponseAsObject
	isSet bool
}

func (v NullableCreateRoaminghostResponseAsObject) Get() *CreateRoaminghostResponseAsObject {
	return v.value
}

func (v *NullableCreateRoaminghostResponseAsObject) Set(val *CreateRoaminghostResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoaminghostResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoaminghostResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoaminghostResponseAsObject(val *CreateRoaminghostResponseAsObject) *NullableCreateRoaminghostResponseAsObject {
	return &NullableCreateRoaminghostResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateRoaminghostResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoaminghostResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
