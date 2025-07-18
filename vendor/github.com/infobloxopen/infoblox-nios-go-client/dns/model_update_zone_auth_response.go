/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// UpdateZoneAuthResponse - struct for UpdateZoneAuthResponse
type UpdateZoneAuthResponse struct {
	UpdateZoneAuthResponseAsObject *UpdateZoneAuthResponseAsObject
	String                         *string
}

// UpdateZoneAuthResponseAsObjectAsUpdateZoneAuthResponse is a convenience function that returns UpdateZoneAuthResponseAsObject wrapped in UpdateZoneAuthResponse
func UpdateZoneAuthResponseAsObjectAsUpdateZoneAuthResponse(v *UpdateZoneAuthResponseAsObject) UpdateZoneAuthResponse {
	return UpdateZoneAuthResponse{
		UpdateZoneAuthResponseAsObject: v,
	}
}

// stringAsUpdateZoneAuthResponse is a convenience function that returns string wrapped in UpdateZoneAuthResponse
func StringAsUpdateZoneAuthResponse(v *string) UpdateZoneAuthResponse {
	return UpdateZoneAuthResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateZoneAuthResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateZoneAuthResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateZoneAuthResponseAsObject)
	if err == nil {
		jsonUpdateZoneAuthResponseAsObject, _ := json.Marshal(dst.UpdateZoneAuthResponseAsObject)
		if string(jsonUpdateZoneAuthResponseAsObject) == "{}" { // empty struct
			dst.UpdateZoneAuthResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateZoneAuthResponseAsObject = nil
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
		dst.UpdateZoneAuthResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateZoneAuthResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateZoneAuthResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateZoneAuthResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateZoneAuthResponseAsObject != nil {
		return json.Marshal(&src.UpdateZoneAuthResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateZoneAuthResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateZoneAuthResponseAsObject != nil {
		return obj.UpdateZoneAuthResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateZoneAuthResponse struct {
	value *UpdateZoneAuthResponse
	isSet bool
}

func (v NullableUpdateZoneAuthResponse) Get() *UpdateZoneAuthResponse {
	return v.value
}

func (v *NullableUpdateZoneAuthResponse) Set(val *UpdateZoneAuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateZoneAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateZoneAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateZoneAuthResponse(val *UpdateZoneAuthResponse) *NullableUpdateZoneAuthResponse {
	return &NullableUpdateZoneAuthResponse{value: val, isSet: true}
}

func (v NullableUpdateZoneAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateZoneAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
