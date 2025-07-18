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

// ListRecordNsec3Response - struct for ListRecordNsec3Response
type ListRecordNsec3Response struct {
	ListRecordNsec3ResponseObject *ListRecordNsec3ResponseObject
	ArrayOfRecordNsec3            *[]RecordNsec3
}

// ListRecordNsec3ResponseObjectAsListRecordNsec3Response is a convenience function that returns ListRecordNsec3ResponseObject wrapped in ListRecordNsec3Response
func ListRecordNsec3ResponseObjectAsListRecordNsec3Response(v *ListRecordNsec3ResponseObject) ListRecordNsec3Response {
	return ListRecordNsec3Response{
		ListRecordNsec3ResponseObject: v,
	}
}

// []RecordNsec3AsListRecordNsec3Response is a convenience function that returns []RecordNsec3 wrapped in ListRecordNsec3Response
func ArrayOfRecordNsec3AsListRecordNsec3Response(v *[]RecordNsec3) ListRecordNsec3Response {
	return ListRecordNsec3Response{
		ArrayOfRecordNsec3: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRecordNsec3Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRecordNsec3ResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRecordNsec3ResponseObject)
	if err == nil {
		jsonListRecordNsec3ResponseObject, _ := json.Marshal(dst.ListRecordNsec3ResponseObject)
		if string(jsonListRecordNsec3ResponseObject) == "{}" { // empty struct
			dst.ListRecordNsec3ResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRecordNsec3ResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRecordNsec3
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRecordNsec3)
	if err == nil {
		jsonArrayOfRecordNsec3, _ := json.Marshal(dst.ArrayOfRecordNsec3)
		if string(jsonArrayOfRecordNsec3) == "{}" { // empty struct
			dst.ArrayOfRecordNsec3 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRecordNsec3 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRecordNsec3ResponseObject = nil
		dst.ArrayOfRecordNsec3 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRecordNsec3Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRecordNsec3Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRecordNsec3Response) MarshalJSON() ([]byte, error) {
	if src.ListRecordNsec3ResponseObject != nil {
		return json.Marshal(&src.ListRecordNsec3ResponseObject)
	}

	if src.ArrayOfRecordNsec3 != nil {
		return json.Marshal(&src.ArrayOfRecordNsec3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRecordNsec3Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRecordNsec3ResponseObject != nil {
		return obj.ListRecordNsec3ResponseObject
	}

	if obj.ArrayOfRecordNsec3 != nil {
		return obj.ArrayOfRecordNsec3
	}

	// all schemas are nil
	return nil
}

type NullableListRecordNsec3Response struct {
	value *ListRecordNsec3Response
	isSet bool
}

func (v NullableListRecordNsec3Response) Get() *ListRecordNsec3Response {
	return v.value
}

func (v *NullableListRecordNsec3Response) Set(val *ListRecordNsec3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordNsec3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordNsec3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordNsec3Response(val *ListRecordNsec3Response) *NullableListRecordNsec3Response {
	return &NullableListRecordNsec3Response{value: val, isSet: true}
}

func (v NullableListRecordNsec3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordNsec3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
