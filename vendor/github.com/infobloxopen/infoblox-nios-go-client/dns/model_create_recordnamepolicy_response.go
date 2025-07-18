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

// CreateRecordnamepolicyResponse - struct for CreateRecordnamepolicyResponse
type CreateRecordnamepolicyResponse struct {
	CreateRecordnamepolicyResponseAsObject *CreateRecordnamepolicyResponseAsObject
	String                                 *string
}

// CreateRecordnamepolicyResponseAsObjectAsCreateRecordnamepolicyResponse is a convenience function that returns CreateRecordnamepolicyResponseAsObject wrapped in CreateRecordnamepolicyResponse
func CreateRecordnamepolicyResponseAsObjectAsCreateRecordnamepolicyResponse(v *CreateRecordnamepolicyResponseAsObject) CreateRecordnamepolicyResponse {
	return CreateRecordnamepolicyResponse{
		CreateRecordnamepolicyResponseAsObject: v,
	}
}

// stringAsCreateRecordnamepolicyResponse is a convenience function that returns string wrapped in CreateRecordnamepolicyResponse
func StringAsCreateRecordnamepolicyResponse(v *string) CreateRecordnamepolicyResponse {
	return CreateRecordnamepolicyResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateRecordnamepolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateRecordnamepolicyResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateRecordnamepolicyResponseAsObject)
	if err == nil {
		jsonCreateRecordnamepolicyResponseAsObject, _ := json.Marshal(dst.CreateRecordnamepolicyResponseAsObject)
		if string(jsonCreateRecordnamepolicyResponseAsObject) == "{}" { // empty struct
			dst.CreateRecordnamepolicyResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateRecordnamepolicyResponseAsObject = nil
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
		dst.CreateRecordnamepolicyResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateRecordnamepolicyResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateRecordnamepolicyResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateRecordnamepolicyResponse) MarshalJSON() ([]byte, error) {
	if src.CreateRecordnamepolicyResponseAsObject != nil {
		return json.Marshal(&src.CreateRecordnamepolicyResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateRecordnamepolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateRecordnamepolicyResponseAsObject != nil {
		return obj.CreateRecordnamepolicyResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateRecordnamepolicyResponse struct {
	value *CreateRecordnamepolicyResponse
	isSet bool
}

func (v NullableCreateRecordnamepolicyResponse) Get() *CreateRecordnamepolicyResponse {
	return v.value
}

func (v *NullableCreateRecordnamepolicyResponse) Set(val *CreateRecordnamepolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordnamepolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordnamepolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordnamepolicyResponse(val *CreateRecordnamepolicyResponse) *NullableCreateRecordnamepolicyResponse {
	return &NullableCreateRecordnamepolicyResponse{value: val, isSet: true}
}

func (v NullableCreateRecordnamepolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordnamepolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
