/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneAuthAwsRte53ZoneInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneAuthAwsRte53ZoneInfo{}

// ZoneAuthAwsRte53ZoneInfo struct for ZoneAuthAwsRte53ZoneInfo
type ZoneAuthAwsRte53ZoneInfo struct {
	// List of AWS VPC strings that are associated with this zone.
	AssociatedVpcs []string `json:"associated_vpcs,omitempty"`
	// User specified caller reference when zone was created.
	CallerReference *string `json:"caller_reference,omitempty"`
	// ID of delegation set associated with this zone.
	DelegationSetId *string `json:"delegation_set_id,omitempty"`
	// AWS route 53 assigned ID for this zone.
	HostedZoneId *string `json:"hosted_zone_id,omitempty"`
	// List of AWS name servers that are authoritative for this domain name.
	NameServers []string `json:"name_servers,omitempty"`
	// Number of resource record sets in the hosted zone.
	RecordSetCount *int64 `json:"record_set_count,omitempty"`
	// Indicates whether private or public zone.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZoneAuthAwsRte53ZoneInfo ZoneAuthAwsRte53ZoneInfo

// NewZoneAuthAwsRte53ZoneInfo instantiates a new ZoneAuthAwsRte53ZoneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneAuthAwsRte53ZoneInfo() *ZoneAuthAwsRte53ZoneInfo {
	this := ZoneAuthAwsRte53ZoneInfo{}
	return &this
}

// NewZoneAuthAwsRte53ZoneInfoWithDefaults instantiates a new ZoneAuthAwsRte53ZoneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneAuthAwsRte53ZoneInfoWithDefaults() *ZoneAuthAwsRte53ZoneInfo {
	this := ZoneAuthAwsRte53ZoneInfo{}
	return &this
}

// GetAssociatedVpcs returns the AssociatedVpcs field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetAssociatedVpcs() []string {
	if o == nil || IsNil(o.AssociatedVpcs) {
		var ret []string
		return ret
	}
	return o.AssociatedVpcs
}

// GetAssociatedVpcsOk returns a tuple with the AssociatedVpcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetAssociatedVpcsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedVpcs) {
		return nil, false
	}
	return o.AssociatedVpcs, true
}

// HasAssociatedVpcs returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasAssociatedVpcs() bool {
	if o != nil && !IsNil(o.AssociatedVpcs) {
		return true
	}

	return false
}

// SetAssociatedVpcs gets a reference to the given []string and assigns it to the AssociatedVpcs field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetAssociatedVpcs(v []string) {
	o.AssociatedVpcs = v
}

// GetCallerReference returns the CallerReference field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetCallerReference() string {
	if o == nil || IsNil(o.CallerReference) {
		var ret string
		return ret
	}
	return *o.CallerReference
}

// GetCallerReferenceOk returns a tuple with the CallerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetCallerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CallerReference) {
		return nil, false
	}
	return o.CallerReference, true
}

// HasCallerReference returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasCallerReference() bool {
	if o != nil && !IsNil(o.CallerReference) {
		return true
	}

	return false
}

// SetCallerReference gets a reference to the given string and assigns it to the CallerReference field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetCallerReference(v string) {
	o.CallerReference = &v
}

// GetDelegationSetId returns the DelegationSetId field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetDelegationSetId() string {
	if o == nil || IsNil(o.DelegationSetId) {
		var ret string
		return ret
	}
	return *o.DelegationSetId
}

// GetDelegationSetIdOk returns a tuple with the DelegationSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetDelegationSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DelegationSetId) {
		return nil, false
	}
	return o.DelegationSetId, true
}

// HasDelegationSetId returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasDelegationSetId() bool {
	if o != nil && !IsNil(o.DelegationSetId) {
		return true
	}

	return false
}

// SetDelegationSetId gets a reference to the given string and assigns it to the DelegationSetId field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetDelegationSetId(v string) {
	o.DelegationSetId = &v
}

// GetHostedZoneId returns the HostedZoneId field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetHostedZoneId() string {
	if o == nil || IsNil(o.HostedZoneId) {
		var ret string
		return ret
	}
	return *o.HostedZoneId
}

// GetHostedZoneIdOk returns a tuple with the HostedZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetHostedZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostedZoneId) {
		return nil, false
	}
	return o.HostedZoneId, true
}

// HasHostedZoneId returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasHostedZoneId() bool {
	if o != nil && !IsNil(o.HostedZoneId) {
		return true
	}

	return false
}

// SetHostedZoneId gets a reference to the given string and assigns it to the HostedZoneId field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetHostedZoneId(v string) {
	o.HostedZoneId = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetNameServers() []string {
	if o == nil || IsNil(o.NameServers) {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetNameServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NameServers) {
		return nil, false
	}
	return o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasNameServers() bool {
	if o != nil && !IsNil(o.NameServers) {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetNameServers(v []string) {
	o.NameServers = v
}

// GetRecordSetCount returns the RecordSetCount field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetRecordSetCount() int64 {
	if o == nil || IsNil(o.RecordSetCount) {
		var ret int64
		return ret
	}
	return *o.RecordSetCount
}

// GetRecordSetCountOk returns a tuple with the RecordSetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetRecordSetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordSetCount) {
		return nil, false
	}
	return o.RecordSetCount, true
}

// HasRecordSetCount returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasRecordSetCount() bool {
	if o != nil && !IsNil(o.RecordSetCount) {
		return true
	}

	return false
}

// SetRecordSetCount gets a reference to the given int64 and assigns it to the RecordSetCount field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetRecordSetCount(v int64) {
	o.RecordSetCount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ZoneAuthAwsRte53ZoneInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ZoneAuthAwsRte53ZoneInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ZoneAuthAwsRte53ZoneInfo) SetType(v string) {
	o.Type = &v
}

func (o ZoneAuthAwsRte53ZoneInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneAuthAwsRte53ZoneInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedVpcs) {
		toSerialize["associated_vpcs"] = o.AssociatedVpcs
	}
	if !IsNil(o.CallerReference) {
		toSerialize["caller_reference"] = o.CallerReference
	}
	if !IsNil(o.DelegationSetId) {
		toSerialize["delegation_set_id"] = o.DelegationSetId
	}
	if !IsNil(o.HostedZoneId) {
		toSerialize["hosted_zone_id"] = o.HostedZoneId
	}
	if !IsNil(o.NameServers) {
		toSerialize["name_servers"] = o.NameServers
	}
	if !IsNil(o.RecordSetCount) {
		toSerialize["record_set_count"] = o.RecordSetCount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZoneAuthAwsRte53ZoneInfo) UnmarshalJSON(data []byte) (err error) {
	varZoneAuthAwsRte53ZoneInfo := _ZoneAuthAwsRte53ZoneInfo{}

	err = json.Unmarshal(data, &varZoneAuthAwsRte53ZoneInfo)

	if err != nil {
		return err
	}

	*o = ZoneAuthAwsRte53ZoneInfo(varZoneAuthAwsRte53ZoneInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associated_vpcs")
		delete(additionalProperties, "caller_reference")
		delete(additionalProperties, "delegation_set_id")
		delete(additionalProperties, "hosted_zone_id")
		delete(additionalProperties, "name_servers")
		delete(additionalProperties, "record_set_count")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZoneAuthAwsRte53ZoneInfo struct {
	value *ZoneAuthAwsRte53ZoneInfo
	isSet bool
}

func (v NullableZoneAuthAwsRte53ZoneInfo) Get() *ZoneAuthAwsRte53ZoneInfo {
	return v.value
}

func (v *NullableZoneAuthAwsRte53ZoneInfo) Set(val *ZoneAuthAwsRte53ZoneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneAuthAwsRte53ZoneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneAuthAwsRte53ZoneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneAuthAwsRte53ZoneInfo(val *ZoneAuthAwsRte53ZoneInfo) *NullableZoneAuthAwsRte53ZoneInfo {
	return &NullableZoneAuthAwsRte53ZoneInfo{value: val, isSet: true}
}

func (v NullableZoneAuthAwsRte53ZoneInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneAuthAwsRte53ZoneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
