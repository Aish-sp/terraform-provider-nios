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

// CreateRecordNsResponse - struct for CreateRecordNsResponse
type CreateRecordNsResponse struct {
	CreateRecordNsResponseAsObject *CreateRecordNsResponseAsObject
	String                         *string
}

// CreateRecordNsResponseAsObjectAsCreateRecordNsResponse is a convenience function that returns CreateRecordNsResponseAsObject wrapped in CreateRecordNsResponse
func CreateRecordNsResponseAsObjectAsCreateRecordNsResponse(v *CreateRecordNsResponseAsObject) CreateRecordNsResponse {
	return CreateRecordNsResponse{
		CreateRecordNsResponseAsObject: v,
	}
}

// stringAsCreateRecordNsResponse is a convenience function that returns string wrapped in CreateRecordNsResponse
func StringAsCreateRecordNsResponse(v *string) CreateRecordNsResponse {
	return CreateRecordNsResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateRecordNsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateRecordNsResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateRecordNsResponseAsObject)
	if err == nil {
		jsonCreateRecordNsResponseAsObject, _ := json.Marshal(dst.CreateRecordNsResponseAsObject)
		if string(jsonCreateRecordNsResponseAsObject) == "{}" { // empty struct
			dst.CreateRecordNsResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateRecordNsResponseAsObject = nil
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
		dst.CreateRecordNsResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateRecordNsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateRecordNsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateRecordNsResponse) MarshalJSON() ([]byte, error) {
	if src.CreateRecordNsResponseAsObject != nil {
		return json.Marshal(&src.CreateRecordNsResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateRecordNsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateRecordNsResponseAsObject != nil {
		return obj.CreateRecordNsResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateRecordNsResponse struct {
	value *CreateRecordNsResponse
	isSet bool
}

func (v NullableCreateRecordNsResponse) Get() *CreateRecordNsResponse {
	return v.value
}

func (v *NullableCreateRecordNsResponse) Set(val *CreateRecordNsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordNsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordNsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordNsResponse(val *CreateRecordNsResponse) *NullableCreateRecordNsResponse {
	return &NullableCreateRecordNsResponse{value: val, isSet: true}
}

func (v NullableCreateRecordNsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordNsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
