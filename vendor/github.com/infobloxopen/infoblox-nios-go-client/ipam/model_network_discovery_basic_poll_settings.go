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

// checks if the NetworkDiscoveryBasicPollSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDiscoveryBasicPollSettings{}

// NetworkDiscoveryBasicPollSettings struct for NetworkDiscoveryBasicPollSettings
type NetworkDiscoveryBasicPollSettings struct {
	// Determines whether port scanning is enabled or not.
	PortScanning *bool `json:"port_scanning,omitempty"`
	// Determines whether device profile is enabled or not.
	DeviceProfile *bool `json:"device_profile,omitempty"`
	// Determines whether SNMP collection is enabled or not.
	SnmpCollection *bool `json:"snmp_collection,omitempty"`
	// Determines whether CLI collection is enabled or not.
	CliCollection *bool `json:"cli_collection,omitempty"`
	// Determines whether netbios scanning is enabled or not.
	NetbiosScanning *bool `json:"netbios_scanning,omitempty"`
	// Determines whether complete ping sweep is enabled or not.
	CompletePingSweep *bool `json:"complete_ping_sweep,omitempty"`
	// Determines whether smart subnet ping sweep is enabled or not.
	SmartSubnetPingSweep *bool `json:"smart_subnet_ping_sweep,omitempty"`
	// Determines whether auto ARP refresh before switch port polling is enabled or not.
	AutoArpRefreshBeforeSwitchPortPolling *bool `json:"auto_arp_refresh_before_switch_port_polling,omitempty"`
	// A switch port data collection polling mode.
	SwitchPortDataCollectionPolling         *string                                                                   `json:"switch_port_data_collection_polling,omitempty"`
	SwitchPortDataCollectionPollingSchedule *NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule `json:"switch_port_data_collection_polling_schedule,omitempty"`
	// Indicates the interval for switch port data collection polling.
	SwitchPortDataCollectionPollingInterval *int64 `json:"switch_port_data_collection_polling_interval,omitempty"`
	// Credential group.
	CredentialGroup *string `json:"credential_group,omitempty"`
	// Polling Frequency Modifier.
	PollingFrequencyModifier *string `json:"polling_frequency_modifier,omitempty"`
	// Use Global Polling Frequency Modifier.
	UseGlobalPollingFrequencyModifier *bool `json:"use_global_polling_frequency_modifier,omitempty"`
	AdditionalProperties              map[string]interface{}
}

type _NetworkDiscoveryBasicPollSettings NetworkDiscoveryBasicPollSettings

// NewNetworkDiscoveryBasicPollSettings instantiates a new NetworkDiscoveryBasicPollSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDiscoveryBasicPollSettings() *NetworkDiscoveryBasicPollSettings {
	this := NetworkDiscoveryBasicPollSettings{}
	return &this
}

// NewNetworkDiscoveryBasicPollSettingsWithDefaults instantiates a new NetworkDiscoveryBasicPollSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDiscoveryBasicPollSettingsWithDefaults() *NetworkDiscoveryBasicPollSettings {
	this := NetworkDiscoveryBasicPollSettings{}
	return &this
}

// GetPortScanning returns the PortScanning field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetPortScanning() bool {
	if o == nil || IsNil(o.PortScanning) {
		var ret bool
		return ret
	}
	return *o.PortScanning
}

// GetPortScanningOk returns a tuple with the PortScanning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetPortScanningOk() (*bool, bool) {
	if o == nil || IsNil(o.PortScanning) {
		return nil, false
	}
	return o.PortScanning, true
}

// HasPortScanning returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasPortScanning() bool {
	if o != nil && !IsNil(o.PortScanning) {
		return true
	}

	return false
}

// SetPortScanning gets a reference to the given bool and assigns it to the PortScanning field.
func (o *NetworkDiscoveryBasicPollSettings) SetPortScanning(v bool) {
	o.PortScanning = &v
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetDeviceProfile() bool {
	if o == nil || IsNil(o.DeviceProfile) {
		var ret bool
		return ret
	}
	return *o.DeviceProfile
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetDeviceProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceProfile) {
		return nil, false
	}
	return o.DeviceProfile, true
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasDeviceProfile() bool {
	if o != nil && !IsNil(o.DeviceProfile) {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given bool and assigns it to the DeviceProfile field.
func (o *NetworkDiscoveryBasicPollSettings) SetDeviceProfile(v bool) {
	o.DeviceProfile = &v
}

// GetSnmpCollection returns the SnmpCollection field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetSnmpCollection() bool {
	if o == nil || IsNil(o.SnmpCollection) {
		var ret bool
		return ret
	}
	return *o.SnmpCollection
}

// GetSnmpCollectionOk returns a tuple with the SnmpCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetSnmpCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.SnmpCollection) {
		return nil, false
	}
	return o.SnmpCollection, true
}

// HasSnmpCollection returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasSnmpCollection() bool {
	if o != nil && !IsNil(o.SnmpCollection) {
		return true
	}

	return false
}

// SetSnmpCollection gets a reference to the given bool and assigns it to the SnmpCollection field.
func (o *NetworkDiscoveryBasicPollSettings) SetSnmpCollection(v bool) {
	o.SnmpCollection = &v
}

// GetCliCollection returns the CliCollection field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetCliCollection() bool {
	if o == nil || IsNil(o.CliCollection) {
		var ret bool
		return ret
	}
	return *o.CliCollection
}

// GetCliCollectionOk returns a tuple with the CliCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetCliCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.CliCollection) {
		return nil, false
	}
	return o.CliCollection, true
}

// HasCliCollection returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasCliCollection() bool {
	if o != nil && !IsNil(o.CliCollection) {
		return true
	}

	return false
}

// SetCliCollection gets a reference to the given bool and assigns it to the CliCollection field.
func (o *NetworkDiscoveryBasicPollSettings) SetCliCollection(v bool) {
	o.CliCollection = &v
}

// GetNetbiosScanning returns the NetbiosScanning field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetNetbiosScanning() bool {
	if o == nil || IsNil(o.NetbiosScanning) {
		var ret bool
		return ret
	}
	return *o.NetbiosScanning
}

// GetNetbiosScanningOk returns a tuple with the NetbiosScanning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetNetbiosScanningOk() (*bool, bool) {
	if o == nil || IsNil(o.NetbiosScanning) {
		return nil, false
	}
	return o.NetbiosScanning, true
}

// HasNetbiosScanning returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasNetbiosScanning() bool {
	if o != nil && !IsNil(o.NetbiosScanning) {
		return true
	}

	return false
}

// SetNetbiosScanning gets a reference to the given bool and assigns it to the NetbiosScanning field.
func (o *NetworkDiscoveryBasicPollSettings) SetNetbiosScanning(v bool) {
	o.NetbiosScanning = &v
}

// GetCompletePingSweep returns the CompletePingSweep field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetCompletePingSweep() bool {
	if o == nil || IsNil(o.CompletePingSweep) {
		var ret bool
		return ret
	}
	return *o.CompletePingSweep
}

// GetCompletePingSweepOk returns a tuple with the CompletePingSweep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetCompletePingSweepOk() (*bool, bool) {
	if o == nil || IsNil(o.CompletePingSweep) {
		return nil, false
	}
	return o.CompletePingSweep, true
}

// HasCompletePingSweep returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasCompletePingSweep() bool {
	if o != nil && !IsNil(o.CompletePingSweep) {
		return true
	}

	return false
}

// SetCompletePingSweep gets a reference to the given bool and assigns it to the CompletePingSweep field.
func (o *NetworkDiscoveryBasicPollSettings) SetCompletePingSweep(v bool) {
	o.CompletePingSweep = &v
}

// GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetSmartSubnetPingSweep() bool {
	if o == nil || IsNil(o.SmartSubnetPingSweep) {
		var ret bool
		return ret
	}
	return *o.SmartSubnetPingSweep
}

// GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetSmartSubnetPingSweepOk() (*bool, bool) {
	if o == nil || IsNil(o.SmartSubnetPingSweep) {
		return nil, false
	}
	return o.SmartSubnetPingSweep, true
}

// HasSmartSubnetPingSweep returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasSmartSubnetPingSweep() bool {
	if o != nil && !IsNil(o.SmartSubnetPingSweep) {
		return true
	}

	return false
}

// SetSmartSubnetPingSweep gets a reference to the given bool and assigns it to the SmartSubnetPingSweep field.
func (o *NetworkDiscoveryBasicPollSettings) SetSmartSubnetPingSweep(v bool) {
	o.SmartSubnetPingSweep = &v
}

// GetAutoArpRefreshBeforeSwitchPortPolling returns the AutoArpRefreshBeforeSwitchPortPolling field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPolling() bool {
	if o == nil || IsNil(o.AutoArpRefreshBeforeSwitchPortPolling) {
		var ret bool
		return ret
	}
	return *o.AutoArpRefreshBeforeSwitchPortPolling
}

// GetAutoArpRefreshBeforeSwitchPortPollingOk returns a tuple with the AutoArpRefreshBeforeSwitchPortPolling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetAutoArpRefreshBeforeSwitchPortPollingOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoArpRefreshBeforeSwitchPortPolling) {
		return nil, false
	}
	return o.AutoArpRefreshBeforeSwitchPortPolling, true
}

// HasAutoArpRefreshBeforeSwitchPortPolling returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasAutoArpRefreshBeforeSwitchPortPolling() bool {
	if o != nil && !IsNil(o.AutoArpRefreshBeforeSwitchPortPolling) {
		return true
	}

	return false
}

// SetAutoArpRefreshBeforeSwitchPortPolling gets a reference to the given bool and assigns it to the AutoArpRefreshBeforeSwitchPortPolling field.
func (o *NetworkDiscoveryBasicPollSettings) SetAutoArpRefreshBeforeSwitchPortPolling(v bool) {
	o.AutoArpRefreshBeforeSwitchPortPolling = &v
}

// GetSwitchPortDataCollectionPolling returns the SwitchPortDataCollectionPolling field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPolling() string {
	if o == nil || IsNil(o.SwitchPortDataCollectionPolling) {
		var ret string
		return ret
	}
	return *o.SwitchPortDataCollectionPolling
}

// GetSwitchPortDataCollectionPollingOk returns a tuple with the SwitchPortDataCollectionPolling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchPortDataCollectionPolling) {
		return nil, false
	}
	return o.SwitchPortDataCollectionPolling, true
}

// HasSwitchPortDataCollectionPolling returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPolling() bool {
	if o != nil && !IsNil(o.SwitchPortDataCollectionPolling) {
		return true
	}

	return false
}

// SetSwitchPortDataCollectionPolling gets a reference to the given string and assigns it to the SwitchPortDataCollectionPolling field.
func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPolling(v string) {
	o.SwitchPortDataCollectionPolling = &v
}

// GetSwitchPortDataCollectionPollingSchedule returns the SwitchPortDataCollectionPollingSchedule field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingSchedule() NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule {
	if o == nil || IsNil(o.SwitchPortDataCollectionPollingSchedule) {
		var ret NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule
		return ret
	}
	return *o.SwitchPortDataCollectionPollingSchedule
}

// GetSwitchPortDataCollectionPollingScheduleOk returns a tuple with the SwitchPortDataCollectionPollingSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingScheduleOk() (*NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, bool) {
	if o == nil || IsNil(o.SwitchPortDataCollectionPollingSchedule) {
		return nil, false
	}
	return o.SwitchPortDataCollectionPollingSchedule, true
}

// HasSwitchPortDataCollectionPollingSchedule returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingSchedule() bool {
	if o != nil && !IsNil(o.SwitchPortDataCollectionPollingSchedule) {
		return true
	}

	return false
}

// SetSwitchPortDataCollectionPollingSchedule gets a reference to the given NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule and assigns it to the SwitchPortDataCollectionPollingSchedule field.
func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingSchedule(v NetworkdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule) {
	o.SwitchPortDataCollectionPollingSchedule = &v
}

// GetSwitchPortDataCollectionPollingInterval returns the SwitchPortDataCollectionPollingInterval field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingInterval() int64 {
	if o == nil || IsNil(o.SwitchPortDataCollectionPollingInterval) {
		var ret int64
		return ret
	}
	return *o.SwitchPortDataCollectionPollingInterval
}

// GetSwitchPortDataCollectionPollingIntervalOk returns a tuple with the SwitchPortDataCollectionPollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetSwitchPortDataCollectionPollingIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchPortDataCollectionPollingInterval) {
		return nil, false
	}
	return o.SwitchPortDataCollectionPollingInterval, true
}

// HasSwitchPortDataCollectionPollingInterval returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasSwitchPortDataCollectionPollingInterval() bool {
	if o != nil && !IsNil(o.SwitchPortDataCollectionPollingInterval) {
		return true
	}

	return false
}

// SetSwitchPortDataCollectionPollingInterval gets a reference to the given int64 and assigns it to the SwitchPortDataCollectionPollingInterval field.
func (o *NetworkDiscoveryBasicPollSettings) SetSwitchPortDataCollectionPollingInterval(v int64) {
	o.SwitchPortDataCollectionPollingInterval = &v
}

// GetCredentialGroup returns the CredentialGroup field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetCredentialGroup() string {
	if o == nil || IsNil(o.CredentialGroup) {
		var ret string
		return ret
	}
	return *o.CredentialGroup
}

// GetCredentialGroupOk returns a tuple with the CredentialGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetCredentialGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialGroup) {
		return nil, false
	}
	return o.CredentialGroup, true
}

// HasCredentialGroup returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasCredentialGroup() bool {
	if o != nil && !IsNil(o.CredentialGroup) {
		return true
	}

	return false
}

// SetCredentialGroup gets a reference to the given string and assigns it to the CredentialGroup field.
func (o *NetworkDiscoveryBasicPollSettings) SetCredentialGroup(v string) {
	o.CredentialGroup = &v
}

// GetPollingFrequencyModifier returns the PollingFrequencyModifier field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetPollingFrequencyModifier() string {
	if o == nil || IsNil(o.PollingFrequencyModifier) {
		var ret string
		return ret
	}
	return *o.PollingFrequencyModifier
}

// GetPollingFrequencyModifierOk returns a tuple with the PollingFrequencyModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetPollingFrequencyModifierOk() (*string, bool) {
	if o == nil || IsNil(o.PollingFrequencyModifier) {
		return nil, false
	}
	return o.PollingFrequencyModifier, true
}

// HasPollingFrequencyModifier returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasPollingFrequencyModifier() bool {
	if o != nil && !IsNil(o.PollingFrequencyModifier) {
		return true
	}

	return false
}

// SetPollingFrequencyModifier gets a reference to the given string and assigns it to the PollingFrequencyModifier field.
func (o *NetworkDiscoveryBasicPollSettings) SetPollingFrequencyModifier(v string) {
	o.PollingFrequencyModifier = &v
}

// GetUseGlobalPollingFrequencyModifier returns the UseGlobalPollingFrequencyModifier field value if set, zero value otherwise.
func (o *NetworkDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifier() bool {
	if o == nil || IsNil(o.UseGlobalPollingFrequencyModifier) {
		var ret bool
		return ret
	}
	return *o.UseGlobalPollingFrequencyModifier
}

// GetUseGlobalPollingFrequencyModifierOk returns a tuple with the UseGlobalPollingFrequencyModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBasicPollSettings) GetUseGlobalPollingFrequencyModifierOk() (*bool, bool) {
	if o == nil || IsNil(o.UseGlobalPollingFrequencyModifier) {
		return nil, false
	}
	return o.UseGlobalPollingFrequencyModifier, true
}

// HasUseGlobalPollingFrequencyModifier returns a boolean if a field has been set.
func (o *NetworkDiscoveryBasicPollSettings) HasUseGlobalPollingFrequencyModifier() bool {
	if o != nil && !IsNil(o.UseGlobalPollingFrequencyModifier) {
		return true
	}

	return false
}

// SetUseGlobalPollingFrequencyModifier gets a reference to the given bool and assigns it to the UseGlobalPollingFrequencyModifier field.
func (o *NetworkDiscoveryBasicPollSettings) SetUseGlobalPollingFrequencyModifier(v bool) {
	o.UseGlobalPollingFrequencyModifier = &v
}

func (o NetworkDiscoveryBasicPollSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDiscoveryBasicPollSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortScanning) {
		toSerialize["port_scanning"] = o.PortScanning
	}
	if !IsNil(o.DeviceProfile) {
		toSerialize["device_profile"] = o.DeviceProfile
	}
	if !IsNil(o.SnmpCollection) {
		toSerialize["snmp_collection"] = o.SnmpCollection
	}
	if !IsNil(o.CliCollection) {
		toSerialize["cli_collection"] = o.CliCollection
	}
	if !IsNil(o.NetbiosScanning) {
		toSerialize["netbios_scanning"] = o.NetbiosScanning
	}
	if !IsNil(o.CompletePingSweep) {
		toSerialize["complete_ping_sweep"] = o.CompletePingSweep
	}
	if !IsNil(o.SmartSubnetPingSweep) {
		toSerialize["smart_subnet_ping_sweep"] = o.SmartSubnetPingSweep
	}
	if !IsNil(o.AutoArpRefreshBeforeSwitchPortPolling) {
		toSerialize["auto_arp_refresh_before_switch_port_polling"] = o.AutoArpRefreshBeforeSwitchPortPolling
	}
	if !IsNil(o.SwitchPortDataCollectionPolling) {
		toSerialize["switch_port_data_collection_polling"] = o.SwitchPortDataCollectionPolling
	}
	if !IsNil(o.SwitchPortDataCollectionPollingSchedule) {
		toSerialize["switch_port_data_collection_polling_schedule"] = o.SwitchPortDataCollectionPollingSchedule
	}
	if !IsNil(o.SwitchPortDataCollectionPollingInterval) {
		toSerialize["switch_port_data_collection_polling_interval"] = o.SwitchPortDataCollectionPollingInterval
	}
	if !IsNil(o.CredentialGroup) {
		toSerialize["credential_group"] = o.CredentialGroup
	}
	if !IsNil(o.PollingFrequencyModifier) {
		toSerialize["polling_frequency_modifier"] = o.PollingFrequencyModifier
	}
	if !IsNil(o.UseGlobalPollingFrequencyModifier) {
		toSerialize["use_global_polling_frequency_modifier"] = o.UseGlobalPollingFrequencyModifier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDiscoveryBasicPollSettings) UnmarshalJSON(data []byte) (err error) {
	varNetworkDiscoveryBasicPollSettings := _NetworkDiscoveryBasicPollSettings{}

	err = json.Unmarshal(data, &varNetworkDiscoveryBasicPollSettings)

	if err != nil {
		return err
	}

	*o = NetworkDiscoveryBasicPollSettings(varNetworkDiscoveryBasicPollSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "port_scanning")
		delete(additionalProperties, "device_profile")
		delete(additionalProperties, "snmp_collection")
		delete(additionalProperties, "cli_collection")
		delete(additionalProperties, "netbios_scanning")
		delete(additionalProperties, "complete_ping_sweep")
		delete(additionalProperties, "smart_subnet_ping_sweep")
		delete(additionalProperties, "auto_arp_refresh_before_switch_port_polling")
		delete(additionalProperties, "switch_port_data_collection_polling")
		delete(additionalProperties, "switch_port_data_collection_polling_schedule")
		delete(additionalProperties, "switch_port_data_collection_polling_interval")
		delete(additionalProperties, "credential_group")
		delete(additionalProperties, "polling_frequency_modifier")
		delete(additionalProperties, "use_global_polling_frequency_modifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDiscoveryBasicPollSettings struct {
	value *NetworkDiscoveryBasicPollSettings
	isSet bool
}

func (v NullableNetworkDiscoveryBasicPollSettings) Get() *NetworkDiscoveryBasicPollSettings {
	return v.value
}

func (v *NullableNetworkDiscoveryBasicPollSettings) Set(val *NetworkDiscoveryBasicPollSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDiscoveryBasicPollSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDiscoveryBasicPollSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDiscoveryBasicPollSettings(val *NetworkDiscoveryBasicPollSettings) *NullableNetworkDiscoveryBasicPollSettings {
	return &NullableNetworkDiscoveryBasicPollSettings{value: val, isSet: true}
}

func (v NullableNetworkDiscoveryBasicPollSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDiscoveryBasicPollSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
