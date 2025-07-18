/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamStatistics{}

// IpamStatistics struct for IpamStatistics
type IpamStatistics struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The network CIDR.
	Cidr *int64 `json:"cidr,omitempty"`
	// The number of conflicts discovered via network discovery. This attribute is only valid for a Network object.
	ConflictCount *int64                      `json:"conflict_count,omitempty"`
	MsAdUserData  *IpamStatisticsMsAdUserData `json:"ms_ad_user_data,omitempty"`
	// The network address.
	Network *string `json:"network,omitempty"`
	// The network view.
	NetworkView *string `json:"network_view,omitempty"`
	// The number of unmanaged IP addresses as discovered by network discovery. This attribute is only valid for a Network object.
	UnmanagedCount *int64 `json:"unmanaged_count,omitempty"`
	// The network utilization in percentage.
	Utilization *int64 `json:"utilization,omitempty"`
	// The time that the utilization statistics were updated last. This attribute is only valid for a Network object. For a Network Container object, the return value is undefined.
	UtilizationUpdate *int64 `json:"utilization_update,omitempty"`
}

// NewIpamStatistics instantiates a new IpamStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamStatistics() *IpamStatistics {
	this := IpamStatistics{}
	return &this
}

// NewIpamStatisticsWithDefaults instantiates a new IpamStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamStatisticsWithDefaults() *IpamStatistics {
	this := IpamStatistics{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *IpamStatistics) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *IpamStatistics) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *IpamStatistics) SetRef(v string) {
	o.Ref = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *IpamStatistics) GetCidr() int64 {
	if o == nil || IsNil(o.Cidr) {
		var ret int64
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *IpamStatistics) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int64 and assigns it to the Cidr field.
func (o *IpamStatistics) SetCidr(v int64) {
	o.Cidr = &v
}

// GetConflictCount returns the ConflictCount field value if set, zero value otherwise.
func (o *IpamStatistics) GetConflictCount() int64 {
	if o == nil || IsNil(o.ConflictCount) {
		var ret int64
		return ret
	}
	return *o.ConflictCount
}

// GetConflictCountOk returns a tuple with the ConflictCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetConflictCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ConflictCount) {
		return nil, false
	}
	return o.ConflictCount, true
}

// HasConflictCount returns a boolean if a field has been set.
func (o *IpamStatistics) HasConflictCount() bool {
	if o != nil && !IsNil(o.ConflictCount) {
		return true
	}

	return false
}

// SetConflictCount gets a reference to the given int64 and assigns it to the ConflictCount field.
func (o *IpamStatistics) SetConflictCount(v int64) {
	o.ConflictCount = &v
}

// GetMsAdUserData returns the MsAdUserData field value if set, zero value otherwise.
func (o *IpamStatistics) GetMsAdUserData() IpamStatisticsMsAdUserData {
	if o == nil || IsNil(o.MsAdUserData) {
		var ret IpamStatisticsMsAdUserData
		return ret
	}
	return *o.MsAdUserData
}

// GetMsAdUserDataOk returns a tuple with the MsAdUserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetMsAdUserDataOk() (*IpamStatisticsMsAdUserData, bool) {
	if o == nil || IsNil(o.MsAdUserData) {
		return nil, false
	}
	return o.MsAdUserData, true
}

// HasMsAdUserData returns a boolean if a field has been set.
func (o *IpamStatistics) HasMsAdUserData() bool {
	if o != nil && !IsNil(o.MsAdUserData) {
		return true
	}

	return false
}

// SetMsAdUserData gets a reference to the given IpamStatisticsMsAdUserData and assigns it to the MsAdUserData field.
func (o *IpamStatistics) SetMsAdUserData(v IpamStatisticsMsAdUserData) {
	o.MsAdUserData = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *IpamStatistics) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *IpamStatistics) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *IpamStatistics) SetNetwork(v string) {
	o.Network = &v
}

// GetNetworkView returns the NetworkView field value if set, zero value otherwise.
func (o *IpamStatistics) GetNetworkView() string {
	if o == nil || IsNil(o.NetworkView) {
		var ret string
		return ret
	}
	return *o.NetworkView
}

// GetNetworkViewOk returns a tuple with the NetworkView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetNetworkViewOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkView) {
		return nil, false
	}
	return o.NetworkView, true
}

// HasNetworkView returns a boolean if a field has been set.
func (o *IpamStatistics) HasNetworkView() bool {
	if o != nil && !IsNil(o.NetworkView) {
		return true
	}

	return false
}

// SetNetworkView gets a reference to the given string and assigns it to the NetworkView field.
func (o *IpamStatistics) SetNetworkView(v string) {
	o.NetworkView = &v
}

// GetUnmanagedCount returns the UnmanagedCount field value if set, zero value otherwise.
func (o *IpamStatistics) GetUnmanagedCount() int64 {
	if o == nil || IsNil(o.UnmanagedCount) {
		var ret int64
		return ret
	}
	return *o.UnmanagedCount
}

// GetUnmanagedCountOk returns a tuple with the UnmanagedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetUnmanagedCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UnmanagedCount) {
		return nil, false
	}
	return o.UnmanagedCount, true
}

// HasUnmanagedCount returns a boolean if a field has been set.
func (o *IpamStatistics) HasUnmanagedCount() bool {
	if o != nil && !IsNil(o.UnmanagedCount) {
		return true
	}

	return false
}

// SetUnmanagedCount gets a reference to the given int64 and assigns it to the UnmanagedCount field.
func (o *IpamStatistics) SetUnmanagedCount(v int64) {
	o.UnmanagedCount = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *IpamStatistics) GetUtilization() int64 {
	if o == nil || IsNil(o.Utilization) {
		var ret int64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetUtilizationOk() (*int64, bool) {
	if o == nil || IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *IpamStatistics) HasUtilization() bool {
	if o != nil && !IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given int64 and assigns it to the Utilization field.
func (o *IpamStatistics) SetUtilization(v int64) {
	o.Utilization = &v
}

// GetUtilizationUpdate returns the UtilizationUpdate field value if set, zero value otherwise.
func (o *IpamStatistics) GetUtilizationUpdate() int64 {
	if o == nil || IsNil(o.UtilizationUpdate) {
		var ret int64
		return ret
	}
	return *o.UtilizationUpdate
}

// GetUtilizationUpdateOk returns a tuple with the UtilizationUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamStatistics) GetUtilizationUpdateOk() (*int64, bool) {
	if o == nil || IsNil(o.UtilizationUpdate) {
		return nil, false
	}
	return o.UtilizationUpdate, true
}

// HasUtilizationUpdate returns a boolean if a field has been set.
func (o *IpamStatistics) HasUtilizationUpdate() bool {
	if o != nil && !IsNil(o.UtilizationUpdate) {
		return true
	}

	return false
}

// SetUtilizationUpdate gets a reference to the given int64 and assigns it to the UtilizationUpdate field.
func (o *IpamStatistics) SetUtilizationUpdate(v int64) {
	o.UtilizationUpdate = &v
}

func (o IpamStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.ConflictCount) {
		toSerialize["conflict_count"] = o.ConflictCount
	}
	if !IsNil(o.MsAdUserData) {
		toSerialize["ms_ad_user_data"] = o.MsAdUserData
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.NetworkView) {
		toSerialize["network_view"] = o.NetworkView
	}
	if !IsNil(o.UnmanagedCount) {
		toSerialize["unmanaged_count"] = o.UnmanagedCount
	}
	if !IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	if !IsNil(o.UtilizationUpdate) {
		toSerialize["utilization_update"] = o.UtilizationUpdate
	}
	return toSerialize, nil
}

type NullableIpamStatistics struct {
	value *IpamStatistics
	isSet bool
}

func (v NullableIpamStatistics) Get() *IpamStatistics {
	return v.value
}

func (v *NullableIpamStatistics) Set(val *IpamStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamStatistics(val *IpamStatistics) *NullableIpamStatistics {
	return &NullableIpamStatistics{value: val, isSet: true}
}

func (v NullableIpamStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
