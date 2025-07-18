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

// checks if the RecordNsAddresses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordNsAddresses{}

// RecordNsAddresses struct for RecordNsAddresses
type RecordNsAddresses struct {
	// The address of the Zone Name Server.
	Address *string `json:"address,omitempty"`
	// Flag to indicate if ptr records need to be auto created.
	AutoCreatePtr        *bool `json:"auto_create_ptr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecordNsAddresses RecordNsAddresses

// NewRecordNsAddresses instantiates a new RecordNsAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordNsAddresses() *RecordNsAddresses {
	this := RecordNsAddresses{}
	return &this
}

// NewRecordNsAddressesWithDefaults instantiates a new RecordNsAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordNsAddressesWithDefaults() *RecordNsAddresses {
	this := RecordNsAddresses{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RecordNsAddresses) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordNsAddresses) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RecordNsAddresses) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RecordNsAddresses) SetAddress(v string) {
	o.Address = &v
}

// GetAutoCreatePtr returns the AutoCreatePtr field value if set, zero value otherwise.
func (o *RecordNsAddresses) GetAutoCreatePtr() bool {
	if o == nil || IsNil(o.AutoCreatePtr) {
		var ret bool
		return ret
	}
	return *o.AutoCreatePtr
}

// GetAutoCreatePtrOk returns a tuple with the AutoCreatePtr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordNsAddresses) GetAutoCreatePtrOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreatePtr) {
		return nil, false
	}
	return o.AutoCreatePtr, true
}

// HasAutoCreatePtr returns a boolean if a field has been set.
func (o *RecordNsAddresses) HasAutoCreatePtr() bool {
	if o != nil && !IsNil(o.AutoCreatePtr) {
		return true
	}

	return false
}

// SetAutoCreatePtr gets a reference to the given bool and assigns it to the AutoCreatePtr field.
func (o *RecordNsAddresses) SetAutoCreatePtr(v bool) {
	o.AutoCreatePtr = &v
}

func (o RecordNsAddresses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordNsAddresses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AutoCreatePtr) {
		toSerialize["auto_create_ptr"] = o.AutoCreatePtr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecordNsAddresses) UnmarshalJSON(data []byte) (err error) {
	varRecordNsAddresses := _RecordNsAddresses{}

	err = json.Unmarshal(data, &varRecordNsAddresses)

	if err != nil {
		return err
	}

	*o = RecordNsAddresses(varRecordNsAddresses)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "auto_create_ptr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecordNsAddresses struct {
	value *RecordNsAddresses
	isSet bool
}

func (v NullableRecordNsAddresses) Get() *RecordNsAddresses {
	return v.value
}

func (v *NullableRecordNsAddresses) Set(val *RecordNsAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordNsAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordNsAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordNsAddresses(val *RecordNsAddresses) *NullableRecordNsAddresses {
	return &NullableRecordNsAddresses{value: val, isSet: true}
}

func (v NullableRecordNsAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordNsAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
