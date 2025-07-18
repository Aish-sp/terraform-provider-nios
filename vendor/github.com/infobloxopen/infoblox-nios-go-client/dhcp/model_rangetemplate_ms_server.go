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

// checks if the RangetemplateMsServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangetemplateMsServer{}

// RangetemplateMsServer struct for RangetemplateMsServer
type RangetemplateMsServer struct {
	// The IPv4 Address or FQDN of the Microsoft server.
	Ipv4addr             *string `json:"ipv4addr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RangetemplateMsServer RangetemplateMsServer

// NewRangetemplateMsServer instantiates a new RangetemplateMsServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangetemplateMsServer() *RangetemplateMsServer {
	this := RangetemplateMsServer{}
	return &this
}

// NewRangetemplateMsServerWithDefaults instantiates a new RangetemplateMsServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangetemplateMsServerWithDefaults() *RangetemplateMsServer {
	this := RangetemplateMsServer{}
	return &this
}

// GetIpv4addr returns the Ipv4addr field value if set, zero value otherwise.
func (o *RangetemplateMsServer) GetIpv4addr() string {
	if o == nil || IsNil(o.Ipv4addr) {
		var ret string
		return ret
	}
	return *o.Ipv4addr
}

// GetIpv4addrOk returns a tuple with the Ipv4addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangetemplateMsServer) GetIpv4addrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4addr) {
		return nil, false
	}
	return o.Ipv4addr, true
}

// HasIpv4addr returns a boolean if a field has been set.
func (o *RangetemplateMsServer) HasIpv4addr() bool {
	if o != nil && !IsNil(o.Ipv4addr) {
		return true
	}

	return false
}

// SetIpv4addr gets a reference to the given string and assigns it to the Ipv4addr field.
func (o *RangetemplateMsServer) SetIpv4addr(v string) {
	o.Ipv4addr = &v
}

func (o RangetemplateMsServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangetemplateMsServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4addr) {
		toSerialize["ipv4addr"] = o.Ipv4addr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RangetemplateMsServer) UnmarshalJSON(data []byte) (err error) {
	varRangetemplateMsServer := _RangetemplateMsServer{}

	err = json.Unmarshal(data, &varRangetemplateMsServer)

	if err != nil {
		return err
	}

	*o = RangetemplateMsServer(varRangetemplateMsServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ipv4addr")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRangetemplateMsServer struct {
	value *RangetemplateMsServer
	isSet bool
}

func (v NullableRangetemplateMsServer) Get() *RangetemplateMsServer {
	return v.value
}

func (v *NullableRangetemplateMsServer) Set(val *RangetemplateMsServer) {
	v.value = val
	v.isSet = true
}

func (v NullableRangetemplateMsServer) IsSet() bool {
	return v.isSet
}

func (v *NullableRangetemplateMsServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangetemplateMsServer(val *RangetemplateMsServer) *NullableRangetemplateMsServer {
	return &NullableRangetemplateMsServer{value: val, isSet: true}
}

func (v NullableRangetemplateMsServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangetemplateMsServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
