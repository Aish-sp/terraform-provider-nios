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

// UpdatePxgridEndpointResponse - struct for UpdatePxgridEndpointResponse
type UpdatePxgridEndpointResponse struct {
	UpdatePxgridEndpointResponseAsObject *UpdatePxgridEndpointResponseAsObject
	String                               *string
}

// UpdatePxgridEndpointResponseAsObjectAsUpdatePxgridEndpointResponse is a convenience function that returns UpdatePxgridEndpointResponseAsObject wrapped in UpdatePxgridEndpointResponse
func UpdatePxgridEndpointResponseAsObjectAsUpdatePxgridEndpointResponse(v *UpdatePxgridEndpointResponseAsObject) UpdatePxgridEndpointResponse {
	return UpdatePxgridEndpointResponse{
		UpdatePxgridEndpointResponseAsObject: v,
	}
}

// stringAsUpdatePxgridEndpointResponse is a convenience function that returns string wrapped in UpdatePxgridEndpointResponse
func StringAsUpdatePxgridEndpointResponse(v *string) UpdatePxgridEndpointResponse {
	return UpdatePxgridEndpointResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdatePxgridEndpointResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdatePxgridEndpointResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdatePxgridEndpointResponseAsObject)
	if err == nil {
		jsonUpdatePxgridEndpointResponseAsObject, _ := json.Marshal(dst.UpdatePxgridEndpointResponseAsObject)
		if string(jsonUpdatePxgridEndpointResponseAsObject) == "{}" { // empty struct
			dst.UpdatePxgridEndpointResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdatePxgridEndpointResponseAsObject = nil
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
		dst.UpdatePxgridEndpointResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdatePxgridEndpointResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdatePxgridEndpointResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdatePxgridEndpointResponse) MarshalJSON() ([]byte, error) {
	if src.UpdatePxgridEndpointResponseAsObject != nil {
		return json.Marshal(&src.UpdatePxgridEndpointResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdatePxgridEndpointResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdatePxgridEndpointResponseAsObject != nil {
		return obj.UpdatePxgridEndpointResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdatePxgridEndpointResponse struct {
	value *UpdatePxgridEndpointResponse
	isSet bool
}

func (v NullableUpdatePxgridEndpointResponse) Get() *UpdatePxgridEndpointResponse {
	return v.value
}

func (v *NullableUpdatePxgridEndpointResponse) Set(val *UpdatePxgridEndpointResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePxgridEndpointResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePxgridEndpointResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePxgridEndpointResponse(val *UpdatePxgridEndpointResponse) *NullableUpdatePxgridEndpointResponse {
	return &NullableUpdatePxgridEndpointResponse{value: val, isSet: true}
}

func (v NullableUpdatePxgridEndpointResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePxgridEndpointResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
