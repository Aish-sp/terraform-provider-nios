/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the GridSecurityBannerSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridSecurityBannerSetting{}

// GridSecurityBannerSetting struct for GridSecurityBannerSetting
type GridSecurityBannerSetting struct {
	// The security level color.
	Color *string `json:"color,omitempty"`
	// The security level.
	Level *string `json:"level,omitempty"`
	// The classification message to be displayed.
	Message *string `json:"message,omitempty"`
	// If set to True, the security banner will be displayed on the header and footer of the Grid Manager screen, including the Login screen.
	Enable               *bool `json:"enable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridSecurityBannerSetting GridSecurityBannerSetting

// NewGridSecurityBannerSetting instantiates a new GridSecurityBannerSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridSecurityBannerSetting() *GridSecurityBannerSetting {
	this := GridSecurityBannerSetting{}
	return &this
}

// NewGridSecurityBannerSettingWithDefaults instantiates a new GridSecurityBannerSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridSecurityBannerSettingWithDefaults() *GridSecurityBannerSetting {
	this := GridSecurityBannerSetting{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *GridSecurityBannerSetting) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSecurityBannerSetting) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *GridSecurityBannerSetting) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *GridSecurityBannerSetting) SetColor(v string) {
	o.Color = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GridSecurityBannerSetting) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSecurityBannerSetting) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GridSecurityBannerSetting) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GridSecurityBannerSetting) SetLevel(v string) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GridSecurityBannerSetting) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSecurityBannerSetting) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GridSecurityBannerSetting) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GridSecurityBannerSetting) SetMessage(v string) {
	o.Message = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GridSecurityBannerSetting) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSecurityBannerSetting) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GridSecurityBannerSetting) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *GridSecurityBannerSetting) SetEnable(v bool) {
	o.Enable = &v
}

func (o GridSecurityBannerSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridSecurityBannerSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridSecurityBannerSetting) UnmarshalJSON(data []byte) (err error) {
	varGridSecurityBannerSetting := _GridSecurityBannerSetting{}

	err = json.Unmarshal(data, &varGridSecurityBannerSetting)

	if err != nil {
		return err
	}

	*o = GridSecurityBannerSetting(varGridSecurityBannerSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "color")
		delete(additionalProperties, "level")
		delete(additionalProperties, "message")
		delete(additionalProperties, "enable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridSecurityBannerSetting struct {
	value *GridSecurityBannerSetting
	isSet bool
}

func (v NullableGridSecurityBannerSetting) Get() *GridSecurityBannerSetting {
	return v.value
}

func (v *NullableGridSecurityBannerSetting) Set(val *GridSecurityBannerSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableGridSecurityBannerSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableGridSecurityBannerSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridSecurityBannerSetting(val *GridSecurityBannerSetting) *NullableGridSecurityBannerSetting {
	return &NullableGridSecurityBannerSetting{value: val, isSet: true}
}

func (v NullableGridSecurityBannerSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridSecurityBannerSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
