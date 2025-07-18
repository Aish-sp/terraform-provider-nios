/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
	"fmt"
)

// UpdateDbObjectsResponse - struct for UpdateDbObjectsResponse
type UpdateDbObjectsResponse struct {
	UpdateDbObjectsResponseAsObject *UpdateDbObjectsResponseAsObject
	String                          *string
}

// UpdateDbObjectsResponseAsObjectAsUpdateDbObjectsResponse is a convenience function that returns UpdateDbObjectsResponseAsObject wrapped in UpdateDbObjectsResponse
func UpdateDbObjectsResponseAsObjectAsUpdateDbObjectsResponse(v *UpdateDbObjectsResponseAsObject) UpdateDbObjectsResponse {
	return UpdateDbObjectsResponse{
		UpdateDbObjectsResponseAsObject: v,
	}
}

// stringAsUpdateDbObjectsResponse is a convenience function that returns string wrapped in UpdateDbObjectsResponse
func StringAsUpdateDbObjectsResponse(v *string) UpdateDbObjectsResponse {
	return UpdateDbObjectsResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateDbObjectsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateDbObjectsResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateDbObjectsResponseAsObject)
	if err == nil {
		jsonUpdateDbObjectsResponseAsObject, _ := json.Marshal(dst.UpdateDbObjectsResponseAsObject)
		if string(jsonUpdateDbObjectsResponseAsObject) == "{}" { // empty struct
			dst.UpdateDbObjectsResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateDbObjectsResponseAsObject = nil
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
		dst.UpdateDbObjectsResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateDbObjectsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateDbObjectsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDbObjectsResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateDbObjectsResponseAsObject != nil {
		return json.Marshal(&src.UpdateDbObjectsResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDbObjectsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateDbObjectsResponseAsObject != nil {
		return obj.UpdateDbObjectsResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDbObjectsResponse struct {
	value *UpdateDbObjectsResponse
	isSet bool
}

func (v NullableUpdateDbObjectsResponse) Get() *UpdateDbObjectsResponse {
	return v.value
}

func (v *NullableUpdateDbObjectsResponse) Set(val *UpdateDbObjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDbObjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDbObjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDbObjectsResponse(val *UpdateDbObjectsResponse) *NullableUpdateDbObjectsResponse {
	return &NullableUpdateDbObjectsResponse{value: val, isSet: true}
}

func (v NullableUpdateDbObjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDbObjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
