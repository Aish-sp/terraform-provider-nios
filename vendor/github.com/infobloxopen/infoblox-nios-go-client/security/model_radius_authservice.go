/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
)

// checks if the RadiusAuthservice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadiusAuthservice{}

// RadiusAuthservice struct for RadiusAuthservice
type RadiusAuthservice struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The number of times to attempt to contact an accounting RADIUS server.
	AcctRetries *int64 `json:"acct_retries,omitempty"`
	// The number of seconds to wait for a response from the RADIUS server.
	AcctTimeout *int64 `json:"acct_timeout,omitempty"`
	// The number of times to attempt to contact an authentication RADIUS server.
	AuthRetries *int64 `json:"auth_retries,omitempty"`
	// The number of seconds to wait for a response from the RADIUS server.
	AuthTimeout *int64 `json:"auth_timeout,omitempty"`
	// The TTL of cached authentication data in seconds.
	CacheTtl *int64 `json:"cache_ttl,omitempty"`
	// The RADIUS descriptive comment.
	Comment *string `json:"comment,omitempty"`
	// Determines whether the RADIUS authentication service is disabled.
	Disable *bool `json:"disable,omitempty"`
	// Determines whether the authentication cache is enabled.
	EnableCache *bool `json:"enable_cache,omitempty"`
	// The way to contact the RADIUS server.
	Mode *string `json:"mode,omitempty"`
	// The RADIUS authentication service name.
	Name *string `json:"name,omitempty"`
	// The time period to wait before retrying a server that has been marked as down.
	RecoveryInterval *int64 `json:"recovery_interval,omitempty"`
	// The ordered list of RADIUS authentication servers.
	Servers []RadiusAuthserviceServers `json:"servers,omitempty"`
}

// NewRadiusAuthservice instantiates a new RadiusAuthservice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusAuthservice() *RadiusAuthservice {
	this := RadiusAuthservice{}
	return &this
}

// NewRadiusAuthserviceWithDefaults instantiates a new RadiusAuthservice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusAuthserviceWithDefaults() *RadiusAuthservice {
	this := RadiusAuthservice{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *RadiusAuthservice) SetRef(v string) {
	o.Ref = &v
}

// GetAcctRetries returns the AcctRetries field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetAcctRetries() int64 {
	if o == nil || IsNil(o.AcctRetries) {
		var ret int64
		return ret
	}
	return *o.AcctRetries
}

// GetAcctRetriesOk returns a tuple with the AcctRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetAcctRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.AcctRetries) {
		return nil, false
	}
	return o.AcctRetries, true
}

// HasAcctRetries returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasAcctRetries() bool {
	if o != nil && !IsNil(o.AcctRetries) {
		return true
	}

	return false
}

// SetAcctRetries gets a reference to the given int64 and assigns it to the AcctRetries field.
func (o *RadiusAuthservice) SetAcctRetries(v int64) {
	o.AcctRetries = &v
}

// GetAcctTimeout returns the AcctTimeout field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetAcctTimeout() int64 {
	if o == nil || IsNil(o.AcctTimeout) {
		var ret int64
		return ret
	}
	return *o.AcctTimeout
}

// GetAcctTimeoutOk returns a tuple with the AcctTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetAcctTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.AcctTimeout) {
		return nil, false
	}
	return o.AcctTimeout, true
}

// HasAcctTimeout returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasAcctTimeout() bool {
	if o != nil && !IsNil(o.AcctTimeout) {
		return true
	}

	return false
}

// SetAcctTimeout gets a reference to the given int64 and assigns it to the AcctTimeout field.
func (o *RadiusAuthservice) SetAcctTimeout(v int64) {
	o.AcctTimeout = &v
}

// GetAuthRetries returns the AuthRetries field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetAuthRetries() int64 {
	if o == nil || IsNil(o.AuthRetries) {
		var ret int64
		return ret
	}
	return *o.AuthRetries
}

// GetAuthRetriesOk returns a tuple with the AuthRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetAuthRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthRetries) {
		return nil, false
	}
	return o.AuthRetries, true
}

// HasAuthRetries returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasAuthRetries() bool {
	if o != nil && !IsNil(o.AuthRetries) {
		return true
	}

	return false
}

// SetAuthRetries gets a reference to the given int64 and assigns it to the AuthRetries field.
func (o *RadiusAuthservice) SetAuthRetries(v int64) {
	o.AuthRetries = &v
}

// GetAuthTimeout returns the AuthTimeout field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetAuthTimeout() int64 {
	if o == nil || IsNil(o.AuthTimeout) {
		var ret int64
		return ret
	}
	return *o.AuthTimeout
}

// GetAuthTimeoutOk returns a tuple with the AuthTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetAuthTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTimeout) {
		return nil, false
	}
	return o.AuthTimeout, true
}

// HasAuthTimeout returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasAuthTimeout() bool {
	if o != nil && !IsNil(o.AuthTimeout) {
		return true
	}

	return false
}

// SetAuthTimeout gets a reference to the given int64 and assigns it to the AuthTimeout field.
func (o *RadiusAuthservice) SetAuthTimeout(v int64) {
	o.AuthTimeout = &v
}

// GetCacheTtl returns the CacheTtl field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetCacheTtl() int64 {
	if o == nil || IsNil(o.CacheTtl) {
		var ret int64
		return ret
	}
	return *o.CacheTtl
}

// GetCacheTtlOk returns a tuple with the CacheTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetCacheTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheTtl) {
		return nil, false
	}
	return o.CacheTtl, true
}

// HasCacheTtl returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasCacheTtl() bool {
	if o != nil && !IsNil(o.CacheTtl) {
		return true
	}

	return false
}

// SetCacheTtl gets a reference to the given int64 and assigns it to the CacheTtl field.
func (o *RadiusAuthservice) SetCacheTtl(v int64) {
	o.CacheTtl = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RadiusAuthservice) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *RadiusAuthservice) SetDisable(v bool) {
	o.Disable = &v
}

// GetEnableCache returns the EnableCache field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetEnableCache() bool {
	if o == nil || IsNil(o.EnableCache) {
		var ret bool
		return ret
	}
	return *o.EnableCache
}

// GetEnableCacheOk returns a tuple with the EnableCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetEnableCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCache) {
		return nil, false
	}
	return o.EnableCache, true
}

// HasEnableCache returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasEnableCache() bool {
	if o != nil && !IsNil(o.EnableCache) {
		return true
	}

	return false
}

// SetEnableCache gets a reference to the given bool and assigns it to the EnableCache field.
func (o *RadiusAuthservice) SetEnableCache(v bool) {
	o.EnableCache = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RadiusAuthservice) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RadiusAuthservice) SetName(v string) {
	o.Name = &v
}

// GetRecoveryInterval returns the RecoveryInterval field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetRecoveryInterval() int64 {
	if o == nil || IsNil(o.RecoveryInterval) {
		var ret int64
		return ret
	}
	return *o.RecoveryInterval
}

// GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetRecoveryIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.RecoveryInterval) {
		return nil, false
	}
	return o.RecoveryInterval, true
}

// HasRecoveryInterval returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasRecoveryInterval() bool {
	if o != nil && !IsNil(o.RecoveryInterval) {
		return true
	}

	return false
}

// SetRecoveryInterval gets a reference to the given int64 and assigns it to the RecoveryInterval field.
func (o *RadiusAuthservice) SetRecoveryInterval(v int64) {
	o.RecoveryInterval = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *RadiusAuthservice) GetServers() []RadiusAuthserviceServers {
	if o == nil || IsNil(o.Servers) {
		var ret []RadiusAuthserviceServers
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthservice) GetServersOk() ([]RadiusAuthserviceServers, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *RadiusAuthservice) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []RadiusAuthserviceServers and assigns it to the Servers field.
func (o *RadiusAuthservice) SetServers(v []RadiusAuthserviceServers) {
	o.Servers = v
}

func (o RadiusAuthservice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadiusAuthservice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AcctRetries) {
		toSerialize["acct_retries"] = o.AcctRetries
	}
	if !IsNil(o.AcctTimeout) {
		toSerialize["acct_timeout"] = o.AcctTimeout
	}
	if !IsNil(o.AuthRetries) {
		toSerialize["auth_retries"] = o.AuthRetries
	}
	if !IsNil(o.AuthTimeout) {
		toSerialize["auth_timeout"] = o.AuthTimeout
	}
	if !IsNil(o.CacheTtl) {
		toSerialize["cache_ttl"] = o.CacheTtl
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.EnableCache) {
		toSerialize["enable_cache"] = o.EnableCache
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RecoveryInterval) {
		toSerialize["recovery_interval"] = o.RecoveryInterval
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableRadiusAuthservice struct {
	value *RadiusAuthservice
	isSet bool
}

func (v NullableRadiusAuthservice) Get() *RadiusAuthservice {
	return v.value
}

func (v *NullableRadiusAuthservice) Set(val *RadiusAuthservice) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusAuthservice) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusAuthservice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusAuthservice(val *RadiusAuthservice) *NullableRadiusAuthservice {
	return &NullableRadiusAuthservice{value: val, isSet: true}
}

func (v NullableRadiusAuthservice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusAuthservice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
