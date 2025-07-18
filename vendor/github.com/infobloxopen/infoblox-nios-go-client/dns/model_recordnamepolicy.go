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

// checks if the Recordnamepolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recordnamepolicy{}

// Recordnamepolicy struct for Recordnamepolicy
type Recordnamepolicy struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// Determines whether the record name policy is Grid default.
	IsDefault *bool `json:"is_default,omitempty"`
	// The name of the record name policy object.
	Name *string `json:"name,omitempty"`
	// Determines whether the record name policy is a predefined one.
	PreDefined *bool `json:"pre_defined,omitempty"`
	// The POSIX regular expression the record names should match in order to comply with the record name policy.
	Regex *string `json:"regex,omitempty"`
}

// NewRecordnamepolicy instantiates a new Recordnamepolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordnamepolicy() *Recordnamepolicy {
	this := Recordnamepolicy{}
	return &this
}

// NewRecordnamepolicyWithDefaults instantiates a new Recordnamepolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordnamepolicyWithDefaults() *Recordnamepolicy {
	this := Recordnamepolicy{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Recordnamepolicy) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recordnamepolicy) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Recordnamepolicy) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Recordnamepolicy) SetRef(v string) {
	o.Ref = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *Recordnamepolicy) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recordnamepolicy) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *Recordnamepolicy) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *Recordnamepolicy) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Recordnamepolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recordnamepolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Recordnamepolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Recordnamepolicy) SetName(v string) {
	o.Name = &v
}

// GetPreDefined returns the PreDefined field value if set, zero value otherwise.
func (o *Recordnamepolicy) GetPreDefined() bool {
	if o == nil || IsNil(o.PreDefined) {
		var ret bool
		return ret
	}
	return *o.PreDefined
}

// GetPreDefinedOk returns a tuple with the PreDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recordnamepolicy) GetPreDefinedOk() (*bool, bool) {
	if o == nil || IsNil(o.PreDefined) {
		return nil, false
	}
	return o.PreDefined, true
}

// HasPreDefined returns a boolean if a field has been set.
func (o *Recordnamepolicy) HasPreDefined() bool {
	if o != nil && !IsNil(o.PreDefined) {
		return true
	}

	return false
}

// SetPreDefined gets a reference to the given bool and assigns it to the PreDefined field.
func (o *Recordnamepolicy) SetPreDefined(v bool) {
	o.PreDefined = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *Recordnamepolicy) GetRegex() string {
	if o == nil || IsNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recordnamepolicy) GetRegexOk() (*string, bool) {
	if o == nil || IsNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *Recordnamepolicy) HasRegex() bool {
	if o != nil && !IsNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *Recordnamepolicy) SetRegex(v string) {
	o.Regex = &v
}

func (o Recordnamepolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recordnamepolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PreDefined) {
		toSerialize["pre_defined"] = o.PreDefined
	}
	if !IsNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	return toSerialize, nil
}

type NullableRecordnamepolicy struct {
	value *Recordnamepolicy
	isSet bool
}

func (v NullableRecordnamepolicy) Get() *Recordnamepolicy {
	return v.value
}

func (v *NullableRecordnamepolicy) Set(val *Recordnamepolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordnamepolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordnamepolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordnamepolicy(val *Recordnamepolicy) *NullableRecordnamepolicy {
	return &NullableRecordnamepolicy{value: val, isSet: true}
}

func (v NullableRecordnamepolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordnamepolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
