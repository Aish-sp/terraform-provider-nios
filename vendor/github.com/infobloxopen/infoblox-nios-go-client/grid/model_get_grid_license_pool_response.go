/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
	"fmt"
)

// GetGridLicensePoolResponse - struct for GetGridLicensePoolResponse
type GetGridLicensePoolResponse struct {
	GetGridLicensePoolResponseObjectAsResult *GetGridLicensePoolResponseObjectAsResult
	GridLicensePool                          *GridLicensePool
}

// GetGridLicensePoolResponseObjectAsResultAsGetGridLicensePoolResponse is a convenience function that returns GetGridLicensePoolResponseObjectAsResult wrapped in GetGridLicensePoolResponse
func GetGridLicensePoolResponseObjectAsResultAsGetGridLicensePoolResponse(v *GetGridLicensePoolResponseObjectAsResult) GetGridLicensePoolResponse {
	return GetGridLicensePoolResponse{
		GetGridLicensePoolResponseObjectAsResult: v,
	}
}

// GridLicensePoolAsGetGridLicensePoolResponse is a convenience function that returns GridLicensePool wrapped in GetGridLicensePoolResponse
func GridLicensePoolAsGetGridLicensePoolResponse(v *GridLicensePool) GetGridLicensePoolResponse {
	return GetGridLicensePoolResponse{
		GridLicensePool: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGridLicensePoolResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetGridLicensePoolResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetGridLicensePoolResponseObjectAsResult)
	if err == nil {
		jsonGetGridLicensePoolResponseObjectAsResult, _ := json.Marshal(dst.GetGridLicensePoolResponseObjectAsResult)
		if string(jsonGetGridLicensePoolResponseObjectAsResult) == "{}" { // empty struct
			dst.GetGridLicensePoolResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetGridLicensePoolResponseObjectAsResult = nil
	}

	// try to unmarshal data into GridLicensePool
	err = newStrictDecoder(data).Decode(&dst.GridLicensePool)
	if err == nil {
		jsonGridLicensePool, _ := json.Marshal(dst.GridLicensePool)
		if string(jsonGridLicensePool) == "{}" { // empty struct
			dst.GridLicensePool = nil
		} else {
			match++
		}
	} else {
		dst.GridLicensePool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetGridLicensePoolResponseObjectAsResult = nil
		dst.GridLicensePool = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGridLicensePoolResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGridLicensePoolResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGridLicensePoolResponse) MarshalJSON() ([]byte, error) {
	if src.GetGridLicensePoolResponseObjectAsResult != nil {
		return json.Marshal(&src.GetGridLicensePoolResponseObjectAsResult)
	}

	if src.GridLicensePool != nil {
		return json.Marshal(&src.GridLicensePool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGridLicensePoolResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetGridLicensePoolResponseObjectAsResult != nil {
		return obj.GetGridLicensePoolResponseObjectAsResult
	}

	if obj.GridLicensePool != nil {
		return obj.GridLicensePool
	}

	// all schemas are nil
	return nil
}

type NullableGetGridLicensePoolResponse struct {
	value *GetGridLicensePoolResponse
	isSet bool
}

func (v NullableGetGridLicensePoolResponse) Get() *GetGridLicensePoolResponse {
	return v.value
}

func (v *NullableGetGridLicensePoolResponse) Set(val *GetGridLicensePoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridLicensePoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridLicensePoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridLicensePoolResponse(val *GetGridLicensePoolResponse) *NullableGetGridLicensePoolResponse {
	return &NullableGetGridLicensePoolResponse{value: val, isSet: true}
}

func (v NullableGetGridLicensePoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridLicensePoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
