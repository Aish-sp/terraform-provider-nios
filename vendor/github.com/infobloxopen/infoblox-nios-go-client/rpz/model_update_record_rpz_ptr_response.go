/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
	"fmt"
)

// UpdateRecordRpzPtrResponse - struct for UpdateRecordRpzPtrResponse
type UpdateRecordRpzPtrResponse struct {
	UpdateRecordRpzPtrResponseAsObject *UpdateRecordRpzPtrResponseAsObject
	String                             *string
}

// UpdateRecordRpzPtrResponseAsObjectAsUpdateRecordRpzPtrResponse is a convenience function that returns UpdateRecordRpzPtrResponseAsObject wrapped in UpdateRecordRpzPtrResponse
func UpdateRecordRpzPtrResponseAsObjectAsUpdateRecordRpzPtrResponse(v *UpdateRecordRpzPtrResponseAsObject) UpdateRecordRpzPtrResponse {
	return UpdateRecordRpzPtrResponse{
		UpdateRecordRpzPtrResponseAsObject: v,
	}
}

// stringAsUpdateRecordRpzPtrResponse is a convenience function that returns string wrapped in UpdateRecordRpzPtrResponse
func StringAsUpdateRecordRpzPtrResponse(v *string) UpdateRecordRpzPtrResponse {
	return UpdateRecordRpzPtrResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRecordRpzPtrResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateRecordRpzPtrResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateRecordRpzPtrResponseAsObject)
	if err == nil {
		jsonUpdateRecordRpzPtrResponseAsObject, _ := json.Marshal(dst.UpdateRecordRpzPtrResponseAsObject)
		if string(jsonUpdateRecordRpzPtrResponseAsObject) == "{}" { // empty struct
			dst.UpdateRecordRpzPtrResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRecordRpzPtrResponseAsObject = nil
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
		dst.UpdateRecordRpzPtrResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateRecordRpzPtrResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateRecordRpzPtrResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRecordRpzPtrResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateRecordRpzPtrResponseAsObject != nil {
		return json.Marshal(&src.UpdateRecordRpzPtrResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRecordRpzPtrResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateRecordRpzPtrResponseAsObject != nil {
		return obj.UpdateRecordRpzPtrResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRecordRpzPtrResponse struct {
	value *UpdateRecordRpzPtrResponse
	isSet bool
}

func (v NullableUpdateRecordRpzPtrResponse) Get() *UpdateRecordRpzPtrResponse {
	return v.value
}

func (v *NullableUpdateRecordRpzPtrResponse) Set(val *UpdateRecordRpzPtrResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordRpzPtrResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordRpzPtrResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordRpzPtrResponse(val *UpdateRecordRpzPtrResponse) *NullableUpdateRecordRpzPtrResponse {
	return &NullableUpdateRecordRpzPtrResponse{value: val, isSet: true}
}

func (v NullableUpdateRecordRpzPtrResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordRpzPtrResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
