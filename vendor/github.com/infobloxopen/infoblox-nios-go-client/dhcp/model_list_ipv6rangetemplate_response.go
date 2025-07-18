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

// ListIpv6rangetemplateResponse - struct for ListIpv6rangetemplateResponse
type ListIpv6rangetemplateResponse struct {
	ListIpv6rangetemplateResponseObject *ListIpv6rangetemplateResponseObject
	ArrayOfIpv6rangetemplate            *[]Ipv6rangetemplate
}

// ListIpv6rangetemplateResponseObjectAsListIpv6rangetemplateResponse is a convenience function that returns ListIpv6rangetemplateResponseObject wrapped in ListIpv6rangetemplateResponse
func ListIpv6rangetemplateResponseObjectAsListIpv6rangetemplateResponse(v *ListIpv6rangetemplateResponseObject) ListIpv6rangetemplateResponse {
	return ListIpv6rangetemplateResponse{
		ListIpv6rangetemplateResponseObject: v,
	}
}

// []Ipv6rangetemplateAsListIpv6rangetemplateResponse is a convenience function that returns []Ipv6rangetemplate wrapped in ListIpv6rangetemplateResponse
func ArrayOfIpv6rangetemplateAsListIpv6rangetemplateResponse(v *[]Ipv6rangetemplate) ListIpv6rangetemplateResponse {
	return ListIpv6rangetemplateResponse{
		ArrayOfIpv6rangetemplate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListIpv6rangetemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListIpv6rangetemplateResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListIpv6rangetemplateResponseObject)
	if err == nil {
		jsonListIpv6rangetemplateResponseObject, _ := json.Marshal(dst.ListIpv6rangetemplateResponseObject)
		if string(jsonListIpv6rangetemplateResponseObject) == "{}" { // empty struct
			dst.ListIpv6rangetemplateResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListIpv6rangetemplateResponseObject = nil
	}

	// try to unmarshal data into ArrayOfIpv6rangetemplate
	err = newStrictDecoder(data).Decode(&dst.ArrayOfIpv6rangetemplate)
	if err == nil {
		jsonArrayOfIpv6rangetemplate, _ := json.Marshal(dst.ArrayOfIpv6rangetemplate)
		if string(jsonArrayOfIpv6rangetemplate) == "{}" { // empty struct
			dst.ArrayOfIpv6rangetemplate = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfIpv6rangetemplate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListIpv6rangetemplateResponseObject = nil
		dst.ArrayOfIpv6rangetemplate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListIpv6rangetemplateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListIpv6rangetemplateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListIpv6rangetemplateResponse) MarshalJSON() ([]byte, error) {
	if src.ListIpv6rangetemplateResponseObject != nil {
		return json.Marshal(&src.ListIpv6rangetemplateResponseObject)
	}

	if src.ArrayOfIpv6rangetemplate != nil {
		return json.Marshal(&src.ArrayOfIpv6rangetemplate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListIpv6rangetemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListIpv6rangetemplateResponseObject != nil {
		return obj.ListIpv6rangetemplateResponseObject
	}

	if obj.ArrayOfIpv6rangetemplate != nil {
		return obj.ArrayOfIpv6rangetemplate
	}

	// all schemas are nil
	return nil
}

type NullableListIpv6rangetemplateResponse struct {
	value *ListIpv6rangetemplateResponse
	isSet bool
}

func (v NullableListIpv6rangetemplateResponse) Get() *ListIpv6rangetemplateResponse {
	return v.value
}

func (v *NullableListIpv6rangetemplateResponse) Set(val *ListIpv6rangetemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListIpv6rangetemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListIpv6rangetemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIpv6rangetemplateResponse(val *ListIpv6rangetemplateResponse) *NullableListIpv6rangetemplateResponse {
	return &NullableListIpv6rangetemplateResponse{value: val, isSet: true}
}

func (v NullableListIpv6rangetemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIpv6rangetemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
