/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
)

// checks if the DiscoveryGridpropertiesSnmpv3Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryGridpropertiesSnmpv3Credentials{}

// DiscoveryGridpropertiesSnmpv3Credentials struct for DiscoveryGridpropertiesSnmpv3Credentials
type DiscoveryGridpropertiesSnmpv3Credentials struct {
	// The SNMPv3 user name.
	User *string `json:"user,omitempty"`
	// Authentication protocol for the SNMPv3 user.
	AuthenticationProtocol *string `json:"authentication_protocol,omitempty"`
	// Authentication password for the SNMPv3 user.
	AuthenticationPassword *string `json:"authentication_password,omitempty"`
	// Privacy protocol for the SNMPv3 user.
	PrivacyProtocol *string `json:"privacy_protocol,omitempty"`
	// Privacy password for the SNMPv3 user.
	PrivacyPassword *string `json:"privacy_password,omitempty"`
	// Comments for the SNMPv3 user.
	Comment *string `json:"comment,omitempty"`
	// Group for the SNMPv3 credential.
	CredentialGroup      *string `json:"credential_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryGridpropertiesSnmpv3Credentials DiscoveryGridpropertiesSnmpv3Credentials

// NewDiscoveryGridpropertiesSnmpv3Credentials instantiates a new DiscoveryGridpropertiesSnmpv3Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryGridpropertiesSnmpv3Credentials() *DiscoveryGridpropertiesSnmpv3Credentials {
	this := DiscoveryGridpropertiesSnmpv3Credentials{}
	return &this
}

// NewDiscoveryGridpropertiesSnmpv3CredentialsWithDefaults instantiates a new DiscoveryGridpropertiesSnmpv3Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryGridpropertiesSnmpv3CredentialsWithDefaults() *DiscoveryGridpropertiesSnmpv3Credentials {
	this := DiscoveryGridpropertiesSnmpv3Credentials{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetUser(v string) {
	o.User = &v
}

// GetAuthenticationProtocol returns the AuthenticationProtocol field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationProtocol() string {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		var ret string
		return ret
	}
	return *o.AuthenticationProtocol
}

// GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationProtocol) {
		return nil, false
	}
	return o.AuthenticationProtocol, true
}

// HasAuthenticationProtocol returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasAuthenticationProtocol() bool {
	if o != nil && !IsNil(o.AuthenticationProtocol) {
		return true
	}

	return false
}

// SetAuthenticationProtocol gets a reference to the given string and assigns it to the AuthenticationProtocol field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetAuthenticationProtocol(v string) {
	o.AuthenticationProtocol = &v
}

// GetAuthenticationPassword returns the AuthenticationPassword field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationPassword() string {
	if o == nil || IsNil(o.AuthenticationPassword) {
		var ret string
		return ret
	}
	return *o.AuthenticationPassword
}

// GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetAuthenticationPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationPassword) {
		return nil, false
	}
	return o.AuthenticationPassword, true
}

// HasAuthenticationPassword returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasAuthenticationPassword() bool {
	if o != nil && !IsNil(o.AuthenticationPassword) {
		return true
	}

	return false
}

// SetAuthenticationPassword gets a reference to the given string and assigns it to the AuthenticationPassword field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetAuthenticationPassword(v string) {
	o.AuthenticationPassword = &v
}

// GetPrivacyProtocol returns the PrivacyProtocol field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyProtocol() string {
	if o == nil || IsNil(o.PrivacyProtocol) {
		var ret string
		return ret
	}
	return *o.PrivacyProtocol
}

// GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyProtocol) {
		return nil, false
	}
	return o.PrivacyProtocol, true
}

// HasPrivacyProtocol returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasPrivacyProtocol() bool {
	if o != nil && !IsNil(o.PrivacyProtocol) {
		return true
	}

	return false
}

// SetPrivacyProtocol gets a reference to the given string and assigns it to the PrivacyProtocol field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetPrivacyProtocol(v string) {
	o.PrivacyProtocol = &v
}

// GetPrivacyPassword returns the PrivacyPassword field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyPassword() string {
	if o == nil || IsNil(o.PrivacyPassword) {
		var ret string
		return ret
	}
	return *o.PrivacyPassword
}

// GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetPrivacyPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPassword) {
		return nil, false
	}
	return o.PrivacyPassword, true
}

// HasPrivacyPassword returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasPrivacyPassword() bool {
	if o != nil && !IsNil(o.PrivacyPassword) {
		return true
	}

	return false
}

// SetPrivacyPassword gets a reference to the given string and assigns it to the PrivacyPassword field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetPrivacyPassword(v string) {
	o.PrivacyPassword = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetComment(v string) {
	o.Comment = &v
}

// GetCredentialGroup returns the CredentialGroup field value if set, zero value otherwise.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCredentialGroup() string {
	if o == nil || IsNil(o.CredentialGroup) {
		var ret string
		return ret
	}
	return *o.CredentialGroup
}

// GetCredentialGroupOk returns a tuple with the CredentialGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) GetCredentialGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialGroup) {
		return nil, false
	}
	return o.CredentialGroup, true
}

// HasCredentialGroup returns a boolean if a field has been set.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) HasCredentialGroup() bool {
	if o != nil && !IsNil(o.CredentialGroup) {
		return true
	}

	return false
}

// SetCredentialGroup gets a reference to the given string and assigns it to the CredentialGroup field.
func (o *DiscoveryGridpropertiesSnmpv3Credentials) SetCredentialGroup(v string) {
	o.CredentialGroup = &v
}

func (o DiscoveryGridpropertiesSnmpv3Credentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryGridpropertiesSnmpv3Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.AuthenticationProtocol) {
		toSerialize["authentication_protocol"] = o.AuthenticationProtocol
	}
	if !IsNil(o.AuthenticationPassword) {
		toSerialize["authentication_password"] = o.AuthenticationPassword
	}
	if !IsNil(o.PrivacyProtocol) {
		toSerialize["privacy_protocol"] = o.PrivacyProtocol
	}
	if !IsNil(o.PrivacyPassword) {
		toSerialize["privacy_password"] = o.PrivacyPassword
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CredentialGroup) {
		toSerialize["credential_group"] = o.CredentialGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryGridpropertiesSnmpv3Credentials) UnmarshalJSON(data []byte) (err error) {
	varDiscoveryGridpropertiesSnmpv3Credentials := _DiscoveryGridpropertiesSnmpv3Credentials{}

	err = json.Unmarshal(data, &varDiscoveryGridpropertiesSnmpv3Credentials)

	if err != nil {
		return err
	}

	*o = DiscoveryGridpropertiesSnmpv3Credentials(varDiscoveryGridpropertiesSnmpv3Credentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "authentication_protocol")
		delete(additionalProperties, "authentication_password")
		delete(additionalProperties, "privacy_protocol")
		delete(additionalProperties, "privacy_password")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "credential_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryGridpropertiesSnmpv3Credentials struct {
	value *DiscoveryGridpropertiesSnmpv3Credentials
	isSet bool
}

func (v NullableDiscoveryGridpropertiesSnmpv3Credentials) Get() *DiscoveryGridpropertiesSnmpv3Credentials {
	return v.value
}

func (v *NullableDiscoveryGridpropertiesSnmpv3Credentials) Set(val *DiscoveryGridpropertiesSnmpv3Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryGridpropertiesSnmpv3Credentials) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryGridpropertiesSnmpv3Credentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryGridpropertiesSnmpv3Credentials(val *DiscoveryGridpropertiesSnmpv3Credentials) *NullableDiscoveryGridpropertiesSnmpv3Credentials {
	return &NullableDiscoveryGridpropertiesSnmpv3Credentials{value: val, isSet: true}
}

func (v NullableDiscoveryGridpropertiesSnmpv3Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryGridpropertiesSnmpv3Credentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
