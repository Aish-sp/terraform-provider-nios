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

// checks if the RecordPtrCloudInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordPtrCloudInfo{}

// RecordPtrCloudInfo struct for RecordPtrCloudInfo
type RecordPtrCloudInfo struct {
	DelegatedMember *RecordptrcloudinfoDelegatedMember `json:"delegated_member,omitempty"`
	// Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).
	DelegatedScope *string `json:"delegated_scope,omitempty"`
	// Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.
	DelegatedRoot *string `json:"delegated_root,omitempty"`
	// Determines whether the object was created by the cloud adapter or not.
	OwnedByAdaptor *bool `json:"owned_by_adaptor,omitempty"`
	// Indicates the cloud origin of the object.
	Usage *string `json:"usage,omitempty"`
	// Reference to the tenant object associated with the object, if any.
	Tenant *string `json:"tenant,omitempty"`
	// Indicates the specified cloud management platform.
	MgmtPlatform *string `json:"mgmt_platform,omitempty"`
	// Type of authority over the object.
	AuthorityType        *string `json:"authority_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecordPtrCloudInfo RecordPtrCloudInfo

// NewRecordPtrCloudInfo instantiates a new RecordPtrCloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordPtrCloudInfo() *RecordPtrCloudInfo {
	this := RecordPtrCloudInfo{}
	return &this
}

// NewRecordPtrCloudInfoWithDefaults instantiates a new RecordPtrCloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordPtrCloudInfoWithDefaults() *RecordPtrCloudInfo {
	this := RecordPtrCloudInfo{}
	return &this
}

// GetDelegatedMember returns the DelegatedMember field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetDelegatedMember() RecordptrcloudinfoDelegatedMember {
	if o == nil || IsNil(o.DelegatedMember) {
		var ret RecordptrcloudinfoDelegatedMember
		return ret
	}
	return *o.DelegatedMember
}

// GetDelegatedMemberOk returns a tuple with the DelegatedMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetDelegatedMemberOk() (*RecordptrcloudinfoDelegatedMember, bool) {
	if o == nil || IsNil(o.DelegatedMember) {
		return nil, false
	}
	return o.DelegatedMember, true
}

// HasDelegatedMember returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasDelegatedMember() bool {
	if o != nil && !IsNil(o.DelegatedMember) {
		return true
	}

	return false
}

// SetDelegatedMember gets a reference to the given RecordptrcloudinfoDelegatedMember and assigns it to the DelegatedMember field.
func (o *RecordPtrCloudInfo) SetDelegatedMember(v RecordptrcloudinfoDelegatedMember) {
	o.DelegatedMember = &v
}

// GetDelegatedScope returns the DelegatedScope field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetDelegatedScope() string {
	if o == nil || IsNil(o.DelegatedScope) {
		var ret string
		return ret
	}
	return *o.DelegatedScope
}

// GetDelegatedScopeOk returns a tuple with the DelegatedScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetDelegatedScopeOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedScope) {
		return nil, false
	}
	return o.DelegatedScope, true
}

// HasDelegatedScope returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasDelegatedScope() bool {
	if o != nil && !IsNil(o.DelegatedScope) {
		return true
	}

	return false
}

// SetDelegatedScope gets a reference to the given string and assigns it to the DelegatedScope field.
func (o *RecordPtrCloudInfo) SetDelegatedScope(v string) {
	o.DelegatedScope = &v
}

// GetDelegatedRoot returns the DelegatedRoot field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetDelegatedRoot() string {
	if o == nil || IsNil(o.DelegatedRoot) {
		var ret string
		return ret
	}
	return *o.DelegatedRoot
}

// GetDelegatedRootOk returns a tuple with the DelegatedRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetDelegatedRootOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedRoot) {
		return nil, false
	}
	return o.DelegatedRoot, true
}

// HasDelegatedRoot returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasDelegatedRoot() bool {
	if o != nil && !IsNil(o.DelegatedRoot) {
		return true
	}

	return false
}

// SetDelegatedRoot gets a reference to the given string and assigns it to the DelegatedRoot field.
func (o *RecordPtrCloudInfo) SetDelegatedRoot(v string) {
	o.DelegatedRoot = &v
}

// GetOwnedByAdaptor returns the OwnedByAdaptor field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetOwnedByAdaptor() bool {
	if o == nil || IsNil(o.OwnedByAdaptor) {
		var ret bool
		return ret
	}
	return *o.OwnedByAdaptor
}

// GetOwnedByAdaptorOk returns a tuple with the OwnedByAdaptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetOwnedByAdaptorOk() (*bool, bool) {
	if o == nil || IsNil(o.OwnedByAdaptor) {
		return nil, false
	}
	return o.OwnedByAdaptor, true
}

// HasOwnedByAdaptor returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasOwnedByAdaptor() bool {
	if o != nil && !IsNil(o.OwnedByAdaptor) {
		return true
	}

	return false
}

// SetOwnedByAdaptor gets a reference to the given bool and assigns it to the OwnedByAdaptor field.
func (o *RecordPtrCloudInfo) SetOwnedByAdaptor(v bool) {
	o.OwnedByAdaptor = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *RecordPtrCloudInfo) SetUsage(v string) {
	o.Usage = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *RecordPtrCloudInfo) SetTenant(v string) {
	o.Tenant = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetMgmtPlatform() string {
	if o == nil || IsNil(o.MgmtPlatform) {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtPlatform) {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasMgmtPlatform() bool {
	if o != nil && !IsNil(o.MgmtPlatform) {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *RecordPtrCloudInfo) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetAuthorityType returns the AuthorityType field value if set, zero value otherwise.
func (o *RecordPtrCloudInfo) GetAuthorityType() string {
	if o == nil || IsNil(o.AuthorityType) {
		var ret string
		return ret
	}
	return *o.AuthorityType
}

// GetAuthorityTypeOk returns a tuple with the AuthorityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPtrCloudInfo) GetAuthorityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityType) {
		return nil, false
	}
	return o.AuthorityType, true
}

// HasAuthorityType returns a boolean if a field has been set.
func (o *RecordPtrCloudInfo) HasAuthorityType() bool {
	if o != nil && !IsNil(o.AuthorityType) {
		return true
	}

	return false
}

// SetAuthorityType gets a reference to the given string and assigns it to the AuthorityType field.
func (o *RecordPtrCloudInfo) SetAuthorityType(v string) {
	o.AuthorityType = &v
}

func (o RecordPtrCloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordPtrCloudInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelegatedMember) {
		toSerialize["delegated_member"] = o.DelegatedMember
	}
	if !IsNil(o.DelegatedScope) {
		toSerialize["delegated_scope"] = o.DelegatedScope
	}
	if !IsNil(o.DelegatedRoot) {
		toSerialize["delegated_root"] = o.DelegatedRoot
	}
	if !IsNil(o.OwnedByAdaptor) {
		toSerialize["owned_by_adaptor"] = o.OwnedByAdaptor
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.MgmtPlatform) {
		toSerialize["mgmt_platform"] = o.MgmtPlatform
	}
	if !IsNil(o.AuthorityType) {
		toSerialize["authority_type"] = o.AuthorityType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecordPtrCloudInfo) UnmarshalJSON(data []byte) (err error) {
	varRecordPtrCloudInfo := _RecordPtrCloudInfo{}

	err = json.Unmarshal(data, &varRecordPtrCloudInfo)

	if err != nil {
		return err
	}

	*o = RecordPtrCloudInfo(varRecordPtrCloudInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegated_member")
		delete(additionalProperties, "delegated_scope")
		delete(additionalProperties, "delegated_root")
		delete(additionalProperties, "owned_by_adaptor")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "mgmt_platform")
		delete(additionalProperties, "authority_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecordPtrCloudInfo struct {
	value *RecordPtrCloudInfo
	isSet bool
}

func (v NullableRecordPtrCloudInfo) Get() *RecordPtrCloudInfo {
	return v.value
}

func (v *NullableRecordPtrCloudInfo) Set(val *RecordPtrCloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordPtrCloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordPtrCloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordPtrCloudInfo(val *RecordPtrCloudInfo) *NullableRecordPtrCloudInfo {
	return &NullableRecordPtrCloudInfo{value: val, isSet: true}
}

func (v NullableRecordPtrCloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordPtrCloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
