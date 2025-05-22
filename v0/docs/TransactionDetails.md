# TransactionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Transaction hash | 
**ExplorerUrl** | **string** | Explorer URL for the transaction | 

## Methods

### NewTransactionDetails

`func NewTransactionDetails(hash string, explorerUrl string, ) *TransactionDetails`

NewTransactionDetails instantiates a new TransactionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDetailsWithDefaults

`func NewTransactionDetailsWithDefaults() *TransactionDetails`

NewTransactionDetailsWithDefaults instantiates a new TransactionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *TransactionDetails) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TransactionDetails) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TransactionDetails) SetHash(v string)`

SetHash sets Hash field to given value.


### GetExplorerUrl

`func (o *TransactionDetails) GetExplorerUrl() string`

GetExplorerUrl returns the ExplorerUrl field if non-nil, zero value otherwise.

### GetExplorerUrlOk

`func (o *TransactionDetails) GetExplorerUrlOk() (*string, bool)`

GetExplorerUrlOk returns a tuple with the ExplorerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerUrl

`func (o *TransactionDetails) SetExplorerUrl(v string)`

SetExplorerUrl sets ExplorerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


