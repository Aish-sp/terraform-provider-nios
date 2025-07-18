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

// checks if the CreateIpv6sharednetworkResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIpv6sharednetworkResponseAsObject{}

// CreateIpv6sharednetworkResponseAsObject The response format to create __Ipv6sharednetwork__ in object format.
type CreateIpv6sharednetworkResponseAsObject struct {
	Result               *Ipv6sharednetwork `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateIpv6sharednetworkResponseAsObject CreateIpv6sharednetworkResponseAsObject

// NewCreateIpv6sharednetworkResponseAsObject instantiates a new CreateIpv6sharednetworkResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIpv6sharednetworkResponseAsObject() *CreateIpv6sharednetworkResponseAsObject {
	this := CreateIpv6sharednetworkResponseAsObject{}
	return &this
}

// NewCreateIpv6sharednetworkResponseAsObjectWithDefaults instantiates a new CreateIpv6sharednetworkResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIpv6sharednetworkResponseAsObjectWithDefaults() *CreateIpv6sharednetworkResponseAsObject {
	this := CreateIpv6sharednetworkResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateIpv6sharednetworkResponseAsObject) GetResult() Ipv6sharednetwork {
	if o == nil || IsNil(o.Result) {
		var ret Ipv6sharednetwork
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIpv6sharednetworkResponseAsObject) GetResultOk() (*Ipv6sharednetwork, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateIpv6sharednetworkResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv6sharednetwork and assigns it to the Result field.
func (o *CreateIpv6sharednetworkResponseAsObject) SetResult(v Ipv6sharednetwork) {
	o.Result = &v
}

func (o CreateIpv6sharednetworkResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIpv6sharednetworkResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateIpv6sharednetworkResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateIpv6sharednetworkResponseAsObject := _CreateIpv6sharednetworkResponseAsObject{}

	err = json.Unmarshal(data, &varCreateIpv6sharednetworkResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateIpv6sharednetworkResponseAsObject(varCreateIpv6sharednetworkResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateIpv6sharednetworkResponseAsObject struct {
	value *CreateIpv6sharednetworkResponseAsObject
	isSet bool
}

func (v NullableCreateIpv6sharednetworkResponseAsObject) Get() *CreateIpv6sharednetworkResponseAsObject {
	return v.value
}

func (v *NullableCreateIpv6sharednetworkResponseAsObject) Set(val *CreateIpv6sharednetworkResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIpv6sharednetworkResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIpv6sharednetworkResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIpv6sharednetworkResponseAsObject(val *CreateIpv6sharednetworkResponseAsObject) *NullableCreateIpv6sharednetworkResponseAsObject {
	return &NullableCreateIpv6sharednetworkResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateIpv6sharednetworkResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIpv6sharednetworkResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
