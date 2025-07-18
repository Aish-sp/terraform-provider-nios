/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the Allendpoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Allendpoints{}

// Allendpoints struct for Allendpoints
type Allendpoints struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The Grid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN).
	Address *string `json:"address,omitempty"`
	// The Grid endpoint descriptive comment.
	Comment *string `json:"comment,omitempty"`
	// Determines whether a Grid endpoint is disabled or not. When this is set to False, the Grid endpoint is enabled.
	Disable *bool `json:"disable,omitempty"`
	// The name of the Grid Member object that is serving Grid endpoint.
	SubscribingMember *string `json:"subscribing_member,omitempty"`
	// The Grid endpoint type.
	Type *string `json:"type,omitempty"`
	// The Grid endpoint version.
	Version *string `json:"version,omitempty"`
}

// NewAllendpoints instantiates a new Allendpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllendpoints() *Allendpoints {
	this := Allendpoints{}
	return &this
}

// NewAllendpointsWithDefaults instantiates a new Allendpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllendpointsWithDefaults() *Allendpoints {
	this := Allendpoints{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Allendpoints) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Allendpoints) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Allendpoints) SetRef(v string) {
	o.Ref = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Allendpoints) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Allendpoints) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Allendpoints) SetAddress(v string) {
	o.Address = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Allendpoints) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Allendpoints) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Allendpoints) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *Allendpoints) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *Allendpoints) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *Allendpoints) SetDisable(v bool) {
	o.Disable = &v
}

// GetSubscribingMember returns the SubscribingMember field value if set, zero value otherwise.
func (o *Allendpoints) GetSubscribingMember() string {
	if o == nil || IsNil(o.SubscribingMember) {
		var ret string
		return ret
	}
	return *o.SubscribingMember
}

// GetSubscribingMemberOk returns a tuple with the SubscribingMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetSubscribingMemberOk() (*string, bool) {
	if o == nil || IsNil(o.SubscribingMember) {
		return nil, false
	}
	return o.SubscribingMember, true
}

// HasSubscribingMember returns a boolean if a field has been set.
func (o *Allendpoints) HasSubscribingMember() bool {
	if o != nil && !IsNil(o.SubscribingMember) {
		return true
	}

	return false
}

// SetSubscribingMember gets a reference to the given string and assigns it to the SubscribingMember field.
func (o *Allendpoints) SetSubscribingMember(v string) {
	o.SubscribingMember = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Allendpoints) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Allendpoints) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Allendpoints) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Allendpoints) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allendpoints) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Allendpoints) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Allendpoints) SetVersion(v string) {
	o.Version = &v
}

func (o Allendpoints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Allendpoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.SubscribingMember) {
		toSerialize["subscribing_member"] = o.SubscribingMember
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAllendpoints struct {
	value *Allendpoints
	isSet bool
}

func (v NullableAllendpoints) Get() *Allendpoints {
	return v.value
}

func (v *NullableAllendpoints) Set(val *Allendpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableAllendpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableAllendpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllendpoints(val *Allendpoints) *NullableAllendpoints {
	return &NullableAllendpoints{value: val, isSet: true}
}

func (v NullableAllendpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllendpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
