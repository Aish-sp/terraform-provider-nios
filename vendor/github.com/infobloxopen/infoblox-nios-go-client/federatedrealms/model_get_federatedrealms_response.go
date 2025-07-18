/*
Infoblox FEDERATEDREALMS API

OpenAPI specification for Infoblox NIOS WAPI FEDERATEDREALMS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package federatedrealms

import (
	"encoding/json"
	"fmt"
)

// GetFederatedrealmsResponse - struct for GetFederatedrealmsResponse
type GetFederatedrealmsResponse struct {
	Federatedrealms                          *Federatedrealms
	GetFederatedrealmsResponseObjectAsResult *GetFederatedrealmsResponseObjectAsResult
}

// FederatedrealmsAsGetFederatedrealmsResponse is a convenience function that returns Federatedrealms wrapped in GetFederatedrealmsResponse
func FederatedrealmsAsGetFederatedrealmsResponse(v *Federatedrealms) GetFederatedrealmsResponse {
	return GetFederatedrealmsResponse{
		Federatedrealms: v,
	}
}

// GetFederatedrealmsResponseObjectAsResultAsGetFederatedrealmsResponse is a convenience function that returns GetFederatedrealmsResponseObjectAsResult wrapped in GetFederatedrealmsResponse
func GetFederatedrealmsResponseObjectAsResultAsGetFederatedrealmsResponse(v *GetFederatedrealmsResponseObjectAsResult) GetFederatedrealmsResponse {
	return GetFederatedrealmsResponse{
		GetFederatedrealmsResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFederatedrealmsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Federatedrealms
	err = newStrictDecoder(data).Decode(&dst.Federatedrealms)
	if err == nil {
		jsonFederatedrealms, _ := json.Marshal(dst.Federatedrealms)
		if string(jsonFederatedrealms) == "{}" { // empty struct
			dst.Federatedrealms = nil
		} else {
			match++
		}
	} else {
		dst.Federatedrealms = nil
	}

	// try to unmarshal data into GetFederatedrealmsResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetFederatedrealmsResponseObjectAsResult)
	if err == nil {
		jsonGetFederatedrealmsResponseObjectAsResult, _ := json.Marshal(dst.GetFederatedrealmsResponseObjectAsResult)
		if string(jsonGetFederatedrealmsResponseObjectAsResult) == "{}" { // empty struct
			dst.GetFederatedrealmsResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetFederatedrealmsResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Federatedrealms = nil
		dst.GetFederatedrealmsResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetFederatedrealmsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetFederatedrealmsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFederatedrealmsResponse) MarshalJSON() ([]byte, error) {
	if src.Federatedrealms != nil {
		return json.Marshal(&src.Federatedrealms)
	}

	if src.GetFederatedrealmsResponseObjectAsResult != nil {
		return json.Marshal(&src.GetFederatedrealmsResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetFederatedrealmsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Federatedrealms != nil {
		return obj.Federatedrealms
	}

	if obj.GetFederatedrealmsResponseObjectAsResult != nil {
		return obj.GetFederatedrealmsResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetFederatedrealmsResponse struct {
	value *GetFederatedrealmsResponse
	isSet bool
}

func (v NullableGetFederatedrealmsResponse) Get() *GetFederatedrealmsResponse {
	return v.value
}

func (v *NullableGetFederatedrealmsResponse) Set(val *GetFederatedrealmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFederatedrealmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFederatedrealmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFederatedrealmsResponse(val *GetFederatedrealmsResponse) *NullableGetFederatedrealmsResponse {
	return &NullableGetFederatedrealmsResponse{value: val, isSet: true}
}

func (v NullableGetFederatedrealmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFederatedrealmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
