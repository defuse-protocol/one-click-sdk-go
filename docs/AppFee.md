# AppFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | **string** | Intents Account ID where this fee will be transferred to | 
**Fee** | **float32** | Fee for this recipient as part of amountIn in basis points (1/100th of a percent), e.g. 100 for 1% fee | 

## Methods

### NewAppFee

`func NewAppFee(recipient string, fee float32, ) *AppFee`

NewAppFee instantiates a new AppFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppFeeWithDefaults

`func NewAppFeeWithDefaults() *AppFee`

NewAppFeeWithDefaults instantiates a new AppFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *AppFee) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AppFee) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AppFee) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetFee

`func (o *AppFee) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AppFee) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AppFee) SetFee(v float32)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


