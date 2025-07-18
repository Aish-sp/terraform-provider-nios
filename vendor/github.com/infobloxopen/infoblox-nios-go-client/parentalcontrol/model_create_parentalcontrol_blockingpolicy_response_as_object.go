/*
Infoblox PARENTALCONTROL API

OpenAPI specification for Infoblox NIOS WAPI PARENTALCONTROL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package parentalcontrol

import (
	"encoding/json"
)

// checks if the CreateParentalcontrolBlockingpolicyResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateParentalcontrolBlockingpolicyResponseAsObject{}

// CreateParentalcontrolBlockingpolicyResponseAsObject The response format to create __ParentalcontrolBlockingpolicy__ in object format.
type CreateParentalcontrolBlockingpolicyResponseAsObject struct {
	Result               *ParentalcontrolBlockingpolicy `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateParentalcontrolBlockingpolicyResponseAsObject CreateParentalcontrolBlockingpolicyResponseAsObject

// NewCreateParentalcontrolBlockingpolicyResponseAsObject instantiates a new CreateParentalcontrolBlockingpolicyResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateParentalcontrolBlockingpolicyResponseAsObject() *CreateParentalcontrolBlockingpolicyResponseAsObject {
	this := CreateParentalcontrolBlockingpolicyResponseAsObject{}
	return &this
}

// NewCreateParentalcontrolBlockingpolicyResponseAsObjectWithDefaults instantiates a new CreateParentalcontrolBlockingpolicyResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateParentalcontrolBlockingpolicyResponseAsObjectWithDefaults() *CreateParentalcontrolBlockingpolicyResponseAsObject {
	this := CreateParentalcontrolBlockingpolicyResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateParentalcontrolBlockingpolicyResponseAsObject) GetResult() ParentalcontrolBlockingpolicy {
	if o == nil || IsNil(o.Result) {
		var ret ParentalcontrolBlockingpolicy
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateParentalcontrolBlockingpolicyResponseAsObject) GetResultOk() (*ParentalcontrolBlockingpolicy, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateParentalcontrolBlockingpolicyResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ParentalcontrolBlockingpolicy and assigns it to the Result field.
func (o *CreateParentalcontrolBlockingpolicyResponseAsObject) SetResult(v ParentalcontrolBlockingpolicy) {
	o.Result = &v
}

func (o CreateParentalcontrolBlockingpolicyResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateParentalcontrolBlockingpolicyResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateParentalcontrolBlockingpolicyResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateParentalcontrolBlockingpolicyResponseAsObject := _CreateParentalcontrolBlockingpolicyResponseAsObject{}

	err = json.Unmarshal(data, &varCreateParentalcontrolBlockingpolicyResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateParentalcontrolBlockingpolicyResponseAsObject(varCreateParentalcontrolBlockingpolicyResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateParentalcontrolBlockingpolicyResponseAsObject struct {
	value *CreateParentalcontrolBlockingpolicyResponseAsObject
	isSet bool
}

func (v NullableCreateParentalcontrolBlockingpolicyResponseAsObject) Get() *CreateParentalcontrolBlockingpolicyResponseAsObject {
	return v.value
}

func (v *NullableCreateParentalcontrolBlockingpolicyResponseAsObject) Set(val *CreateParentalcontrolBlockingpolicyResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateParentalcontrolBlockingpolicyResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateParentalcontrolBlockingpolicyResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateParentalcontrolBlockingpolicyResponseAsObject(val *CreateParentalcontrolBlockingpolicyResponseAsObject) *NullableCreateParentalcontrolBlockingpolicyResponseAsObject {
	return &NullableCreateParentalcontrolBlockingpolicyResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateParentalcontrolBlockingpolicyResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateParentalcontrolBlockingpolicyResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
