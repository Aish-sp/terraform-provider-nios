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

// checks if the Scavengingtask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scavengingtask{}

// Scavengingtask struct for Scavengingtask
type Scavengingtask struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The scavenging action.
	Action *string `json:"action,omitempty"`
	// The reference to the object associated with the scavenging task.
	AssociatedObject *string `json:"associated_object,omitempty"`
	// The scavenging process end time.
	EndTime *int64 `json:"end_time,omitempty"`
	// The number of processed during scavenging resource records.
	ProcessedRecords *int64 `json:"processed_records,omitempty"`
	// The number of resource records that are allowed to be reclaimed during the scavenging process.
	ReclaimableRecords *int64 `json:"reclaimable_records,omitempty"`
	// The number of reclaimed during the scavenging process resource records.
	ReclaimedRecords *int64 `json:"reclaimed_records,omitempty"`
	// The scavenging process start time.
	StartTime *int64 `json:"start_time,omitempty"`
	// The scavenging process status. This is a read-only attribute.
	Status *string `json:"status,omitempty"`
}

// NewScavengingtask instantiates a new Scavengingtask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScavengingtask() *Scavengingtask {
	this := Scavengingtask{}
	return &this
}

// NewScavengingtaskWithDefaults instantiates a new Scavengingtask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScavengingtaskWithDefaults() *Scavengingtask {
	this := Scavengingtask{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Scavengingtask) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Scavengingtask) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Scavengingtask) SetRef(v string) {
	o.Ref = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Scavengingtask) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Scavengingtask) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Scavengingtask) SetAction(v string) {
	o.Action = &v
}

// GetAssociatedObject returns the AssociatedObject field value if set, zero value otherwise.
func (o *Scavengingtask) GetAssociatedObject() string {
	if o == nil || IsNil(o.AssociatedObject) {
		var ret string
		return ret
	}
	return *o.AssociatedObject
}

// GetAssociatedObjectOk returns a tuple with the AssociatedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetAssociatedObjectOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedObject) {
		return nil, false
	}
	return o.AssociatedObject, true
}

// HasAssociatedObject returns a boolean if a field has been set.
func (o *Scavengingtask) HasAssociatedObject() bool {
	if o != nil && !IsNil(o.AssociatedObject) {
		return true
	}

	return false
}

// SetAssociatedObject gets a reference to the given string and assigns it to the AssociatedObject field.
func (o *Scavengingtask) SetAssociatedObject(v string) {
	o.AssociatedObject = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Scavengingtask) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Scavengingtask) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *Scavengingtask) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetProcessedRecords returns the ProcessedRecords field value if set, zero value otherwise.
func (o *Scavengingtask) GetProcessedRecords() int64 {
	if o == nil || IsNil(o.ProcessedRecords) {
		var ret int64
		return ret
	}
	return *o.ProcessedRecords
}

// GetProcessedRecordsOk returns a tuple with the ProcessedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetProcessedRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.ProcessedRecords) {
		return nil, false
	}
	return o.ProcessedRecords, true
}

// HasProcessedRecords returns a boolean if a field has been set.
func (o *Scavengingtask) HasProcessedRecords() bool {
	if o != nil && !IsNil(o.ProcessedRecords) {
		return true
	}

	return false
}

// SetProcessedRecords gets a reference to the given int64 and assigns it to the ProcessedRecords field.
func (o *Scavengingtask) SetProcessedRecords(v int64) {
	o.ProcessedRecords = &v
}

// GetReclaimableRecords returns the ReclaimableRecords field value if set, zero value otherwise.
func (o *Scavengingtask) GetReclaimableRecords() int64 {
	if o == nil || IsNil(o.ReclaimableRecords) {
		var ret int64
		return ret
	}
	return *o.ReclaimableRecords
}

// GetReclaimableRecordsOk returns a tuple with the ReclaimableRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetReclaimableRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.ReclaimableRecords) {
		return nil, false
	}
	return o.ReclaimableRecords, true
}

// HasReclaimableRecords returns a boolean if a field has been set.
func (o *Scavengingtask) HasReclaimableRecords() bool {
	if o != nil && !IsNil(o.ReclaimableRecords) {
		return true
	}

	return false
}

// SetReclaimableRecords gets a reference to the given int64 and assigns it to the ReclaimableRecords field.
func (o *Scavengingtask) SetReclaimableRecords(v int64) {
	o.ReclaimableRecords = &v
}

// GetReclaimedRecords returns the ReclaimedRecords field value if set, zero value otherwise.
func (o *Scavengingtask) GetReclaimedRecords() int64 {
	if o == nil || IsNil(o.ReclaimedRecords) {
		var ret int64
		return ret
	}
	return *o.ReclaimedRecords
}

// GetReclaimedRecordsOk returns a tuple with the ReclaimedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetReclaimedRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.ReclaimedRecords) {
		return nil, false
	}
	return o.ReclaimedRecords, true
}

// HasReclaimedRecords returns a boolean if a field has been set.
func (o *Scavengingtask) HasReclaimedRecords() bool {
	if o != nil && !IsNil(o.ReclaimedRecords) {
		return true
	}

	return false
}

// SetReclaimedRecords gets a reference to the given int64 and assigns it to the ReclaimedRecords field.
func (o *Scavengingtask) SetReclaimedRecords(v int64) {
	o.ReclaimedRecords = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Scavengingtask) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Scavengingtask) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *Scavengingtask) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Scavengingtask) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scavengingtask) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Scavengingtask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Scavengingtask) SetStatus(v string) {
	o.Status = &v
}

func (o Scavengingtask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scavengingtask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AssociatedObject) {
		toSerialize["associated_object"] = o.AssociatedObject
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.ProcessedRecords) {
		toSerialize["processed_records"] = o.ProcessedRecords
	}
	if !IsNil(o.ReclaimableRecords) {
		toSerialize["reclaimable_records"] = o.ReclaimableRecords
	}
	if !IsNil(o.ReclaimedRecords) {
		toSerialize["reclaimed_records"] = o.ReclaimedRecords
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableScavengingtask struct {
	value *Scavengingtask
	isSet bool
}

func (v NullableScavengingtask) Get() *Scavengingtask {
	return v.value
}

func (v *NullableScavengingtask) Set(val *Scavengingtask) {
	v.value = val
	v.isSet = true
}

func (v NullableScavengingtask) IsSet() bool {
	return v.isSet
}

func (v *NullableScavengingtask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScavengingtask(val *Scavengingtask) *NullableScavengingtask {
	return &NullableScavengingtask{value: val, isSet: true}
}

func (v NullableScavengingtask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScavengingtask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
