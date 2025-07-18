/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
	"fmt"
)

// UpdateIpv6sharednetworkResponse - struct for UpdateIpv6sharednetworkResponse
type UpdateIpv6sharednetworkResponse struct {
	UpdateIpv6sharednetworkResponseAsObject *UpdateIpv6sharednetworkResponseAsObject
	String                                  *string
}

// UpdateIpv6sharednetworkResponseAsObjectAsUpdateIpv6sharednetworkResponse is a convenience function that returns UpdateIpv6sharednetworkResponseAsObject wrapped in UpdateIpv6sharednetworkResponse
func UpdateIpv6sharednetworkResponseAsObjectAsUpdateIpv6sharednetworkResponse(v *UpdateIpv6sharednetworkResponseAsObject) UpdateIpv6sharednetworkResponse {
	return UpdateIpv6sharednetworkResponse{
		UpdateIpv6sharednetworkResponseAsObject: v,
	}
}

// stringAsUpdateIpv6sharednetworkResponse is a convenience function that returns string wrapped in UpdateIpv6sharednetworkResponse
func StringAsUpdateIpv6sharednetworkResponse(v *string) UpdateIpv6sharednetworkResponse {
	return UpdateIpv6sharednetworkResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateIpv6sharednetworkResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateIpv6sharednetworkResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateIpv6sharednetworkResponseAsObject)
	if err == nil {
		jsonUpdateIpv6sharednetworkResponseAsObject, _ := json.Marshal(dst.UpdateIpv6sharednetworkResponseAsObject)
		if string(jsonUpdateIpv6sharednetworkResponseAsObject) == "{}" { // empty struct
			dst.UpdateIpv6sharednetworkResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateIpv6sharednetworkResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateIpv6sharednetworkResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateIpv6sharednetworkResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateIpv6sharednetworkResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateIpv6sharednetworkResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateIpv6sharednetworkResponseAsObject != nil {
		return json.Marshal(&src.UpdateIpv6sharednetworkResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateIpv6sharednetworkResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateIpv6sharednetworkResponseAsObject != nil {
		return obj.UpdateIpv6sharednetworkResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateIpv6sharednetworkResponse struct {
	value *UpdateIpv6sharednetworkResponse
	isSet bool
}

func (v NullableUpdateIpv6sharednetworkResponse) Get() *UpdateIpv6sharednetworkResponse {
	return v.value
}

func (v *NullableUpdateIpv6sharednetworkResponse) Set(val *UpdateIpv6sharednetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpv6sharednetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpv6sharednetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpv6sharednetworkResponse(val *UpdateIpv6sharednetworkResponse) *NullableUpdateIpv6sharednetworkResponse {
	return &NullableUpdateIpv6sharednetworkResponse{value: val, isSet: true}
}

func (v NullableUpdateIpv6sharednetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpv6sharednetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
