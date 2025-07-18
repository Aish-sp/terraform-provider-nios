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

// checks if the Nsgroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Nsgroup{}

// Nsgroup struct for Nsgroup
type Nsgroup struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// Comment for the name server group; maximum 256 characters.
	Comment *string `json:"comment,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// The list of external primary servers.
	ExternalPrimaries []NsgroupExternalPrimaries `json:"external_primaries,omitempty"`
	// The list of external secondary servers.
	ExternalSecondaries []NsgroupExternalSecondaries `json:"external_secondaries,omitempty"`
	// The grid primary servers for this group.
	GridPrimary []NsgroupGridPrimary `json:"grid_primary,omitempty"`
	// The list with Grid members that are secondary servers for this group.
	GridSecondaries []NsgroupGridSecondaries `json:"grid_secondaries,omitempty"`
	// Determines if this name server group is the Grid default.
	IsGridDefault *bool `json:"is_grid_default,omitempty"`
	// Determines if the \"multiple DNS primaries\" feature is enabled for the group.
	IsMultimaster *bool `json:"is_multimaster,omitempty"`
	// The name of this name server group.
	Name *string `json:"name,omitempty"`
	// This flag controls whether the group is using an external primary. Note that modification of this field requires passing values for \"grid_secondaries\" and \"external_primaries\".
	UseExternalPrimary *bool `json:"use_external_primary,omitempty"`
}

// NewNsgroup instantiates a new Nsgroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsgroup() *Nsgroup {
	this := Nsgroup{}
	return &this
}

// NewNsgroupWithDefaults instantiates a new Nsgroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsgroupWithDefaults() *Nsgroup {
	this := Nsgroup{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Nsgroup) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Nsgroup) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Nsgroup) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Nsgroup) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Nsgroup) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Nsgroup) SetComment(v string) {
	o.Comment = &v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Nsgroup) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Nsgroup) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Nsgroup) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetExternalPrimaries returns the ExternalPrimaries field value if set, zero value otherwise.
func (o *Nsgroup) GetExternalPrimaries() []NsgroupExternalPrimaries {
	if o == nil || IsNil(o.ExternalPrimaries) {
		var ret []NsgroupExternalPrimaries
		return ret
	}
	return o.ExternalPrimaries
}

// GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetExternalPrimariesOk() ([]NsgroupExternalPrimaries, bool) {
	if o == nil || IsNil(o.ExternalPrimaries) {
		return nil, false
	}
	return o.ExternalPrimaries, true
}

// HasExternalPrimaries returns a boolean if a field has been set.
func (o *Nsgroup) HasExternalPrimaries() bool {
	if o != nil && !IsNil(o.ExternalPrimaries) {
		return true
	}

	return false
}

// SetExternalPrimaries gets a reference to the given []NsgroupExternalPrimaries and assigns it to the ExternalPrimaries field.
func (o *Nsgroup) SetExternalPrimaries(v []NsgroupExternalPrimaries) {
	o.ExternalPrimaries = v
}

// GetExternalSecondaries returns the ExternalSecondaries field value if set, zero value otherwise.
func (o *Nsgroup) GetExternalSecondaries() []NsgroupExternalSecondaries {
	if o == nil || IsNil(o.ExternalSecondaries) {
		var ret []NsgroupExternalSecondaries
		return ret
	}
	return o.ExternalSecondaries
}

// GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetExternalSecondariesOk() ([]NsgroupExternalSecondaries, bool) {
	if o == nil || IsNil(o.ExternalSecondaries) {
		return nil, false
	}
	return o.ExternalSecondaries, true
}

// HasExternalSecondaries returns a boolean if a field has been set.
func (o *Nsgroup) HasExternalSecondaries() bool {
	if o != nil && !IsNil(o.ExternalSecondaries) {
		return true
	}

	return false
}

// SetExternalSecondaries gets a reference to the given []NsgroupExternalSecondaries and assigns it to the ExternalSecondaries field.
func (o *Nsgroup) SetExternalSecondaries(v []NsgroupExternalSecondaries) {
	o.ExternalSecondaries = v
}

// GetGridPrimary returns the GridPrimary field value if set, zero value otherwise.
func (o *Nsgroup) GetGridPrimary() []NsgroupGridPrimary {
	if o == nil || IsNil(o.GridPrimary) {
		var ret []NsgroupGridPrimary
		return ret
	}
	return o.GridPrimary
}

// GetGridPrimaryOk returns a tuple with the GridPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetGridPrimaryOk() ([]NsgroupGridPrimary, bool) {
	if o == nil || IsNil(o.GridPrimary) {
		return nil, false
	}
	return o.GridPrimary, true
}

// HasGridPrimary returns a boolean if a field has been set.
func (o *Nsgroup) HasGridPrimary() bool {
	if o != nil && !IsNil(o.GridPrimary) {
		return true
	}

	return false
}

// SetGridPrimary gets a reference to the given []NsgroupGridPrimary and assigns it to the GridPrimary field.
func (o *Nsgroup) SetGridPrimary(v []NsgroupGridPrimary) {
	o.GridPrimary = v
}

// GetGridSecondaries returns the GridSecondaries field value if set, zero value otherwise.
func (o *Nsgroup) GetGridSecondaries() []NsgroupGridSecondaries {
	if o == nil || IsNil(o.GridSecondaries) {
		var ret []NsgroupGridSecondaries
		return ret
	}
	return o.GridSecondaries
}

// GetGridSecondariesOk returns a tuple with the GridSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetGridSecondariesOk() ([]NsgroupGridSecondaries, bool) {
	if o == nil || IsNil(o.GridSecondaries) {
		return nil, false
	}
	return o.GridSecondaries, true
}

// HasGridSecondaries returns a boolean if a field has been set.
func (o *Nsgroup) HasGridSecondaries() bool {
	if o != nil && !IsNil(o.GridSecondaries) {
		return true
	}

	return false
}

// SetGridSecondaries gets a reference to the given []NsgroupGridSecondaries and assigns it to the GridSecondaries field.
func (o *Nsgroup) SetGridSecondaries(v []NsgroupGridSecondaries) {
	o.GridSecondaries = v
}

// GetIsGridDefault returns the IsGridDefault field value if set, zero value otherwise.
func (o *Nsgroup) GetIsGridDefault() bool {
	if o == nil || IsNil(o.IsGridDefault) {
		var ret bool
		return ret
	}
	return *o.IsGridDefault
}

// GetIsGridDefaultOk returns a tuple with the IsGridDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetIsGridDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGridDefault) {
		return nil, false
	}
	return o.IsGridDefault, true
}

// HasIsGridDefault returns a boolean if a field has been set.
func (o *Nsgroup) HasIsGridDefault() bool {
	if o != nil && !IsNil(o.IsGridDefault) {
		return true
	}

	return false
}

// SetIsGridDefault gets a reference to the given bool and assigns it to the IsGridDefault field.
func (o *Nsgroup) SetIsGridDefault(v bool) {
	o.IsGridDefault = &v
}

// GetIsMultimaster returns the IsMultimaster field value if set, zero value otherwise.
func (o *Nsgroup) GetIsMultimaster() bool {
	if o == nil || IsNil(o.IsMultimaster) {
		var ret bool
		return ret
	}
	return *o.IsMultimaster
}

// GetIsMultimasterOk returns a tuple with the IsMultimaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetIsMultimasterOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultimaster) {
		return nil, false
	}
	return o.IsMultimaster, true
}

// HasIsMultimaster returns a boolean if a field has been set.
func (o *Nsgroup) HasIsMultimaster() bool {
	if o != nil && !IsNil(o.IsMultimaster) {
		return true
	}

	return false
}

// SetIsMultimaster gets a reference to the given bool and assigns it to the IsMultimaster field.
func (o *Nsgroup) SetIsMultimaster(v bool) {
	o.IsMultimaster = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Nsgroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Nsgroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Nsgroup) SetName(v string) {
	o.Name = &v
}

// GetUseExternalPrimary returns the UseExternalPrimary field value if set, zero value otherwise.
func (o *Nsgroup) GetUseExternalPrimary() bool {
	if o == nil || IsNil(o.UseExternalPrimary) {
		var ret bool
		return ret
	}
	return *o.UseExternalPrimary
}

// GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Nsgroup) GetUseExternalPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.UseExternalPrimary) {
		return nil, false
	}
	return o.UseExternalPrimary, true
}

// HasUseExternalPrimary returns a boolean if a field has been set.
func (o *Nsgroup) HasUseExternalPrimary() bool {
	if o != nil && !IsNil(o.UseExternalPrimary) {
		return true
	}

	return false
}

// SetUseExternalPrimary gets a reference to the given bool and assigns it to the UseExternalPrimary field.
func (o *Nsgroup) SetUseExternalPrimary(v bool) {
	o.UseExternalPrimary = &v
}

func (o Nsgroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Nsgroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.ExternalPrimaries) {
		toSerialize["external_primaries"] = o.ExternalPrimaries
	}
	if !IsNil(o.ExternalSecondaries) {
		toSerialize["external_secondaries"] = o.ExternalSecondaries
	}
	if !IsNil(o.GridPrimary) {
		toSerialize["grid_primary"] = o.GridPrimary
	}
	if !IsNil(o.GridSecondaries) {
		toSerialize["grid_secondaries"] = o.GridSecondaries
	}
	if !IsNil(o.IsGridDefault) {
		toSerialize["is_grid_default"] = o.IsGridDefault
	}
	if !IsNil(o.IsMultimaster) {
		toSerialize["is_multimaster"] = o.IsMultimaster
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UseExternalPrimary) {
		toSerialize["use_external_primary"] = o.UseExternalPrimary
	}
	return toSerialize, nil
}

type NullableNsgroup struct {
	value *Nsgroup
	isSet bool
}

func (v NullableNsgroup) Get() *Nsgroup {
	return v.value
}

func (v *NullableNsgroup) Set(val *Nsgroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNsgroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNsgroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsgroup(val *Nsgroup) *NullableNsgroup {
	return &NullableNsgroup{value: val, isSet: true}
}

func (v NullableNsgroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsgroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
