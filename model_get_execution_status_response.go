/*
1Click Swap API

API for One-Click Swaps

API version: 0.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oneclick

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the GetExecutionStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExecutionStatusResponse{}

// GetExecutionStatusResponse struct for GetExecutionStatusResponse
type GetExecutionStatusResponse struct {
	// Quote response from original request
	QuoteResponse QuoteResponse `json:"quoteResponse"`
	Status        string        `json:"status"`
	// Last time the state was updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Details of actual swaps and withdrawals
	SwapDetails SwapDetails `json:"swapDetails"`
}

type _GetExecutionStatusResponse GetExecutionStatusResponse

// NewGetExecutionStatusResponse instantiates a new GetExecutionStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExecutionStatusResponse(quoteResponse QuoteResponse, status string, updatedAt time.Time, swapDetails SwapDetails) *GetExecutionStatusResponse {
	this := GetExecutionStatusResponse{}
	this.QuoteResponse = quoteResponse
	this.Status = status
	this.UpdatedAt = updatedAt
	this.SwapDetails = swapDetails
	return &this
}

// NewGetExecutionStatusResponseWithDefaults instantiates a new GetExecutionStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExecutionStatusResponseWithDefaults() *GetExecutionStatusResponse {
	this := GetExecutionStatusResponse{}
	return &this
}

// GetQuoteResponse returns the QuoteResponse field value
func (o *GetExecutionStatusResponse) GetQuoteResponse() QuoteResponse {
	if o == nil {
		var ret QuoteResponse
		return ret
	}

	return o.QuoteResponse
}

// GetQuoteResponseOk returns a tuple with the QuoteResponse field value
// and a boolean to check if the value has been set.
func (o *GetExecutionStatusResponse) GetQuoteResponseOk() (*QuoteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteResponse, true
}

// SetQuoteResponse sets field value
func (o *GetExecutionStatusResponse) SetQuoteResponse(v QuoteResponse) {
	o.QuoteResponse = v
}

// GetStatus returns the Status field value
func (o *GetExecutionStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetExecutionStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetExecutionStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetExecutionStatusResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetExecutionStatusResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetExecutionStatusResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSwapDetails returns the SwapDetails field value
func (o *GetExecutionStatusResponse) GetSwapDetails() SwapDetails {
	if o == nil {
		var ret SwapDetails
		return ret
	}

	return o.SwapDetails
}

// GetSwapDetailsOk returns a tuple with the SwapDetails field value
// and a boolean to check if the value has been set.
func (o *GetExecutionStatusResponse) GetSwapDetailsOk() (*SwapDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwapDetails, true
}

// SetSwapDetails sets field value
func (o *GetExecutionStatusResponse) SetSwapDetails(v SwapDetails) {
	o.SwapDetails = v
}

func (o GetExecutionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExecutionStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quoteResponse"] = o.QuoteResponse
	toSerialize["status"] = o.Status
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["swapDetails"] = o.SwapDetails
	return toSerialize, nil
}

func (o *GetExecutionStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quoteResponse",
		"status",
		"updatedAt",
		"swapDetails",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetExecutionStatusResponse := _GetExecutionStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetExecutionStatusResponse)

	if err != nil {
		return err
	}

	*o = GetExecutionStatusResponse(varGetExecutionStatusResponse)

	return err
}

type NullableGetExecutionStatusResponse struct {
	value *GetExecutionStatusResponse
	isSet bool
}

func (v NullableGetExecutionStatusResponse) Get() *GetExecutionStatusResponse {
	return v.value
}

func (v *NullableGetExecutionStatusResponse) Set(val *GetExecutionStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExecutionStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExecutionStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExecutionStatusResponse(val *GetExecutionStatusResponse) *NullableGetExecutionStatusResponse {
	return &NullableGetExecutionStatusResponse{value: val, isSet: true}
}

func (v NullableGetExecutionStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExecutionStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
