# SubmitDepositTxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | **string** | Transaction hash of your deposit | 
**DepositAddress** | **string** | Deposit address for the quote | 

## Methods

### NewSubmitDepositTxRequest

`func NewSubmitDepositTxRequest(txHash string, depositAddress string, ) *SubmitDepositTxRequest`

NewSubmitDepositTxRequest instantiates a new SubmitDepositTxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDepositTxRequestWithDefaults

`func NewSubmitDepositTxRequestWithDefaults() *SubmitDepositTxRequest`

NewSubmitDepositTxRequestWithDefaults instantiates a new SubmitDepositTxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *SubmitDepositTxRequest) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *SubmitDepositTxRequest) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *SubmitDepositTxRequest) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetDepositAddress

`func (o *SubmitDepositTxRequest) GetDepositAddress() string`

GetDepositAddress returns the DepositAddress field if non-nil, zero value otherwise.

### GetDepositAddressOk

`func (o *SubmitDepositTxRequest) GetDepositAddressOk() (*string, bool)`

GetDepositAddressOk returns a tuple with the DepositAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddress

`func (o *SubmitDepositTxRequest) SetDepositAddress(v string)`

SetDepositAddress sets DepositAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


