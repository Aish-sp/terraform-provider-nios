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

// checks if the AdmingroupLockoutSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdmingroupLockoutSetting{}

// AdmingroupLockoutSetting struct for AdmingroupLockoutSetting
type AdmingroupLockoutSetting struct {
	// Enable/disable sequential failed login attempts lockout for local users
	EnableSequentialFailedLoginAttemptsLockout *bool `json:"enable_sequential_failed_login_attempts_lockout,omitempty"`
	// The number of failed login attempts
	SequentialAttempts *int64 `json:"sequential_attempts,omitempty"`
	// Time period the account remains locked after sequential failed login attempt lockout.
	FailedLockoutDuration *int64 `json:"failed_lockout_duration,omitempty"`
	// Never unlock option is also provided and if set then user account is locked forever and only super user can unlock this account
	NeverUnlockUser      *bool `json:"never_unlock_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdmingroupLockoutSetting AdmingroupLockoutSetting

// NewAdmingroupLockoutSetting instantiates a new AdmingroupLockoutSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdmingroupLockoutSetting() *AdmingroupLockoutSetting {
	this := AdmingroupLockoutSetting{}
	return &this
}

// NewAdmingroupLockoutSettingWithDefaults instantiates a new AdmingroupLockoutSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdmingroupLockoutSettingWithDefaults() *AdmingroupLockoutSetting {
	this := AdmingroupLockoutSetting{}
	return &this
}

// GetEnableSequentialFailedLoginAttemptsLockout returns the EnableSequentialFailedLoginAttemptsLockout field value if set, zero value otherwise.
func (o *AdmingroupLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockout() bool {
	if o == nil || IsNil(o.EnableSequentialFailedLoginAttemptsLockout) {
		var ret bool
		return ret
	}
	return *o.EnableSequentialFailedLoginAttemptsLockout
}

// GetEnableSequentialFailedLoginAttemptsLockoutOk returns a tuple with the EnableSequentialFailedLoginAttemptsLockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupLockoutSetting) GetEnableSequentialFailedLoginAttemptsLockoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSequentialFailedLoginAttemptsLockout) {
		return nil, false
	}
	return o.EnableSequentialFailedLoginAttemptsLockout, true
}

// HasEnableSequentialFailedLoginAttemptsLockout returns a boolean if a field has been set.
func (o *AdmingroupLockoutSetting) HasEnableSequentialFailedLoginAttemptsLockout() bool {
	if o != nil && !IsNil(o.EnableSequentialFailedLoginAttemptsLockout) {
		return true
	}

	return false
}

// SetEnableSequentialFailedLoginAttemptsLockout gets a reference to the given bool and assigns it to the EnableSequentialFailedLoginAttemptsLockout field.
func (o *AdmingroupLockoutSetting) SetEnableSequentialFailedLoginAttemptsLockout(v bool) {
	o.EnableSequentialFailedLoginAttemptsLockout = &v
}

// GetSequentialAttempts returns the SequentialAttempts field value if set, zero value otherwise.
func (o *AdmingroupLockoutSetting) GetSequentialAttempts() int64 {
	if o == nil || IsNil(o.SequentialAttempts) {
		var ret int64
		return ret
	}
	return *o.SequentialAttempts
}

// GetSequentialAttemptsOk returns a tuple with the SequentialAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupLockoutSetting) GetSequentialAttemptsOk() (*int64, bool) {
	if o == nil || IsNil(o.SequentialAttempts) {
		return nil, false
	}
	return o.SequentialAttempts, true
}

// HasSequentialAttempts returns a boolean if a field has been set.
func (o *AdmingroupLockoutSetting) HasSequentialAttempts() bool {
	if o != nil && !IsNil(o.SequentialAttempts) {
		return true
	}

	return false
}

// SetSequentialAttempts gets a reference to the given int64 and assigns it to the SequentialAttempts field.
func (o *AdmingroupLockoutSetting) SetSequentialAttempts(v int64) {
	o.SequentialAttempts = &v
}

// GetFailedLockoutDuration returns the FailedLockoutDuration field value if set, zero value otherwise.
func (o *AdmingroupLockoutSetting) GetFailedLockoutDuration() int64 {
	if o == nil || IsNil(o.FailedLockoutDuration) {
		var ret int64
		return ret
	}
	return *o.FailedLockoutDuration
}

// GetFailedLockoutDurationOk returns a tuple with the FailedLockoutDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupLockoutSetting) GetFailedLockoutDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.FailedLockoutDuration) {
		return nil, false
	}
	return o.FailedLockoutDuration, true
}

// HasFailedLockoutDuration returns a boolean if a field has been set.
func (o *AdmingroupLockoutSetting) HasFailedLockoutDuration() bool {
	if o != nil && !IsNil(o.FailedLockoutDuration) {
		return true
	}

	return false
}

// SetFailedLockoutDuration gets a reference to the given int64 and assigns it to the FailedLockoutDuration field.
func (o *AdmingroupLockoutSetting) SetFailedLockoutDuration(v int64) {
	o.FailedLockoutDuration = &v
}

// GetNeverUnlockUser returns the NeverUnlockUser field value if set, zero value otherwise.
func (o *AdmingroupLockoutSetting) GetNeverUnlockUser() bool {
	if o == nil || IsNil(o.NeverUnlockUser) {
		var ret bool
		return ret
	}
	return *o.NeverUnlockUser
}

// GetNeverUnlockUserOk returns a tuple with the NeverUnlockUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupLockoutSetting) GetNeverUnlockUserOk() (*bool, bool) {
	if o == nil || IsNil(o.NeverUnlockUser) {
		return nil, false
	}
	return o.NeverUnlockUser, true
}

// HasNeverUnlockUser returns a boolean if a field has been set.
func (o *AdmingroupLockoutSetting) HasNeverUnlockUser() bool {
	if o != nil && !IsNil(o.NeverUnlockUser) {
		return true
	}

	return false
}

// SetNeverUnlockUser gets a reference to the given bool and assigns it to the NeverUnlockUser field.
func (o *AdmingroupLockoutSetting) SetNeverUnlockUser(v bool) {
	o.NeverUnlockUser = &v
}

func (o AdmingroupLockoutSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdmingroupLockoutSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableSequentialFailedLoginAttemptsLockout) {
		toSerialize["enable_sequential_failed_login_attempts_lockout"] = o.EnableSequentialFailedLoginAttemptsLockout
	}
	if !IsNil(o.SequentialAttempts) {
		toSerialize["sequential_attempts"] = o.SequentialAttempts
	}
	if !IsNil(o.FailedLockoutDuration) {
		toSerialize["failed_lockout_duration"] = o.FailedLockoutDuration
	}
	if !IsNil(o.NeverUnlockUser) {
		toSerialize["never_unlock_user"] = o.NeverUnlockUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdmingroupLockoutSetting) UnmarshalJSON(data []byte) (err error) {
	varAdmingroupLockoutSetting := _AdmingroupLockoutSetting{}

	err = json.Unmarshal(data, &varAdmingroupLockoutSetting)

	if err != nil {
		return err
	}

	*o = AdmingroupLockoutSetting(varAdmingroupLockoutSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_sequential_failed_login_attempts_lockout")
		delete(additionalProperties, "sequential_attempts")
		delete(additionalProperties, "failed_lockout_duration")
		delete(additionalProperties, "never_unlock_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdmingroupLockoutSetting struct {
	value *AdmingroupLockoutSetting
	isSet bool
}

func (v NullableAdmingroupLockoutSetting) Get() *AdmingroupLockoutSetting {
	return v.value
}

func (v *NullableAdmingroupLockoutSetting) Set(val *AdmingroupLockoutSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableAdmingroupLockoutSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableAdmingroupLockoutSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdmingroupLockoutSetting(val *AdmingroupLockoutSetting) *NullableAdmingroupLockoutSetting {
	return &NullableAdmingroupLockoutSetting{value: val, isSet: true}
}

func (v NullableAdmingroupLockoutSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdmingroupLockoutSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
