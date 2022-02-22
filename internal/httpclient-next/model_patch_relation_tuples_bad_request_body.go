/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PatchRelationTuplesBadRequestBody PatchRelationTuplesBadRequestBody PatchRelationTuplesBadRequestBody PatchRelationTuplesBadRequestBody PatchRelationTuplesBadRequestBody PatchRelationTuplesBadRequestBody patch relation tuples bad request body
type PatchRelationTuplesBadRequestBody struct {
	// code
	Code *int64 `json:"code,omitempty"`
	// details
	Details []map[string]interface{} `json:"details,omitempty"`
	// message
	Message *string `json:"message,omitempty"`
	// reason
	Reason *string `json:"reason,omitempty"`
	// request
	Request *string `json:"request,omitempty"`
	// status
	Status *string `json:"status,omitempty"`
}

// NewPatchRelationTuplesBadRequestBody instantiates a new PatchRelationTuplesBadRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchRelationTuplesBadRequestBody() *PatchRelationTuplesBadRequestBody {
	this := PatchRelationTuplesBadRequestBody{}
	return &this
}

// NewPatchRelationTuplesBadRequestBodyWithDefaults instantiates a new PatchRelationTuplesBadRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchRelationTuplesBadRequestBodyWithDefaults() *PatchRelationTuplesBadRequestBody {
	this := PatchRelationTuplesBadRequestBody{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *PatchRelationTuplesBadRequestBody) SetCode(v int64) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetDetails() []map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetDetailsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []map[string]interface{} and assigns it to the Details field.
func (o *PatchRelationTuplesBadRequestBody) SetDetails(v []map[string]interface{}) {
	o.Details = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PatchRelationTuplesBadRequestBody) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PatchRelationTuplesBadRequestBody) SetReason(v string) {
	o.Reason = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *PatchRelationTuplesBadRequestBody) SetRequest(v string) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchRelationTuplesBadRequestBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRelationTuplesBadRequestBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchRelationTuplesBadRequestBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchRelationTuplesBadRequestBody) SetStatus(v string) {
	o.Status = &v
}

func (o PatchRelationTuplesBadRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePatchRelationTuplesBadRequestBody struct {
	value *PatchRelationTuplesBadRequestBody
	isSet bool
}

func (v NullablePatchRelationTuplesBadRequestBody) Get() *PatchRelationTuplesBadRequestBody {
	return v.value
}

func (v *NullablePatchRelationTuplesBadRequestBody) Set(val *PatchRelationTuplesBadRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRelationTuplesBadRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRelationTuplesBadRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRelationTuplesBadRequestBody(val *PatchRelationTuplesBadRequestBody) *NullablePatchRelationTuplesBadRequestBody {
	return &NullablePatchRelationTuplesBadRequestBody{value: val, isSet: true}
}

func (v NullablePatchRelationTuplesBadRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRelationTuplesBadRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
