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
)

// checks if the BadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestResponse{}

// BadRequestResponse struct for BadRequestResponse
type BadRequestResponse struct {
	Message string `json:"message"`
}

type _BadRequestResponse BadRequestResponse

// NewBadRequestResponse instantiates a new BadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestResponse(message string) *BadRequestResponse {
	this := BadRequestResponse{}
	this.Message = message
	return &this
}

// NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestResponseWithDefaults() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *BadRequestResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BadRequestResponse) SetMessage(v string) {
	o.Message = v
}

func (o BadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *BadRequestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varBadRequestResponse := _BadRequestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBadRequestResponse)

	if err != nil {
		return err
	}

	*o = BadRequestResponse(varBadRequestResponse)

	return err
}

type NullableBadRequestResponse struct {
	value *BadRequestResponse
	isSet bool
}

func (v NullableBadRequestResponse) Get() *BadRequestResponse {
	return v.value
}

func (v *NullableBadRequestResponse) Set(val *BadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestResponse(val *BadRequestResponse) *NullableBadRequestResponse {
	return &NullableBadRequestResponse{value: val, isSet: true}
}

func (v NullableBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
