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

// CreateRoaminghostResponse - struct for CreateRoaminghostResponse
type CreateRoaminghostResponse struct {
	CreateRoaminghostResponseAsObject *CreateRoaminghostResponseAsObject
	String                            *string
}

// CreateRoaminghostResponseAsObjectAsCreateRoaminghostResponse is a convenience function that returns CreateRoaminghostResponseAsObject wrapped in CreateRoaminghostResponse
func CreateRoaminghostResponseAsObjectAsCreateRoaminghostResponse(v *CreateRoaminghostResponseAsObject) CreateRoaminghostResponse {
	return CreateRoaminghostResponse{
		CreateRoaminghostResponseAsObject: v,
	}
}

// stringAsCreateRoaminghostResponse is a convenience function that returns string wrapped in CreateRoaminghostResponse
func StringAsCreateRoaminghostResponse(v *string) CreateRoaminghostResponse {
	return CreateRoaminghostResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateRoaminghostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateRoaminghostResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateRoaminghostResponseAsObject)
	if err == nil {
		jsonCreateRoaminghostResponseAsObject, _ := json.Marshal(dst.CreateRoaminghostResponseAsObject)
		if string(jsonCreateRoaminghostResponseAsObject) == "{}" { // empty struct
			dst.CreateRoaminghostResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateRoaminghostResponseAsObject = nil
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
		dst.CreateRoaminghostResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateRoaminghostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateRoaminghostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateRoaminghostResponse) MarshalJSON() ([]byte, error) {
	if src.CreateRoaminghostResponseAsObject != nil {
		return json.Marshal(&src.CreateRoaminghostResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateRoaminghostResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateRoaminghostResponseAsObject != nil {
		return obj.CreateRoaminghostResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateRoaminghostResponse struct {
	value *CreateRoaminghostResponse
	isSet bool
}

func (v NullableCreateRoaminghostResponse) Get() *CreateRoaminghostResponse {
	return v.value
}

func (v *NullableCreateRoaminghostResponse) Set(val *CreateRoaminghostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoaminghostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoaminghostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoaminghostResponse(val *CreateRoaminghostResponse) *NullableCreateRoaminghostResponse {
	return &NullableCreateRoaminghostResponse{value: val, isSet: true}
}

func (v NullableCreateRoaminghostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoaminghostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
